package main

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestArrayNewCommand_ReturnsEmptyArrayString(t *testing.T) {
	app := cli.NewApp()
	flags := &flag.FlagSet{}
	ctx := cli.NewContext(app, flags, nil)

	stdout := &bytes.Buffer{}
	stdoutWriter = stdout

	err := ArrayNewCommand.Run(ctx)

	assert.Nil(t, err)
	assert.Equal(t, "[]", stdout.String())
}
