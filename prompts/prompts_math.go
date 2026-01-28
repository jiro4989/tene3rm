package prompts

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

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
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", appname, path, a, op, b),
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
