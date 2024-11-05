package data2sqlite

import (
	"encoding/csv"
	"os"
)

func newCsvReader(f *os.File) dataReader {
	return csv.NewReader(f)
}

type CsvData struct {
	Data
}

func NewCsvData(path string) *CsvData {
	return &CsvData{
		Data: Data{
			path: path,
		},
	}
}

func (c *CsvData) Header() ([]string, error) {
	return readHeader(&c.Data, newCsvReader)
}

func (c *CsvData) Export() chan []string {
	return exportData(&c.Data, newCsvReader)
}
