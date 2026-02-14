package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYesNoServiceJudgeYesNo(t *testing.T) {
	tests := []struct {
		name  string
		svc   YesNoService
		input string
		want  bool
	}{
		{
			name:  "正常系: yes のときは true",
			svc:   NewYesNoService(),
			input: "yes",
			want:  true,
		},
		{
			name:  "正常系: ye のときは true",
			svc:   NewYesNoService(),
			input: "ye",
			want:  true,
		},
		{
			name:  "正常系: y のときは true",
			svc:   NewYesNoService(),
			input: "y",
			want:  true,
		},
		{
			name:  "正常系: yes でないときは false",
			svc:   NewYesNoService(),
			input: "no",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewYesNoService(),
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

func TestYesNoServiceJudgeYesNoDenial(t *testing.T) {
	tests := []struct {
		name  string
		svc   YesNoService
		input string
		want  bool
	}{
		{
			name:  "正常系: no のときは true",
			svc:   NewYesNoService(),
			input: "no",
			want:  true,
		},
		{
			name:  "正常系: n のときは true",
			svc:   NewYesNoService(),
			input: "n",
			want:  true,
		},
		{
			name:  "正常系: no でないときは false",
			svc:   NewYesNoService(),
			input: "yes",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewYesNoService(),
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

func TestYesNoServiceJudgeYesNoJapanese(t *testing.T) {
	tests := []struct {
		name  string
		svc   YesNoService
		input string
		want  bool
	}{
		{
			name:  "正常系: はいのときは true",
			svc:   NewYesNoService(),
			input: "はい",
			want:  true,
		},
		{
			name:  "正常系: はのときは true",
			svc:   NewYesNoService(),
			input: "は",
			want:  true,
		},
		{
			name:  "正常系: はいでないときは false",
			svc:   NewYesNoService(),
			input: "いいえ",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewYesNoService(),
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
