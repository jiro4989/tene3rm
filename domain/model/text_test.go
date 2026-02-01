package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixes(t *testing.T) {
	tests := []struct {
		name string
		s    Text
		want Texts
	}{
		{
			name: "正常系: yes から y, ye, yes が得られる",
			s:    NewText("yes"),
			want: NewTexts([]string{"y", "ye", "yes"}),
		},
		{
			name: "正常系: いいえ から い, いい, いいえ が得られる",
			s:    NewText("いいえ"),
			want: NewTexts([]string{"い", "いい", "いいえ"}),
		},
		{
			name: "正常系: 一文字の場合でも問題ない",
			s:    NewText("い"),
			want: NewTexts([]string{"い"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.s.Prefixes()
			a.Equal(tt.want, got)
		})
	}
}
