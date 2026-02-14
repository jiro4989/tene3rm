package terminal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jiro4989/tene3rm/domain"
	"github.com/jiro4989/tene3rm/usecase"
	"github.com/manifoldco/promptui"
)

// promptWithSimpleOperation は単純な算数入力を求めるプロンプトを表示する。
func promptWithSimpleOperation(path string) (bool, error) {
	validate := func(input string) error {
		_, err := domain.NewUserInputText(input).ToInt()
		if err != nil {
			return err
		}
		return nil
	}

	svc := usecase.NewSimpleOperationUsecase()
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	want, a, b, op := svc.Execute(r, r, r)

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", appname, path, a, op, b),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	resultNum, err := domain.NewUserInputText(result).ToInt()
	if err != nil {
		return false, err
	}

	return want == resultNum, nil
}
