package main

import (
	"fmt"
	"os"

	"github.com/ququzone/crypto-cli/cast-ext/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "cast-ext",
		Version: "v0.1.0",
		Authors: []*cli.Author{
			{
				Name:  "ququzone",
				Email: "xueping.yang@gmail.com",
			},
		},
		HelpName:  "cast-ext",
		Usage:     "extension for cast",
		UsageText: "cast-ext <SUBCOMMAND>",
		Commands:  commands.Commonds(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}
}
