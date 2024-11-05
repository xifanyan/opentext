package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sourcegraph/conc/pool"
	"github.com/xifanyan/opentext/discovery/data2sqlite"
)

func main() {

	// Create multiple CsvRepo instances
	repos := []*data2sqlite.CsvRepo{
		data2sqlite.NewCsvRepo("../../../testdata/Electric_Vehicle_Population_Data.csv"),
		data2sqlite.NewCsvRepo("../../../testdata/Crime_Data_from_2020_to_Present.csv"),
	}

	// Create an output channel to receive the results
	outputChan := make(chan []string)

	// Create a pool with a limited size
	p := pool.New().WithMaxGoroutines(5)

	// Start a goroutine for each repo to read and send to the output channel
	for _, repo := range repos {
		p.Go(func() {
			rowsChan := repo.Export()
			for row := range rowsChan {
				outputChan <- row
			}
		})
	}

	// Wait for the pool to finish
	go func() {
		p.Wait()
		close(outputChan)
	}()

	// Read from the output channel and log the results
	for row := range outputChan {
		log.Info().Msgf("%v", row)
	}

	log.Info().Msgf("done")
}
