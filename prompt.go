package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type PromptFunc func(string) (bool, error)

type Prompts map[string]PromptFunc

// Prompt はランダムにプロンプトを表示し、ユーザの入力に応じてファイルを削除するかどうかを返す。
func Prompt(path string, seed int64) (bool, error) {
	prompts := Prompts{
		"text":        promptWithText,
		"text_jp":     promptWithTextInJapanese,
		"text_denial": promptWithTextDenial,
		"math":        promptWithMath,
	}

	f := selectFunc(prompts, seed)
	return f(path)
}

// selectFunc はpromptsからランダムに1つPromptFuncを返す。
func selectFunc(prompts Prompts, seed int64) PromptFunc {
	keys := make([]string, 0)
	for k := range prompts {
		keys = append(keys, k)
	}

	r := rand.New(rand.NewSource(seed))
	i := r.Intn(len(keys))
	k := keys[i]
	f := prompts[k]
	return f
}

// promptWithText はシンプルなYes/Noプロンプトを表示する。
func promptWithText(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? [y/n]", Appname, path),
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

// promptWithTextDenial はNoのときだけファイルを削除するプロンプトを表示する。
func promptWithTextDenial(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: DON'T remove file '%s'? [y/n]", Appname, path),
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

// promptWithTextInJapanese はシンプルなはい/いいえプロンプトを表示する。
func promptWithTextInJapanese(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: '%s' ファイルを削除しますか? [はい/いいえ]", Appname, path),
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

// promptWithMath は単純な算数入力を求めるプロンプトを表示する。
func promptWithMath(path string) (bool, error) {
	validate := func(input string) error {
		input = strings.TrimSpace(input)
		_, err := strconv.Atoi(input)
		if err != nil {
			return err
		}
		return nil
	}

	calcs := map[string]func(int, int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"*": func(a, b int) int {
			return a * b
		},
	}

	a := rand.Intn(9) + 1
	b := rand.Intn(9) + 1

	keys := make([]string, 0)
	for k := range calcs {
		keys = append(keys, k)
	}
	i := rand.Intn(len(keys))
	op := keys[i]
	f := calcs[op]
	c := f(a, b)

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", Appname, path, a, op, b),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	result = strings.TrimSpace(result)
	resultNum, err := strconv.Atoi(result)
	if err != nil {
		return false, err
	}

	return resultNum == c, nil
}
