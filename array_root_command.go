package main

import "github.com/urfave/cli/v2"

var ArrayCommand = &cli.Command{
	Name:  "array",
	Usage: "Utils for working with JSON arrays",
	Subcommands: []*cli.Command{
		ArrayNewCommand,
		ArrayAppendCommand,
		ArraySpreadCommand,
	},
}
