package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

const face = "(*'-')!"

// promptWithYesNo はシンプルなYes/Noプロンプトを表示する。
func promptWithYesNo(path string) (bool, error) {
	return promptWithSimpleText(path, face+" < remove file '%s'? [y/n]", "yes", "")
}

// promptWithYesNoDenial はNoのときだけファイルを削除するプロンプトを表示する。
func promptWithYesNoDenial(path string) (bool, error) {
	return promptWithSimpleText(path, face+" < DON't remove file '%s'? [y/n]", "no", "")
}

// promptWithYesNoInJapanese はシンプルなはい/いいえプロンプトを表示する。
func promptWithYesNoInJapanese(path string) (bool, error) {
	return promptWithSimpleText(path, face+" < '%s' ファイルを削除しますか? [はい/いいえ]", "はい", "")
}

// promptWithYesNoInJapanese3 は 3 回確認するプロンプトを表示する。
func promptWithYesNoInJapanese3(path string) (bool, error) {
	ok, err := promptWithYesNoInJapanese(path)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	ok, err = promptWithSimpleText(path, "(*'o')? < '%s' 本当に? [はい/いいえ]", "はい", "")
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	ok, err = promptWithSimpleText(path, "(*-o-)? < '%s' 削除すると復元できなくなるけれど大丈夫? [はい/いいえ]", "はい", "いいえ")
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	return true, nil
}

func promptWithSimpleText(path string, promptFmt string, want string, defaultValue string) (bool, error) {
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

	result = strings.TrimSpace(result)
	found := false
	wants := prefixes(want)
	for _, w := range wants {
		if result == w {
			found = true
			break
		}
	}

	return found, nil
}

// prefixes は yes を y, ye, yes といった感じの配列にして返す。
func prefixes(s string) []string {
	// マルチバイト文字を考慮するため rune に変換
	runes := []rune(s)
	result := make([]string, 0, len(runes))
	for i := 1; i <= len(runes); i++ {
		result = append(result, string(runes[:i]))
	}
	return result
}
