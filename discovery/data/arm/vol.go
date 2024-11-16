package arm

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/xifanyan/opentext/discovery/utils"
)

func (v *Vol) Path() string {
	return filepath.Join(v.basePath, v.armID, v.ID, v.DatFile)
}

// ReadDatHeader reads the header of a .dat file and returns it as a slice of strings.
// If the header has already been read, it is returned from the cache.
// If the header cannot be read, an error is returned.
// The header is cached for further calls to ReadDatHeader.
func (v *Vol) ReadDatHeader() ([]string, error) {
	var err error

	if v.header != nil {
		return v.header, nil
	}

	path := v.Path()
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

// StreamDatToChannel reads the .dat file and streams each row to the returned channel.
// The channel is closed when the end of the file is reached.
// If an error occurs while reading the file, the channel is closed and the error is logged.
// The values of each column are zipped with the header fields, and the result is sent through the channel.
func (v *Vol) StreamDatToChannel() chan map[string]string {
	rowsChan := make(chan map[string]string, 100)

	go func() {
		defer close(rowsChan)

		path := v.Path()
		// Open the file
		f, err := os.Open(path)
		if err != nil {
			log.Error().Msgf("open [%s]: %v", path, err)
			return
		}
		defer f.Close()

		// Create a new reader
		reader := NewReader(f)

		// pull the header
		header, err := reader.Read()
		if err != nil {
			log.Error().Msgf("read header [%s]: %v", path, err)
			return
		}

		// Read and send each row through the channel
		for {
			valules, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					return // Normal end of file
				}
				log.Error().Msgf("read row [%s]: %v", path, err)
				return
			}

			row, err := utils.Zip(header, valules)
			if err != nil {
				log.Error().Msgf("zip row [%s]: %v", path, err)
			}
			rowsChan <- row
		}
	}()

	return rowsChan
}
