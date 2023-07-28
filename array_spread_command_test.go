package main

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestArraySpreadCommand_WhenArrayArgIsNotSet_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArraySpreadCommand.Run(ctx, "spread", "--value", "[\"hei\"]" /* --array not set = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Required flag \"array\" not set")
}

func TestArraySpreadCommand_WhenArrayArgIsEmpty_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArraySpreadCommand.Run(ctx, "spread", "--value", "[\"hei\"]", "--array", "" /* no value for --array = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "argument --array is empty")
}

func TestArraySpreadCommand_WhenValueArgIsNotSet_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArraySpreadCommand.Run(ctx, "spread", "--array", "[]" /* --value not set = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Required flag \"value\" not set")
}

func TestArraySpreadCommand_WhenValueArgIsEmpty_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArraySpreadCommand.Run(ctx, "spread", "--array", "[]", "--value", "" /* no value for --value = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "argument --value is empty")
}

func TestArraySpreadCommand_ArrayArgInvalidJSON(t *testing.T) {
	testCases := []struct {
		name     string
		arrayArg string
	}{
		{
			name:     "when array is text return error",
			arrayArg: "some text",
		},
		{
			name:     "when array is number as string return error",
			arrayArg: "1",
		},
		{
			name:     "when array is JSON object as string return error",
			arrayArg: "{ \"foo\": \"bar\" }",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			app := cli.NewApp()
			flags := flag.NewFlagSet("test", flag.ExitOnError)
			flags.String("value", "", "")
			flags.String("array", "", "")

			parent := &cli.Context{}
			ctx := cli.NewContext(app, flags, parent)

			err := ArraySpreadCommand.Run(ctx, "spread", "--array", testCase.arrayArg, "--value", "foo")

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, "the array is invalid")
		})
	}
}

func TestArraySpreadCommand_ValueArgIsNotArray_ReturnsError(t *testing.T) {
	testCases := []struct {
		name     string
		valueArg string
	}{
		{
			name:     "when value is text return error",
			valueArg: "some text",
		},
		{
			name:     "when value is number as string return error",
			valueArg: "1",
		},
		{
			name:     "when value is JSON object as string return error",
			valueArg: "{ \"foo\": \"bar\" }",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			app := cli.NewApp()
			flags := flag.NewFlagSet("test", flag.ExitOnError)
			flags.String("value", "", "")
			flags.String("array", "", "")

			parent := &cli.Context{}
			ctx := cli.NewContext(app, flags, parent)

			err := ArraySpreadCommand.Run(ctx, "spread", "--array", "[]", "--value", testCase.valueArg)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, "the value is invalid: It must be a JSON array")
		})
	}
}

func TestArraySpreadCommand_ValueArg(t *testing.T) {
	testCases := []struct {
		name           string
		arrayArg       string
		valueArg       string
		expectedStdout string
	}{
		{
			name:           "when value is text append value to array as string",
			arrayArg:       "[]",
			valueArg:       "[\"some text\"]",
			expectedStdout: "[\"some text\"]",
		},
		{
			name:           "when array already has one item, append value to the existing array",
			arrayArg:       "[\"some text\"]",
			valueArg:       "[\"more text\"]",
			expectedStdout: "[\"some text\",\"more text\"]",
		},
		{
			name:           "when array already has one item and value contains many values, append all values to the array",
			arrayArg:       "[\"some text\"]",
			valueArg:       "[\"more text\", \"even more\"]",
			expectedStdout: "[\"some text\",\"more text\",\"even more\"]",
		},
		{
			name:           "when value is number append value to array",
			arrayArg:       "[]",
			valueArg:       "[1]",
			expectedStdout: "[1]",
		},
		{
			name:           "when value is bool append value to array",
			arrayArg:       "[]",
			valueArg:       "[true]",
			expectedStdout: "[true]",
		},
		{
			name:           "when value is JSON object as string append value to array as string",
			arrayArg:       "[]",
			valueArg:       "[{ \"foo\": \"bar\" }]",
			expectedStdout: "[{\"foo\":\"bar\"}]",
		},
		{
			name:           "when array contains a JSON object and value is another JSON object as string append value to array as string",
			arrayArg:       "[{ \"aaa\": \"bbb\" }]",
			valueArg:       "[{ \"ccc\": \"ddd\" }]",
			expectedStdout: "[{\"aaa\":\"bbb\"},{\"ccc\":\"ddd\"}]",
		},
		{
			name:           "when value is JSON array as string append value to array as string",
			arrayArg:       "[]",
			valueArg:       "[[\"foo\", \"bar\"]]",
			expectedStdout: "[[\"foo\",\"bar\"]]",
		},
		{
			name:           "when array already contains an array and value is a JSON array as string append value to array as string",
			arrayArg:       "[[\"aaa\", \"bbb\"]]",
			valueArg:       "[[\"ccc\", \"ddd\"]]",
			expectedStdout: "[[\"aaa\",\"bbb\"],[\"ccc\",\"ddd\"]]",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			app := cli.NewApp()
			flags := flag.NewFlagSet("test", flag.ExitOnError)
			flags.String("value", "", "")
			flags.String("array", "", "")

			parent := &cli.Context{}
			ctx := cli.NewContext(app, flags, parent)

			stdout := &bytes.Buffer{}
			stdoutWriter = stdout

			err := ArraySpreadCommand.Run(ctx, "spread", "--array", testCase.arrayArg, "--value", testCase.valueArg)

			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedStdout, stdout.String())
		})
	}
}
