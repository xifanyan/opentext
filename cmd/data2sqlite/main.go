package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sourcegraph/conc/pool"
	"github.com/xifanyan/opentext/discovery/data2sqlite"
)

func main() {

	// Create multiple CsvData instances
	dataRepos := []*data2sqlite.CsvData{
		data2sqlite.NewCsvData("../../testdata/data2sqlite/Electric_Vehicle_Population_Data.csv"),
		data2sqlite.NewCsvData("../../testdata/data2sqlite/Crime_Data_from_2020_to_Present.csv"),
	}

	// Create an output channel to receive the results
	outputChan := make(chan []string)

	/*
		// Create a pool with a limited size

		for _, repo := range dataRepos {
			p.Go(func() {
				if err := repo.GetHeader(); err != nil {
					log.Error().Msgf("get header: %v", err)
				}
				fmt.Printf("%v\n", repo.Header)
			})
		}

		p.Wait()
	*/

	p := pool.New().WithMaxGoroutines(5)

	// Start a goroutine for each repo to read and send to the output channel
	for _, repo := range dataRepos {
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
