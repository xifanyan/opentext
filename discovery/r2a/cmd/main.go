package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "r2a",
		Version: "0.1.0-alpha",
		Usage:   "Relativity To Axcelerate",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "json config file",
				EnvVars: []string{"R2A_CONFIG"},
				Value:   ".r2a.json",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Debug Mode",
				Value:   false,
			},
		},
		Commands: Commands,
		Before: func(c *cli.Context) error {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
			if c.Bool("debug") {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano})
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error().Msgf("error: %s", err)
		os.Exit(1)
	}
}
