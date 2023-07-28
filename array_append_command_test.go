package main

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestArrayAppendCommand_WhenArrayArgIsNotSet_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArrayAppendCommand.Run(ctx, "append", "--value", "hei" /* --array not set = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Required flag \"array\" not set")
}

func TestArrayAppendCommand_WhenArrayArgIsEmpty_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArrayAppendCommand.Run(ctx, "append", "--value", "hei", "--array", "" /* no value for --array = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "argument --array is empty")
}

func TestArrayAppendCommand_WhenValueArgIsNotSet_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArrayAppendCommand.Run(ctx, "append", "--array", "[]" /* --value not set = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Required flag \"value\" not set")
}

func TestArrayAppendCommand_WhenValueArgIsEmpty_ReturnsError(t *testing.T) {
	app := cli.NewApp()
	flags := flag.NewFlagSet("test", flag.ExitOnError)
	flags.String("value", "", "")
	flags.String("array", "", "")

	parent := &cli.Context{}
	ctx := cli.NewContext(app, flags, parent)

	err := ArrayAppendCommand.Run(ctx, "append", "--array", "[]", "--value", "" /* no value for --value = error */)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "argument --value is empty")
}

func TestArrayAppendCommand_ArrayArgInvalidJSON(t *testing.T) {
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

			err := ArrayAppendCommand.Run(ctx, "append", "--array", testCase.arrayArg, "--value", "foo")

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, "the array is invalid")
		})
	}
}

func TestArrayAppendCommand_ValueArg(t *testing.T) {
	testCases := []struct {
		name           string
		valueArg       string
		expectedStdout string
	}{
		{
			name:           "when value is text append value to array as string",
			valueArg:       "some text",
			expectedStdout: "[\"some text\"]",
		},
		{
			name:           "when value is number as string append value to array as string",
			valueArg:       "1",
			expectedStdout: "[\"1\"]",
		},
		{
			name:           "when value is bool as string append value to array as string",
			valueArg:       "true",
			expectedStdout: "[\"true\"]",
		},
		{
			name:           "when value is JSON object as string append value to array as string",
			valueArg:       "{ \"foo\": \"bar\" }",
			expectedStdout: "[{\"foo\":\"bar\"}]",
		},
		{
			name:           "when value is JSON array as string append value to array as string",
			valueArg:       "[\"foo\", \"bar\"]",
			expectedStdout: "[[\"foo\",\"bar\"]]",
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

			err := ArrayAppendCommand.Run(ctx, "append", "--array", "[]", "--value", testCase.valueArg)

			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedStdout, stdout.String())
		})
	}
}

func TestArrayAppendCommand_TypeArg(t *testing.T) {
	testCases := []struct {
		name           string
		valueArg       string
		typeArg        string
		expectedStdout string
	}{
		{
			name:           "when value is number as string and type is number append value to array as number",
			valueArg:       "1",
			typeArg:        "number",
			expectedStdout: "[1]",
		},
		{
			name:           "when value is bool (true) as string and type is bool append value to array as bool",
			valueArg:       "true",
			typeArg:        "bool",
			expectedStdout: "[true]",
		},
		{
			name:           "when value is bool (false) as string and type is bool append value to array as bool",
			valueArg:       "false",
			typeArg:        "bool",
			expectedStdout: "[false]",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			app := cli.NewApp()
			flags := flag.NewFlagSet("test", flag.ExitOnError)
			flags.String("value", "", "")
			flags.String("array", "", "")
			flags.String("type", "", "")

			parent := &cli.Context{}
			ctx := cli.NewContext(app, flags, parent)

			stdout := &bytes.Buffer{}
			stdoutWriter = stdout

			err := ArrayAppendCommand.Run(ctx, "append", "--array", "[]", "--value", testCase.valueArg, "--type", testCase.typeArg)

			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedStdout, stdout.String())
		})
	}
}
