package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// promptWithYesNo はシンプルなYes/Noプロンプトを表示する。
func promptWithYesNo(path string) (bool, error) {
	return promptWithSimpleText(path, "(*'-')! < remove file '%s'? [y/n]", []string{"y", "ye", "yes"})
}

// promptWithYesNoDenial はNoのときだけファイルを削除するプロンプトを表示する。
func promptWithYesNoDenial(path string) (bool, error) {
	return promptWithSimpleText(path, "(*'-')! < DON't remove file '%s'? [y/n]", []string{"n", "no"})
}

// promptWithYesNoInJapanese はシンプルなはい/いいえプロンプトを表示する。
func promptWithYesNoInJapanese(path string) (bool, error) {
	return promptWithSimpleText(path, "(*'-')! < '%s' ファイルを削除しますか? [はい/いいえ]", []string{"は", "はい"})
}

func promptWithSimpleText(path string, promptFmt string, wants []string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf(promptFmt, path),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	result = strings.TrimSpace(result)
	found := false
	for _, want := range wants {
		if result == want {
			found = true
			break
		}
	}

	return found, nil
}
