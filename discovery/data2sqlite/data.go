package data2sqlite

import (
	"os"

	"github.com/rs/zerolog/log"
)

type Data struct {
	header []string
	path   string
}

type dataReader interface {
	Read() ([]string, error)
}

// readHeader reads the header of a file and caches it. It returns the header on
// subsequent calls.
func readHeader(r *Data, newReader func(*os.File) dataReader) ([]string, error) {
	var err error

	if r.header != nil {
		return r.header, nil
	}

	f, err := os.Open(r.path)
	if err != nil {
		log.Error().Msgf("open [%s]: %v", r.path, err)
		return nil, err
	}
	defer f.Close()

	// Create a new reader
	reader := newReader(f)

	// pull the header
	r.header, err = reader.Read()
	if err != nil {
		log.Error().Msgf("read header [%s]: %v", r.path, err)
		return nil, err
	}

	return r.header, err
}

// exportData opens a file and reads it row by row, sending each row through a
// channel. The function takes a Data object and a newReader function that takes
// an open file and returns a dataReader interface. The function returns a channel
// that receives each row as it is read, or an error if there is a problem.
func exportData(r *Data, newReader func(*os.File) dataReader) chan []string {
	rowsChan := make(chan []string, 100)

	go func() {
		defer close(rowsChan)

		// Open the file
		f, err := os.Open(r.path)
		if err != nil {
			log.Error().Msgf("open [%s]: %v", r.path, err)
			return
		}
		defer f.Close()

		// Create a new reader
		reader := newReader(f)

		// pull the header
		r.header, err = reader.Read()
		if err != nil {
			log.Error().Msgf("read header [%s]: %v", r.path, err)
			return
		}

		// Read and send each row through the channel
		for {
			row, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					return // Normal end of file
				}
				log.Error().Msgf("read row [%s]: %v", r.path, err)
				return
			}
			rowsChan <- row
		}
	}()

	return rowsChan
}
