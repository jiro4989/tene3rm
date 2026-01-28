package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

type CheckFunc func(string) (bool, error)

type Checks map[string]CheckFunc

// Check はランダムにプロンプトを表示し、ユーザの入力に応じてファイルを削除するかどうかを返す。
func Check(path string, seed int64) (bool, error) {
	checks := Checks{
		"text":    checkWithText,
		"text_jp": checkWithTextInJapanese,
		"math":    checkWithMath,
	}

	f := selectFunc(checks, seed)
	return f(path)
}

// selectFunc はchecksからランダムに1つCheckFuncを返す。
func selectFunc(checks Checks, seed int64) CheckFunc {
	keys := make([]string, 0)
	for k := range checks {
		keys = append(keys, k)
	}

	r := rand.New(rand.NewSource(seed))
	i := r.Intn(len(keys))
	k := keys[i]
	f := checks[k]
	return f
}

// checkWithText はシンプルなYes/Noプロンプトを表示する。
func checkWithText(path string) (bool, error) {
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

// checkWithTextInJapanese はシンプルなはい/いいえプロンプトを表示する。
func checkWithTextInJapanese(path string) (bool, error) {
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

// checkWithTextInJapanese はシンプルなはい/いいえプロンプトを表示する。
func checkWithMath(path string) (bool, error) {
	validate := func(input string) error {
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
