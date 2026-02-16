package terminal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jiro4989/tene3rm/usecase"
	"github.com/manifoldco/promptui"
)

// promptWithCaptcha はキャプチャ画像をファイルに出力し、
// そのファイルの内容の入力を求めるプロンプトを表示する。
func promptWithCaptcha(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	tmpFile := filepath.Join(os.TempDir(), "tmp.png")
	fp, err := os.Create(tmpFile)
	if err != nil {
		return false, err
	}
	defer fp.Close()

	uc := usecase.NewCaptchaUsecase()
	want, err := uc.Execute(fp)
	if err != nil {
		return false, err
	}

	fmt.Println(fmt.Sprintf("%s: remove file '%s'.", appname, path))
	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: what was written in the '%s' file?", appname, tmpFile),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	return want == result, nil
}
