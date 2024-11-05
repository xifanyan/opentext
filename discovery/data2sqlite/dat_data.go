package data2sqlite

import (
	"os"

	"github.com/xifanyan/opentext/discovery/dat"
)

func newDatReader(f *os.File) dataReader {
	return dat.NewReader(f)
}

type DatData struct {
	Data
}

func NewDatData(path string) *DatData {
	return &DatData{
		Data: Data{
			path: path,
		},
	}
}

func (c *DatData) Header() ([]string, error) {
	return readHeader(&c.Data, newDatReader)
}

func (c *DatData) Export() chan []string {
	return exportData(&c.Data, newDatReader)
}
