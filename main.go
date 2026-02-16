package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
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
		logger.Warn("action locked. please try again in one hour")
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
				l.Error("failed to increment fail count", "err", err)
			}
			continue
		}

		if state, err = uc.ResetFailCount(stateFile, state); err != nil {
			l.Error("failed to reset fail count", "err", err)
		}

		if err := remover.Remove(path); err != nil {
			l.Error("failed to remove a file", "err", err)
			return exitcodeErrorFailedToRemoveFile
		}
	}

	return exitcodeOK
}

// incrementFailCount はファイルの失敗回数を加算する。
func incrementFailCount() error {
	tmp := os.TempDir()
	basename := fmt.Sprintf(".%s.json", Appname)
	file := filepath.Join(tmp, basename)

	var data infra.StateDTO
	_, err := os.Stat(file)
	if os.IsExist(err) {
		// すでにファイルが存在する場合は読み取る
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &data); err != nil {
			return err
		}
		data.FailCount++
	}

	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	if err := os.WriteFile(file, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// resetFailCount はファイルの失敗回数を0にする。
func resetFailCount() error {
	tmp := os.TempDir()
	basename := fmt.Sprintf(".%s.json", Appname)
	file := filepath.Join(tmp, basename)

	var data infra.StateDTO
	_, err := os.Stat(file)
	if os.IsExist(err) {
		// すでにファイルが存在する場合は読み取る
		b, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &data); err != nil {
			return err
		}
	}
	data.FailCount = 0

	b, err := json.Marshal(&data)
	if err != nil {
		return err
	}
	if err := os.WriteFile(file, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func readData() (infra.StateDTO, error) {
	tmp := os.TempDir()
	basename := fmt.Sprintf(".%s.json", Appname)
	file := filepath.Join(tmp, basename)

	var data infra.StateDTO
	_, err := os.Stat(file)
	if os.IsExist(err) {
		// すでにファイルが存在する場合は読み取る
		b, err := os.ReadFile(file)
		if err != nil {
			return infra.StateDTO{}, err
		}
		if err := json.Unmarshal(b, &data); err != nil {
			return infra.StateDTO{}, err
		}
	}

	return data, nil
}
