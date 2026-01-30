package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// promptWithYesNoDenial はNoのときだけファイルを削除するプロンプトを表示する。
func promptWithYesNoDenial(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: DON'T remove file '%s'? [y/n]", appname, path),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	result = strings.TrimSpace(result)
	switch result {
	case "n", "no":
		return true, nil
	}

	return false, nil
}
