package prompts

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/jiro4989/tene3rm/domain"
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

	ops := []domain.Operator{
		&domain.PlusOperator{},
		&domain.MinusOperator{},
		&domain.MultiOperator{},
	}

	a := rand.Intn(9) + 1
	b := rand.Intn(9) + 1

	i := rand.Intn(len(ops))
	op := ops[i]
	c := op.Do(a, b)

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", appname, path, a, op.Op(), b),
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
