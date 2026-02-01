package terminal

import (
	"math/rand"
)

type promptFunc func(string) (bool, error)

type promptFuncs map[string]promptFunc

const appname = "tene3rm"

// Prompt はランダムにプロンプトを表示し、ユーザの入力に応じてファイルを削除するかどうかを返す。
func Prompt(path string, seed int64) (bool, error) {
	pfs := []promptFunc{
		promptWithYesNo,
		promptWithYesNoInJapanese,
		promptWithYesNoInJapanese3,
		promptWithYesNoDenial,
		promptWithMath,
		promptWithTimer,
	}

	r := rand.New(rand.NewSource(seed))
	i := r.Intn(len(pfs))
	pf := pfs[i]
	return pf(path)
}
