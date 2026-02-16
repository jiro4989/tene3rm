package terminal

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/jiro4989/tene3rm/usecase"
	"github.com/manifoldco/promptui"
)

// promptWithRandomCharacterFile はランダムな文字列が書かれたファイルを出力し、
// そのファイルの内容の入力を求めるプロンプトを表示する。
func promptWithRandomStringFile(path string) (bool, error) {
	validate := func(input string) error {
		return nil
	}

	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	fr := infra.NewFileRepo(os.TempDir())
	uc := usecase.NewGenerateStringUsecase(r, fr)
	const filename = "tmp.txt"
	want, err := uc.Execute(64, filename)
	if err != nil {
		return false, err
	}

	fmt.Println(fmt.Sprintf("%s: remove file '%s'.", appname, path))

	tmpFile := fr.FullPath(filename)
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
