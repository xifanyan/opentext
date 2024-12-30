package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"github.com/xifanyan/adp"
	"github.com/xifanyan/opentext/discovery/r2a/data/arm"
	"github.com/xifanyan/opentext/discovery/r2a/dataproc"
)

var (
	PrintCmd = &cli.Command{
		Name: "print",
		Subcommands: []*cli.Command{
			PrintFieldCountsCmd,
			PrintFieldValuesCmd,
		},
	}

	PrintFieldCountsCmd = &cli.Command{
		Name:     "fieldCounts",
		Aliases:  []string{"fc"},
		Category: "print",
		Action:   execute,
		Flags: []cli.Flag{
			BasePath,
			DataSet,
			DisplayFieldMapping,
			FieldMapping,
			GreaterThanZeroOnly,
			Sort,
		},
	}

	PrintFieldValuesCmd = &cli.Command{
		Name:     "fieldValues",
		Aliases:  []string{"fv"},
		Category: "print",
		Action:   execute,
		Flags: []cli.Flag{
			BasePath,
			DataSet,
			Name,
			MaxNumLines,
		},
	}

	CreateCmd = &cli.Command{
		Name:   "create",
		Action: execute,
		Subcommands: []*cli.Command{
			CreateTicketsCmd,
		},
	}

	CreateTicketsCmd = &cli.Command{
		Name:     "tickets",
		Category: "create",
		Action:   execute,
		Flags: []cli.Flag{
			BasePath,
			DataSet,
			ApplicationID,
			Prefix,
			Template,
			Output,
		},
	}

	VerifyCmd = &cli.Command{
		Name:   "verify",
		Action: execute,
		Subcommands: []*cli.Command{
			VerifyFieldsCmd,
		},
	}

	VerifyFieldsCmd = &cli.Command{
		Name:     "fields",
		Category: "verify",
		Action:   execute,
		Flags: []cli.Flag{
			ApplicationID,
			FieldMapping,
		},
	}

	Commands = []*cli.Command{
		CreateCmd,
		PrintCmd,
		VerifyCmd,
	}
)

func execute(ctx *cli.Context) error {
	switch ctx.Command.Category {
	case "print":
		dataSet := arm.NewDataSetBuilder().WithBasePath(ctx.String("basePath")).WithID(ctx.String("dataSet")).Build()
		proc := dataproc.NewDataProc(dataSet)

		if err := proc.Initialize(); err != nil {
			return fmt.Errorf("initialize error: %v", err)
		}

		switch ctx.Command.Name {
		case "fieldCounts":
			if ctx.Bool("displayFieldMapping") {
				if proc.InitFieldMapping(ctx.String("fieldMapping")) != nil {
					return fmt.Errorf("load field mapping error")
				}
				proc.MapFieldProperties()
			}
			return proc.PrintFieldCounts(ctx.Bool("sort"), ctx.Bool("greaterThanZeroOnly"))
		case "fieldValues":
			return proc.PrintTopNFieldValues(ctx.String("name"), ctx.Int("maxNumLines"))
		}
	case "verify":
		var app = ctx.String("applicationID")

		client := adp.NewClientBuilder().WithDomain(ctx.String("domain")).
			WithPort(ctx.Int("port")).
			WithUser(ctx.String("user")).
			WithPassword(ctx.String("password")).
			Build()
		svc := adp.Service{ADPClient: client}

		switch ctx.Command.Name {
		case "fields":
			fieldProps, err := svc.GetFieldProperties(app)
			if err != nil {
				return err
			}

			mapping, err := dataproc.LoadFieldMapping(ctx.String("fieldMapping"))
			if err != nil {
				return err
			}

			for key, prop := range mapping {
				if prop.FieldType == "Text" {
					if _, ok := fieldProps[key]; !ok {
						log.Error().Msgf("field %s does not exist", key)
					}
				}
			}

		}

	}

	return nil
}
