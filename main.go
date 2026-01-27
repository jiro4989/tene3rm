package main

import (
	"fmt"
	"os"
)

type exitcode int

const (
	// Goがpanicしたときの終了コードは2なので、
	// 2と衝突しないように1桁の終了コードは使わない
	exitcodeOK             exitcode = 0
	exitcodeErrorParseArgs exitcode = 10 + iota
	exitcodeErrorFailedToRemoveFile
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

	var remover Remover = &OSRemover{}
	if args.DryRun {
		remover = &NilRemover{}
	}

	for _, path := range args.Args {
		logger := logger.With("path", path)

		if err := remover.Remove(path); err != nil {
			logger.With("err", err).Error("failed to remove a file")
			return exitcodeErrorFailedToRemoveFile
		}
	}

	return exitcodeOK
}
