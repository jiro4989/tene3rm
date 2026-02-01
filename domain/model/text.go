package model

import "strings"

type Text struct {
	value string
}

func NewText(s string) Text {
	s = strings.TrimSpace(s)
	return Text{
		value: s,
	}
}

func (t Text) Equal(t2 Text) bool {
	return t.value == t2.value
}

// prefixes は yes を y, ye, yes といった感じの配列にして返す。
func (t Text) Prefixes() Texts {
	// マルチバイト文字を考慮するため rune に変換
	runes := []rune(t.value)
	result := make([]string, 0, len(runes))
	for i := 1; i <= len(runes); i++ {
		result = append(result, string(runes[:i]))
	}
	return NewTexts(result)
}

type Texts struct {
	value []Text
}

func NewTexts(s []string) Texts {
	t := Texts{}
	for _, v := range s {
		t.value = append(t.value, NewText(v))
	}
	return t
}

func (t Texts) Contains(t2 Text) bool {
	for _, v := range t.value {
		if v.Equal(t2) {
			return true
		}
	}
	return false
}
