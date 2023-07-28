package main

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "JSON Utils",
		Usage:       "A collection of tools to work with JSON on the commandline alongside jq",
		Version:     "0.0.1",
		Description: "See readme at https://github.com/eaardal/jsonutils for details",
		Commands: []*cli.Command{
			ArrayCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	if os.Args[1] == "new-array" {
		stdout("[]")
		return
	}

	if os.Args[1] == "append-array-item" {
		arrayArg := os.Args[2]
		var arr []any
		if err := json.Unmarshal([]byte(arrayArg), &arr); err != nil {
			stderr("the array is invalid: %v. Array arg: %+v", err, arrayArg)
			return
		}
		nextItemArg := os.Args[3]
		var nextItem any
		if err := json.Unmarshal([]byte(nextItemArg), &nextItem); err != nil {
			stderr("the next item is invalid: %v. Next item arg: %+v", err, nextItemArg)
			return
		}
		arr = append(arr, nextItem)
		nextArr, err := json.Marshal(arr)
		if err != nil {
			stderr("failed to marshal next array: %v. Array contents: %+v", err, arr)
			return
		}
		stdout(string(nextArr))
		return
	}
}
