package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
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
		Name:     "fieldcounts",
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
		Name:     "fieldvalues",
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

	Commands = []*cli.Command{
		CreateCmd,
		PrintCmd,
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
		case "fieldcounts":
			if ctx.Bool("displayFieldMapping") {
				if proc.LoadFieldMapping(ctx.String("fieldMapping")) != nil {
					return fmt.Errorf("load field mapping error")
				}
				proc.MapFieldProperties()
			}
			return proc.PrintFieldCounts(ctx.Bool("sort"), ctx.Bool("greaterThanZeroOnly"))
		case "fieldvalues":
			return proc.PrintTopNFieldValues(ctx.String("name"), ctx.Int("maxNumLines"))
		}
	}
	return nil
}
