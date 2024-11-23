package main

import "github.com/urfave/cli/v2"

var (
	User = &cli.StringFlag{
		Name:  "user",
		Usage: "ADP user name",
		Value: "adpuser",
	}

	Password = &cli.StringFlag{
		Name:  "password",
		Usage: "ADP user password",
	}

	Domain = &cli.StringFlag{
		Name:  "domain",
		Usage: "domain",
		Value: "vm-rhauswirth2.otxlab.net",
	}

	Port = &cli.IntFlag{
		Name:  "port",
		Usage: "port",
		Value: 8443,
	}

	Endpoint = &cli.StringFlag{
		Name:  "endpoint",
		Usage: "endpoint url",
	}

	Home = &cli.StringFlag{
		Name:  "home",
		Usage: "home dir",
	}

	MaxRunningDatasourcesPerHost = &cli.IntFlag{
		Name:  "maxRunningDatasourcesPerHost",
		Usage: "MAX_RUNNING_DATASOURCES_PER_CRAWLER_HOST",
		Value: 5,
	}

	MaxRunningDatasourcesPerEngine = &cli.IntFlag{
		Name:  "maxRunningDatasourcesPerEngine",
		Usage: "MAX_RUNNING_DATASOURCES_PER_ENGINE",
		Value: 2,
	}

	MaxDocumentCountEngine = &cli.IntFlag{
		Name:  "maxDocumentCountPerEngine",
		Usage: "MAX_DOCUMENT_COUNT_PER_ENGINE",
		Value: 4000000,
	}

	MaxNumberOfTicketsToBeProcessed = &cli.IntFlag{
		Name:  "maxNumberOfTicketsToBeProcessed",
		Usage: "MAX_NUMBER_OF_TICKETS_TO_BE_PROCESSED",
		Value: 20,
	}

	CrawlerHostPattern = &cli.StringFlag{
		Name:  "crawlerHostPattern",
		Usage: "CRAWLER_HOST_PATTERN",
	}

	Output = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Value:   "./",
		Usage:   "output folder path",
	}

	Name = &cli.StringFlag{
		Name:     "name",
		Aliases:  []string{"n"},
		Required: true,
		Usage:    "field name",
	}

	MaxNumLines = &cli.IntFlag{
		Name:    "maxNumLines",
		Aliases: []string{"l"},
		Value:   -1,
		Usage:   "max number of lines",
	}

	GreaterThanZeroOnly = &cli.BoolFlag{
		Name:    "greaterThanZeroOnly",
		Aliases: []string{"z"},
		Value:   false,
		Usage:   "greater than zero only",
	}

	Sort = &cli.BoolFlag{
		Name:    "sort",
		Aliases: []string{"s"},
		Value:   false,
		Usage:   "sort",
	}

	ApplicationID = &cli.StringFlag{
		Name:     "applicationID",
		Required: true,
		Usage:    "Application ID e.g., axcelerate.DataMigrationDemo",
	}

	Prefix = &cli.StringFlag{
		Name:  "prefix",
		Value: "tick",
		Usage: "Prefix for datasource e.g., tick",
	}

	Template = &cli.StringFlag{
		Name:  "template",
		Value: "_CSV_ARMLOAD_TEMPLATE",
		Usage: "datasource template name",
	}

	ImportConfig = &cli.StringFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Required: true,
		Usage:    "json config file",
	}

	BasePath = &cli.StringFlag{
		Name:     "basePath",
		Aliases:  []string{"base"},
		Required: true,
		Usage:    "base path",
	}

	DataSet = &cli.StringFlag{
		Name:     "dataSet",
		Aliases:  []string{"ds"},
		Required: true,
		Usage:    "dataset name",
	}

	Volumn = &cli.StringFlag{
		Name:    "volumn",
		Aliases: []string{"vol"},
		Usage:   "volumn name",
	}

	DisplayFieldMapping = &cli.BoolFlag{
		Name:    "displayFieldMapping",
		Aliases: []string{"m"},
		Value:   false,
		Usage:   "display field mapping",
	}

	FieldMapping = &cli.StringFlag{
		Name:  "fieldMapping",
		Usage: "field mapping file",
		Value: "field_mapping.json",
	}
)
