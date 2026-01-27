package main

import (
	"log/slog"
	"os"
)

const (
	logOutputStdout = "stdout"
	logOutputStderr = "stderr"
)

func newLogger(output string) *slog.Logger {
	out := os.Stderr
	if output == logOutputStdout {
		out = os.Stdout
	}
	l := slog.New(slog.NewTextHandler(out, nil))
	return l
}
