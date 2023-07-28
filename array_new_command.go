package main

import "github.com/urfave/cli/v2"

var ArrayNewCommand = &cli.Command{
	Name:  "new",
	Usage: "Returns an empty JSON array",
	Action: func(c *cli.Context) error {
		stdout("[]")
		return nil
	},
}
