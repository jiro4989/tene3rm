package main

import (
	"fmt"
	"os"
)

type exitcode int

const (
	exitcodeOK             exitcode = 0
	exitcodeErrorParseArgs exitcode = 11 + iota
)

func main() {
	args, err := ParseArgs()
	if err != nil {
		msg := fmt.Sprintf("parse args error: %v", err)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(int(exitcodeErrorParseArgs))
	}

	os.Exit(int(Main(args)))
}

func Main(args *CmdArgs) exitcode {
	logger := newLogger(args.LogOutput)
	logger.Info("hello")
	return exitcodeOK
}
