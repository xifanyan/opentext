package arm

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func (v *Volumn) DatPath() string {
	return filepath.Join(v.BasePath, v.DataSetID, v.ID, v.DatFile)
}

func (v *Volumn) OptPath() string {
	return filepath.Join(v.BasePath, v.DataSetID, v.ID, v.OptFile)
}

// ReadDatHeader reads the header of a .dat file and returns it as a slice of strings.
// If the header has already been read, it is returned from the cache.
// If the header cannot be read, an error is returned.
// The header is cached for further calls to ReadDatHeader.
func (v *Volumn) ReadDatHeader() ([]string, error) {
	var err error

	if v.header != nil {
		return v.header, nil
	}

	path := v.DatPath()
	f, err := os.Open(path)
	if err != nil {
		log.Error().Msgf("open [%s]: %v", path, err)
		return nil, err
	}
	defer f.Close()

	// Create a new reader
	reader := NewReader(f)

	// pull the header
	v.header, err = reader.Read()
	if err != nil {
		log.Error().Msgf("read header [%s]: %v", path, err)
		return nil, err
	}

	return v.header, err
}

// CountColumnWithValue reads the .dat file of a Vol and counts the non-empty values
// in each column, returning a map with column headers as keys and their respective
// non-empty value counts as values.
//
// Parameters:
//
//	None
//
// Returns:
//
//	map[string]int: A map where the keys are column headers and values are counts
//	of non-empty entries in each column.
//	error: An error is returned if there is a problem reading the .dat file.
func (v *Volumn) CountColumnWithValue() (map[string]int, error) {
	path := v.DatPath()
	f, err := os.Open(path)
	if err != nil {
		log.Error().Msgf("open [%s]: %v", path, err)
		return nil, err
	}
	defer f.Close()

	reader := NewReader(f)
	header, err := reader.Read()
	if err != nil {
		log.Error().Msgf("read header [%s]: %v", path, err)
		return nil, err
	}

	countMap := make(map[string]int)
	for _, h := range header {
		countMap[h] = 0
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Error().Msgf("read row [%s]: %v", path, err)
			return nil, err
		}

		for i, value := range record {
			if len(value) > 0 {
				countMap[header[i]]++
			}
		}
	}

	return countMap, nil
}
