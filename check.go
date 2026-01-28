package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/manifoldco/promptui"
)

type CheckFunc func(string) (bool, error)

type Checks map[string]CheckFunc

// Check はランダムにプロンプトを表示し、ユーザの入力に応じてファイルを削除するかどうかを返す。
func Check(path string, seed int64) (bool, error) {
	checks := Checks{
		"text": checkWithText,
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
		Label:    fmt.Sprintf("%s: remove file '%s'?", Appname, path),
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
