package main

import (
	"log"
	"os"

	commandshandler "github.com/abinashphulkonwar/rmgen/src/commands_handler"
	"github.com/urfave/cli/v2"
)

func main() {
	var file_path string
	var count int
	app := &cli.App{
		Name:  "rmgen",
		Usage: "generate ramdon data",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Usage:       "config file path",
				Aliases:     []string{"f"},
				Destination: &file_path,
			},
			&cli.IntFlag{
				Name:        "count",
				Usage:       "number of random data row",
				Aliases:     []string{"c"},
				Destination: &count,
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	commandshandler.Handler(file_path, count)
}
