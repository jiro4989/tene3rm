package terminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "正常系: yes から y, ye, yes が得られる",
			s:    "yes",
			want: []string{"y", "ye", "yes"},
		},
		{
			name: "正常系: いいえ から い, いい, いいえ が得られる",
			s:    "いいえ",
			want: []string{"い", "いい", "いいえ"},
		},
		{
			name: "正常系: 一文字の場合でも問題ない",
			s:    "い",
			want: []string{"い"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := prefixes(tt.s)
			a.Equal(tt.want, got)
		})
	}
}
