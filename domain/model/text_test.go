package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextEqual(t *testing.T) {
	tests := []struct {
		name string
		a    Text
		b    Text
		want bool
	}{
		{
			name: "正常系: 等しい場合は true",
			a:    NewText("yes"),
			b:    NewText("yes"),
			want: true,
		},
		{
			name: "正常系: 等しくない場合は false",
			a:    NewText("yes"),
			b:    NewText("ye"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.a.Equal(tt.b)
			a.Equal(tt.want, got)
		})
	}
}

func TestTextPrefixes(t *testing.T) {
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

func TestTextContains(t *testing.T) {
	tests := []struct {
		name string
		a    Texts
		b    Text
		want bool
	}{
		{
			name: "正常系: 含まれるなら true",
			a:    NewTexts([]string{"y", "ye", "yes"}),
			b:    NewText("yes"),
			want: true,
		},
		{
			name: "正常系: 含まれるなら true",
			a:    NewTexts([]string{"y", "ye", "yes"}),
			b:    NewText("ye"),
			want: true,
		},
		{
			name: "正常系: 含まれるなら true",
			a:    NewTexts([]string{"は", "はい"}),
			b:    NewText("はい"),
			want: true,
		},
		{
			name: "正常系: 含まれないなら false",
			a:    NewTexts([]string{"y", "ye", "yes"}),
			b:    NewText("no"),
			want: false,
		},
		{
			name: "正常系: 空配列でもエラーにならない",
			a:    NewTexts([]string{}),
			b:    NewText("no"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.a.Contains(tt.b)
			a.Equal(tt.want, got)
		})
	}
}
