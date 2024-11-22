package dataproc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/sourcegraph/conc/pool"
	"github.com/xifanyan/opentext/discovery/r2a/data/arm"
)

type DataProc struct {
	Fields       []Field
	FieldAliases map[string][]string
	FieldTypes   map[string][]string
	*arm.DataSet
}

type Field struct {
	Header    string `json:"header"`
	MappedTo  string `json:"mappedTo"`
	FieldType string `json:"fieldType"`
	Count     int    `json:"count"`
}

func NewDataProc(arm *arm.DataSet) *DataProc {
	return &DataProc{
		DataSet: arm,
	}
}
func (proc *DataProc) Initialize() error {
	// collect volumns information
	rootPath := fmt.Sprintf("%s/%s", proc.BasePath, proc.ID)
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if directory
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".dat") {
			return nil
		}

		// Get relative path from root
		relPath, err := filepath.Rel(proc.BasePath, path)
		if err != nil {
			return err
		}

		// Split path components
		parts := strings.Split(relPath, string(os.PathSeparator))
		if len(parts) != 3 {
			return nil
		}

		vol := arm.Volumn{
			BasePath:  proc.BasePath,
			DataSetID: proc.ID,
			ID:        parts[1],
			DatFile:   parts[2],
			OptFile:   arm.GetOptFileByDatFile(path),
		}

		proc.Volumns = append(proc.Volumns, vol)

		return nil
	})

	if err != nil {
		return err
	}

	// check header consistency and initialize fields info
	header, err := proc.checkHeaderConsistency()
	if err != nil {
		log.Error().Msgf("checkHeaderConsistency: %v", err)
		return err
	}

	proc.Fields = make([]Field, len(header))
	for i, h := range header {
		proc.Fields[i].Header = h
	}

	return err
}

func (t *DataProc) checkHeaderConsistency() ([]string, error) {
	var header []string

	for _, vol := range t.Volumns {
		current, err := vol.ReadDatHeader()
		if err != nil || current == nil || len(current) == 0 {
			return header, err
		}

		if header == nil {
			header = current
		}

		if !reflect.DeepEqual(header, current) {
			return header, fmt.Errorf("header inconsistency: %v != %v", header, current)
		}
	}

	return header, nil
}

func (proc *DataProc) countTotalColumnWithValue() error {

	results := make([]map[string]int, len(proc.Volumns))
	p := pool.New().WithMaxGoroutines(runtime.GOMAXPROCS(0)).WithErrors()
	for i, vol := range proc.Volumns {
		p.Go(func() error {
			countMap, err := vol.CountColumnWithValue()
			if err != nil {
				return err
			}
			results[i] = countMap
			return nil
		})
	}

	if err := p.Wait(); err != nil {
		log.Error().Msgf("p.Wait: %v", err)
		return err
	}

	for i, field := range proc.Fields {
		for _, result := range results {
			if count, ok := result[field.Header]; ok {
				proc.Fields[i].Count += count
			}
		}
	}

	return nil
}

func (proc *DataProc) PrintFieldCounts(sortFlag bool, greaterThanZeroOnlyFlag bool) error {
	if err := proc.countTotalColumnWithValue(); err != nil {
		return err
	}

	if sortFlag {
		sort.Slice(proc.Fields, func(i, j int) bool {
			return proc.Fields[i].Count > proc.Fields[j].Count
		})
	}

	for _, field := range proc.Fields {
		if greaterThanZeroOnlyFlag {
			if field.Count > 0 {
				fmt.Printf("%-50s %-10d %-25s %-25s\n", field.Header, field.Count, field.MappedTo, field.FieldType)
			}
		} else {
			fmt.Printf("%-50s %-10d %-25s %-25s\n", field.Header, field.Count, field.MappedTo, field.FieldType)
		}
	}
	return nil
}

func (proc *DataProc) findHeaderIndex(fieldName string) (int, error) {
	for i, field := range proc.Fields {
		if field.Header == fieldName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("header '%s' not found", fieldName)
}

func (proc *DataProc) PrintTopNFieldValues(fieldName string, maxNumLines int) error {
	headerIndex, err := proc.findHeaderIndex(fieldName)
	if err != nil {
		return err
	}

	printedLines := 0
	for _, vol := range proc.DataSet.Volumns {
		// Open the file
		datPath := fmt.Sprintf("%s/%s/%s/%s", vol.BasePath, vol.DataSetID, vol.ID, vol.DatFile)
		f, err := os.Open(datPath)
		if err != nil {
			return err
		}
		defer f.Close()

		// Read the first line to get the headers
		r := arm.NewReader(f)

		// skip header
		r.NextLine()

		for {
			fields, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			if len(fields) > headerIndex && fields[headerIndex] != "" {
				if maxNumLines == -1 {
					fmt.Println(fields[headerIndex])
					continue
				}

				// Print the specified number of lines when the field has a value
				if printedLines < maxNumLines {
					fmt.Println(fields[headerIndex])
					printedLines++
				} else {
					break
				}
			}
		}

	}

	return nil
}

func (proc *DataProc) LoadFieldAliases(path string) error {
	// Load alias definitions from file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&proc.FieldAliases)
	if err != nil {
		return err
	}

	return nil
}

func (proc *DataProc) MapHeaderToField() error {
	// Create a map for fast lookups
	aliasLookupTable := make(map[string]string)
	for fieldName, headers := range proc.FieldAliases {
		for _, header := range headers {
			aliasLookupTable[header] = fieldName
		}
		aliasLookupTable[fieldName] = fieldName
	}

	for i, field := range proc.Fields {
		if fieldName, ok := aliasLookupTable[field.Header]; ok {
			field.MappedTo = fieldName
			proc.Fields[i] = field
		}
	}

	return nil
}
