package prompts

import (
	"math/rand"
)

type promptFunc func(string) (bool, error)

type promptFuncs map[string]promptFunc

const appname = "tene3rm"

// Prompt はランダムにプロンプトを表示し、ユーザの入力に応じてファイルを削除するかどうかを返す。
func Prompt(path string, seed int64) (bool, error) {
	pfs := promptFuncs{
		"yes_no":        promptWithYesNo,
		"yes_no_jp":     promptWithYesNoInJapanese,
		"yes_no_denial": promptWithYesNoDenial,
		"math":          promptWithMath,
	}

	f := selectFunc(pfs, seed)
	return f(path)
}

// selectFunc はpfsからランダムに1つpromptFuncを返す。
func selectFunc(pfs promptFuncs, seed int64) promptFunc {
	keys := make([]string, 0)
	for k := range pfs {
		keys = append(keys, k)
	}

	r := rand.New(rand.NewSource(seed))
	i := r.Intn(len(keys))
	k := keys[i]
	f := pfs[k]
	return f
}
