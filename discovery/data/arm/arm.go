package arm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ARM struct {
	ID       string
	basePath string
	Vols     []Vol
}

type Vol struct {
	basePath string

	armID   string
	ID      string
	OptFile string
	DatFile string

	header []string
}

type ARMBuilder struct {
	*ARM
}

func NewArmDBBuilder() *ARMBuilder {
	return &ARMBuilder{
		ARM: &ARM{},
	}
}

func (b *ARMBuilder) WithID(id string) *ARMBuilder {
	b.ARM.ID = id
	return b
}

func (b *ARMBuilder) WithBasePath(basePath string) *ARMBuilder {
	b.ARM.basePath = basePath
	return b
}

func (b *ARMBuilder) Build() *ARM {
	return b.ARM
}

// Initialize reads the file system under the given basePath and ID, and
// populates the ARM's Vols field with Vol structs. Each Vol is populated with
// the ARM's basePath and ID, and the file name and opt file name found in the
// file system.
func (a *ARM) Initialize() error {
	rootPath := fmt.Sprintf("%s/%s", a.basePath, a.ID)
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if directory
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".dat") {
			return nil
		}

		// Get relative path from root
		relPath, err := filepath.Rel(a.basePath, path)
		if err != nil {
			return err
		}

		// Split path components
		parts := strings.Split(relPath, string(os.PathSeparator))
		if len(parts) != 3 {
			return nil
		}

		vol := Vol{
			basePath: a.basePath,
			armID:    a.ID,
			ID:       parts[1],
			DatFile:  parts[2],
			OptFile:  getOptFileByDatFile(path),
		}

		a.Vols = append(a.Vols, vol)

		return nil
	})

	return err
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
func getOptFileByDatFile(path string) string {
	dir, fileName := filepath.Split(path)
	baseName := strings.TrimSuffix(fileName, ".dat")

	optFilePath := filepath.Join(dir, baseName+".opt")
	if _, err := os.Stat(optFilePath); err != nil {
		return ""
	}

	return filepath.Base(optFilePath)
}
