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
				Name:    "user",
				Usage:   "adp user",
				EnvVars: []string{"ADPUSER"},
				Value:   "adpuser",
			},
			&cli.StringFlag{
				Name:    "password",
				Usage:   "adp user password",
				EnvVars: []string{"ADPUSERPASSWORD"},
			},
			&cli.StringFlag{
				Name:  "domain",
				Usage: "domain",
				Value: "localhost",
			},
			&cli.IntFlag{
				Name:  "port",
				Usage: "port",
				Value: 8443,
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
