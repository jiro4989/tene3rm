package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// promptWithYesNo はシンプルなYes/Noプロンプトを表示する。
func promptWithYesNo(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? [y/n]", appname, path),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	result = strings.TrimSpace(result)
	switch result {
	case "y", "ye", "yes":
		return true, nil
	}

	return false, nil
}
