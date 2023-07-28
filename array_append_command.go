package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
)

var ArrayAppendCommand = &cli.Command{
	Name:  "append",
	Usage: "Appends the given input to the given array",
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
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
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

		typeArg := c.String("type")
		if err := validateTypeArg(typeArg); err != nil {
			return fmt.Errorf("argument --type is invalid: %v", err)
		}

		var array []any
		if err := json.Unmarshal([]byte(arrayArg), &array); err != nil {
			return fmt.Errorf("the array is invalid: %v. Array arg: %+v", err, arrayArg)
		}

		value, err := convertValueToTypeIfNeeded(valueArg, typeArg)
		if err != nil {
			return err
		}

		if strings.HasPrefix(valueArg, "{") || strings.HasPrefix(valueArg, "[") {
			if err := json.Unmarshal([]byte(valueArg), &value); err != nil {
				return fmt.Errorf("the --value arg is invalid JSON: %v. Value arg: %+v", err, valueArg)
			}
		}

		array = append(array, value)

		arrayJson, err := json.Marshal(array)
		if err != nil {
			return fmt.Errorf("failed to marshal updated array to JSON: %v. Updated array: %+v", err, array)
		}

		stdout(string(arrayJson))
		return nil
	},
}

func convertValueToTypeIfNeeded(value string, toType string) (any, error) {
	if toType == "" {
		return value, nil
	}

	if toType == "number" {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value %q to number: %v", value, err)
		}
		return num, nil
	}

	if toType == "bool" {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value %q to bool: %v", value, err)
		}
		return b, nil
	}

	return nil, fmt.Errorf("unhandled conversion case: type %s is not supported", value)
}

func validateTypeArg(val string) error {
	if val == "" {
		return nil
	}

	if val != "number" && val != "bool" {
		return fmt.Errorf("must be either 'number' or 'bool' but was '%s'", val)
	}

	return nil
}
