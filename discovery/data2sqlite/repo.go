package data2sqlite

import (
	"encoding/csv"
	"os"

	"github.com/rs/zerolog/log"
)

type repo interface {
	Export() chan []string
}

type CsvRepo struct {
	fields []string
	path   string
}

func NewCsvRepo(path string) *CsvRepo {
	return &CsvRepo{
		path: path,
	}
}

func (r *CsvRepo) Export() chan []string {
	rowsChan := make(chan []string, 100)

	go func() {
		defer close(rowsChan)

		// Open the CSV file
		file, err := os.Open(r.path)
		if err != nil {
			log.Error().Msgf("open csv [%s]: %v", r.path, err)
			return
		}
		defer file.Close()

		// Create a new CSV reader
		reader := csv.NewReader(file)
		if r.fields, err = reader.Read(); err != nil {
			log.Error().Msgf("read CSV [%s] header: %v", r.path, err)
			return
		}

		// Read and send each row through the channel
		for {
			row, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					return // Normal end of file
				}
				log.Error().Msgf("read CSV [%s] row: %v", r.path, err)
				return
			}
			rowsChan <- row
		}
	}()

	return rowsChan
}
