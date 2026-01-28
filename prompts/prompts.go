package prompts

import (
	"math/rand"
)

type PromptFunc func(string) (bool, error)

type Prompts map[string]PromptFunc

const appname = "tene3rm"

// Prompt はランダムにプロンプトを表示し、ユーザの入力に応じてファイルを削除するかどうかを返す。
func Prompt(path string, seed int64) (bool, error) {
	prompts := Prompts{
		"yes_no":        promptWithYesNo,
		"yes_no_jp":     promptWithYesNoInJapanese,
		"yes_no_denial": promptWithYesNoDenial,
		"math":          promptWithMath,
	}

	f := selectFunc(prompts, seed)
	return f(path)
}

// selectFunc はpromptsからランダムに1つPromptFuncを返す。
func selectFunc(prompts Prompts, seed int64) PromptFunc {
	keys := make([]string, 0)
	for k := range prompts {
		keys = append(keys, k)
	}

	r := rand.New(rand.NewSource(seed))
	i := r.Intn(len(keys))
	k := keys[i]
	f := prompts[k]
	return f
}
