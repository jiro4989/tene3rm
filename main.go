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
		os.Exit(int(exitcodeErrorParseArgs))
	}

	os.Exit(int(Main(args)))
}

func Main(args *CmdArgs) exitcode {
	fmt.Println(args)
	return exitcodeOK
}
