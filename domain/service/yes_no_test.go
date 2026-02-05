package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextServiceJudgeYesNo(t *testing.T) {
	tests := []struct {
		name  string
		svc   TextService
		input string
		want  bool
	}{
		{
			name:  "正常系: yes のときは true",
			svc:   NewTextService(),
			input: "yes",
			want:  true,
		},
		{
			name:  "正常系: ye のときは true",
			svc:   NewTextService(),
			input: "ye",
			want:  true,
		},
		{
			name:  "正常系: y のときは true",
			svc:   NewTextService(),
			input: "y",
			want:  true,
		},
		{
			name:  "正常系: yes でないときは false",
			svc:   NewTextService(),
			input: "no",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewTextService(),
			input: "",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.svc.JudgeYesNo(tt.input)
			a.Equal(tt.want, got)
		})
	}
}

func TestTextServiceJudgeYesNoDenial(t *testing.T) {
	tests := []struct {
		name  string
		svc   TextService
		input string
		want  bool
	}{
		{
			name:  "正常系: no のときは true",
			svc:   NewTextService(),
			input: "no",
			want:  true,
		},
		{
			name:  "正常系: n のときは true",
			svc:   NewTextService(),
			input: "n",
			want:  true,
		},
		{
			name:  "正常系: no でないときは false",
			svc:   NewTextService(),
			input: "yes",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewTextService(),
			input: "",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.svc.JudgeYesNoDenial(tt.input)
			a.Equal(tt.want, got)
		})
	}
}

func TestTextServiceJudgeYesNoJapanese(t *testing.T) {
	tests := []struct {
		name  string
		svc   TextService
		input string
		want  bool
	}{
		{
			name:  "正常系: はいのときは true",
			svc:   NewTextService(),
			input: "はい",
			want:  true,
		},
		{
			name:  "正常系: はのときは true",
			svc:   NewTextService(),
			input: "は",
			want:  true,
		},
		{
			name:  "正常系: はいでないときは false",
			svc:   NewTextService(),
			input: "いいえ",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewTextService(),
			input: "",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.svc.JudgeYesNoJapanese(tt.input)
			a.Equal(tt.want, got)
		})
	}
}
