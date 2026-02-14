package terminal

import (
	"fmt"

	"github.com/jiro4989/tene3rm/usecase"
	"github.com/manifoldco/promptui"
)

const face = "(*'-')!"

// promptWithYesNo はシンプルなYes/Noプロンプトを表示する。
func promptWithYesNo(path string) (bool, error) {
	svc := usecase.NewYesNoService()
	return promptWithSimpleText(path, face+" < remove file '%s'? [y/n]", "", svc.JudgeYesNo)
}

// promptWithYesNoDenial はNoのときだけファイルを削除するプロンプトを表示する。
func promptWithYesNoDenial(path string) (bool, error) {
	svc := usecase.NewYesNoService()
	return promptWithSimpleText(path, face+" < DON't remove file '%s'? [y/n]", "", svc.JudgeYesNoDenial)
}

// promptWithYesNoJapanese はシンプルなはい/いいえプロンプトを表示する。
func promptWithYesNoJapanese(path string) (bool, error) {
	svc := usecase.NewYesNoService()
	return promptWithSimpleText(path, face+" < '%s' ファイルを削除しますか? [はい/いいえ]", "", svc.JudgeYesNoJapanese)
}

// promptWithYesNoJapanese3 は 3 回確認するプロンプトを表示する。
func promptWithYesNoJapanese3(path string) (bool, error) {
	svc := usecase.NewYesNoService()
	ok, err := promptWithYesNoJapanese(path)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	ok, err = promptWithSimpleText(path, "(*'o')? < '%s' 本当に? [はい/いいえ]", "", svc.JudgeYesNoJapanese)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	ok, err = promptWithSimpleText(path, "(*-o-)? < '%s' 削除すると復元できなくなるけれど大丈夫? [はい/いいえ]", "いいえ", svc.JudgeYesNoJapanese)
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
