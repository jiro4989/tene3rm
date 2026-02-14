package domain

import (
	"strconv"
	"strings"
)

type UserInputText struct {
	value string
}

func NewUserInputText(s string) UserInputText {
	s = strings.TrimSpace(s)
	return UserInputText{
		value: s,
	}
}

func (t UserInputText) Equal(t2 UserInputText) bool {
	return t.value == t2.value
}

// prefixes は yes を y, ye, yes といった感じの配列にして返す。
func (t UserInputText) Prefixes() []string {
	// マルチバイト文字を考慮するため rune に変換
	runes := []rune(t.value)
	result := make([]string, 0, len(runes))
	for i := 1; i <= len(runes); i++ {
		result = append(result, string(runes[:i]))
	}
	return result
}

func (t UserInputText) ToInt() (int, error) {
	return strconv.Atoi(t.value)
}

func (t UserInputText) In(vals []string) bool {
	for _, v := range vals {
		if t.value == v {
			return true
		}
	}
	return false
}
