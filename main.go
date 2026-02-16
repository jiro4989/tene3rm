package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/jiro4989/tene3rm/ui/terminal"
	"github.com/jiro4989/tene3rm/usecase"
)

type exitcode int

const (
	// Goがpanicしたときの終了コードは2なので、
	// 2と衝突しないように1桁の終了コードは使わない
	exitcodeOK             exitcode = 0
	exitcodeErrorParseArgs exitcode = 10 + iota
	exitcodeErrorFailedToRemoveFile
	exitcodeErrorFailedToCheckState
	exitcodeErrorFailedToLoadState
	exitcodeErrorActionLocked
)

func main() {
	args, err := ParseArgs()
	if err != nil {
		msg := fmt.Sprintf("parse args error: %v", err)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(int(exitcodeErrorParseArgs))
	}

	// シグナルで作業を中断出来なくする
	signal.Ignore(syscall.SIGTERM, syscall.SIGINT)

	seed := time.Now().Unix()
	os.Exit(int(Main(args, seed)))
}

func Main(args *CmdArgs, seed int64) exitcode {
	logger := newLogger(args.LogOutput)

	const stateFile = "." + Appname + ".json"

	fileRepo := infra.NewFileRepo(os.TempDir())
	timeImpl := infra.NewTimeGeneratorImpl()
	uc := usecase.NewStateUsecase(fileRepo, timeImpl)

	state, err := uc.LoadState(stateFile)
	if err != nil {
		logger.Error("failed to load state", "err", err)
		return exitcodeErrorFailedToLoadState
	}

	if ok, err := uc.IsActionLocked(stateFile, state); err != nil {
		logger.Error("failed to check state", "err", err)
		return exitcodeErrorFailedToCheckState
	} else if ok {
		fmt.Println(fmt.Sprintf("%s: action locked. please try again in one hour", Appname))
		return exitcodeErrorActionLocked
	}

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
			state, err = uc.IncrementFailCount(stateFile, state)
			if err != nil {
				// ステートの更新が失敗しても中断はしない
				l.Error("failed to increment fail count", "err", err)
			}
			continue
		}

		if state, err = uc.ResetFailCount(stateFile, state); err != nil {
			// ステートの更新が失敗しても中断はしない
			l.Error("failed to reset fail count", "err", err)
		}

		if err := remover.Remove(path); err != nil {
			l.Error("failed to remove a file", "err", err)
			return exitcodeErrorFailedToRemoveFile
		}
	}

	return exitcodeOK
}
