package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// promptWithYesNoInJapanese はシンプルなはい/いいえプロンプトを表示する。
func promptWithYesNoInJapanese(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: '%s' ファイルを削除しますか? [はい/いいえ]", appname, path),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	result = strings.TrimSpace(result)
	switch result {
	case "は", "はい":
		return true, nil
	}

	return false, nil
}
