package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInputTextEqual(t *testing.T) {
	tests := []struct {
		name string
		a    UserInputText
		b    UserInputText
		want bool
	}{
		{
			name: "正常系: 等しい場合は true",
			a:    NewUserInputText("yes"),
			b:    NewUserInputText("yes"),
			want: true,
		},
		{
			name: "正常系: 等しくない場合は false",
			a:    NewUserInputText("yes"),
			b:    NewUserInputText("ye"),
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

func TestUserInputTextPrefixes(t *testing.T) {
	tests := []struct {
		name string
		s    UserInputText
		want []string
	}{
		{
			name: "正常系: yes から y, ye, yes が得られる",
			s:    NewUserInputText("yes"),
			want: []string{"y", "ye", "yes"},
		},
		{
			name: "正常系: いいえ から い, いい, いいえ が得られる",
			s:    NewUserInputText("いいえ"),
			want: []string{"い", "いい", "いいえ"},
		},
		{
			name: "正常系: 一文字の場合でも問題ない",
			s:    NewUserInputText("い"),
			want: []string{"い"},
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

func TestUserInputTextIn(t *testing.T) {
	tests := []struct {
		name string
		a    UserInputText
		b    []string
		want bool
	}{
		{
			name: "正常系: 含まれるなら true",
			a:    NewUserInputText("yes"),
			b:    []string{"y", "ye", "yes"},
			want: true,
		},
		{
			name: "正常系: 含まれるなら true",
			a:    NewUserInputText("ye"),
			b:    []string{"y", "ye", "yes"},
			want: true,
		},
		{
			name: "正常系: 含まれるなら true",
			a:    NewUserInputText("はい"),
			b:    []string{"は", "はい"},
			want: true,
		},
		{
			name: "正常系: 含まれないなら false",
			a:    NewUserInputText("no"),
			b:    []string{"y", "ye", "yes"},
			want: false,
		},
		{
			name: "正常系: 空配列でもエラーにならない",
			a:    NewUserInputText("no"),
			b:    []string{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.a.In(tt.b)
			a.Equal(tt.want, got)
		})
	}
}
