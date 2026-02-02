package terminal

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jiro4989/tene3rm/domain/model"
	"github.com/jiro4989/tene3rm/domain/service"
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

	svc := service.NewMathService()
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	want, a, b, op := svc.SimpleOperations(r, r, r)

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", appname, path, a.Value(), op, b.Value()),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	resultNum, err := model.NewNumberWithText(model.NewText(result))
	if err != nil {
		return false, err
	}

	return want.Equal(resultNum), nil
}

// promptWithMath2 は筆算での計算結果を求めるプロンプトを表示する。
func promptWithMath2(path string) (bool, error) {
	return false, nil
}
