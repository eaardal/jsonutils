package main

import (
	"fmt"
	"io"
	"os"
)

var stdoutWriter io.Writer
var stderrWriter io.Writer

func stdout(format string, a ...any) {
	if stdoutWriter == nil {
		stdoutWriter = os.Stdout
	}
	_, _ = fmt.Fprintf(stdoutWriter, fmt.Sprintf(format, a...))
}

func stderr(format string, a ...any) {
	if stderrWriter == nil {
		stderrWriter = os.Stderr
	}
	_, _ = fmt.Fprintf(stderrWriter, fmt.Sprintf(format, a...))
}
