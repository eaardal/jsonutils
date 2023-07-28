package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

var ArraySpreadCommand = &cli.Command{
	Name:  "spread",
	Usage: "Spreads the given input array to the given array",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "array",
			Aliases:  []string{"a"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "value",
			Aliases:  []string{"v"},
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		arrayArg := c.String("array")
		if arrayArg == "" {
			return fmt.Errorf("argument --array is empty")
		}

		valueArg := c.String("value")
		if valueArg == "" {
			return fmt.Errorf("argument --value is empty")
		}

		var array []any
		if err := json.Unmarshal([]byte(arrayArg), &array); err != nil {
			return fmt.Errorf("the array is invalid: %v. Array arg: %+v", err, arrayArg)
		}

		if !strings.HasPrefix(valueArg, "[") {
			return fmt.Errorf("the value is invalid: It must be a JSON array so the first character should be '[' but it was %s. Value arg: %q", valueArg[:1], valueArg)
		}

		var value []any
		if err := json.Unmarshal([]byte(valueArg), &value); err != nil {
			return fmt.Errorf("the --value arg is invalid JSON: %v. Value arg: %+v", err, valueArg)
		}

		array = append(array, value...)

		arrayJson, err := json.Marshal(array)
		if err != nil {
			return fmt.Errorf("failed to marshal updated array to JSON: %v. Updated array: %+v", err, array)
		}

		stdout(string(arrayJson))
		return nil
	},
}
