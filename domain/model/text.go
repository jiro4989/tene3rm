package model

import "strings"

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
func (t UserInputText) Prefixes() UserInputTexts {
	// マルチバイト文字を考慮するため rune に変換
	runes := []rune(t.value)
	result := make([]string, 0, len(runes))
	for i := 1; i <= len(runes); i++ {
		result = append(result, string(runes[:i]))
	}
	return NewUserInputTexts(result)
}

type UserInputTexts struct {
	value []UserInputText
}

func NewUserInputTexts(s []string) UserInputTexts {
	t := UserInputTexts{}
	for _, v := range s {
		t.value = append(t.value, NewUserInputText(v))
	}
	return t
}

func (t UserInputTexts) Contains(t2 UserInputText) bool {
	for _, v := range t.value {
		if v.Equal(t2) {
			return true
		}
	}
	return false
}
