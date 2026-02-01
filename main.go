package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jiro4989/tene3rm/ui/terminal"
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

	seed := time.Now().Unix()
	os.Exit(int(Main(args, seed)))
}

func Main(args *CmdArgs, seed int64) exitcode {
	logger := newLogger(args.LogOutput)

	var remover Remover = &OSRemover{}
	if args.DryRun {
		remover = &NilRemover{}
	}

	for _, path := range args.Args {
		l := logger.With("path", path)

		ok, err := terminal.Prompt(path, seed)
		if err != nil {
			l.Error("failed to check", "err", err)
			return exitcodeErrorFailedToRemoveFile
		}
		if !ok {
			continue
		}

		if err := remover.Remove(path); err != nil {
			l.Error("failed to remove a file", "err", err)
			return exitcodeErrorFailedToRemoveFile
		}
	}

	return exitcodeOK
}
