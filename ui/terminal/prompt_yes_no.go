package terminal

import (
	"fmt"

	"github.com/jiro4989/tene3rm/usecase"
	"github.com/manifoldco/promptui"
)

// promptWithYesNo はシンプルなYes/Noプロンプトを表示する。
func promptWithYesNo(path string) (bool, error) {
	uc := usecase.NewYesNoUsecase()
	return promptWithSimpleText(path, appname+": remove file '%s'? [y/n]", "", uc.JudgeYesNo)
}

// promptWithYesNoDenial はNoのときだけファイルを削除するプロンプトを表示する。
func promptWithYesNoDenial(path string) (bool, error) {
	uc := usecase.NewYesNoUsecase()
	return promptWithSimpleText(path, appname+": DON't remove file '%s'? [y/n]", "", uc.JudgeYesNoDenial)
}

// promptWithYesNoJapanese はシンプルなはい/いいえプロンプトを表示する。
func promptWithYesNoJapanese(path string) (bool, error) {
	uc := usecase.NewYesNoUsecase()
	return promptWithSimpleText(path, appname+": '%s' ファイルを削除しますか? [はい/いいえ]", "", uc.JudgeYesNoJapanese)
}

// promptWithYesNoJapanese3 は 3 回確認するプロンプトを表示する。
func promptWithYesNoJapanese3(path string) (bool, error) {
	uc := usecase.NewYesNoUsecase()
	ok, err := promptWithYesNoJapanese(path)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	ok, err = promptWithSimpleText(path, appname+": '%s' 本当に? [はい/いいえ]", "", uc.JudgeYesNoJapanese)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	ok, err = promptWithSimpleText(path, appname+": '%s' 削除すると復元できなくなるけれど大丈夫? [はい/いいえ]", "いいえ", uc.JudgeYesNoJapanese)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	return true, nil
}

func promptWithSimpleText(path string, promptFmt string, defaultValue string, f func(string) bool) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf(promptFmt, path),
		Validate: validate,
		Default:  defaultValue,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	return f(result), nil
}
