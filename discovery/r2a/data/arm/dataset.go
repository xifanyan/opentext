package arm

import (
	"os"
	"path/filepath"
	"strings"
)

type DataSet struct {
	ID       string
	BasePath string
	Volumns  []Volumn
}

type Volumn struct {
	BasePath string

	DataSetID string
	ID        string
	OptFile   string
	DatFile   string

	header []string
}

type DataSetBuilder struct {
	*DataSet
}

func NewDataSetBuilder() *DataSetBuilder {
	return &DataSetBuilder{
		DataSet: &DataSet{},
	}
}

func (b *DataSetBuilder) WithID(id string) *DataSetBuilder {
	b.DataSet.ID = id
	return b
}

func (b *DataSetBuilder) WithBasePath(basePath string) *DataSetBuilder {
	b.DataSet.BasePath = basePath
	return b
}

func (b *DataSetBuilder) Build() *DataSet {
	return b.DataSet
}

/*
GetOptFileByDatFile gets an opt file name by a given .dat file path.
The opt file is assumed to be in the same directory as the .dat file, and
with the same name but with a .opt extension.

param:

	path - the path to the .dat file.

return:

	the name of the corresponding .opt file, or an empty string if
	the file does not exist.
*/
func GetOptFileByDatFile(path string) string {
	dir, fileName := filepath.Split(path)
	baseName := strings.TrimSuffix(fileName, ".dat")

	optFilePath := filepath.Join(dir, baseName+".opt")
	if _, err := os.Stat(optFilePath); err != nil {
		return ""
	}

	return filepath.Base(optFilePath)
}
