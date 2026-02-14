package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYesNoUsecaseJudgeYesNo(t *testing.T) {
	tests := []struct {
		name  string
		svc   YesNoUsecase
		input string
		want  bool
	}{
		{
			name:  "正常系: yes のときは true",
			svc:   NewYesNoUsecase(),
			input: "yes",
			want:  true,
		},
		{
			name:  "正常系: ye のときは true",
			svc:   NewYesNoUsecase(),
			input: "ye",
			want:  true,
		},
		{
			name:  "正常系: y のときは true",
			svc:   NewYesNoUsecase(),
			input: "y",
			want:  true,
		},
		{
			name:  "正常系: yes でないときは false",
			svc:   NewYesNoUsecase(),
			input: "no",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewYesNoUsecase(),
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

func TestYesNoUsecaseJudgeYesNoDenial(t *testing.T) {
	tests := []struct {
		name  string
		svc   YesNoUsecase
		input string
		want  bool
	}{
		{
			name:  "正常系: no のときは true",
			svc:   NewYesNoUsecase(),
			input: "no",
			want:  true,
		},
		{
			name:  "正常系: n のときは true",
			svc:   NewYesNoUsecase(),
			input: "n",
			want:  true,
		},
		{
			name:  "正常系: no でないときは false",
			svc:   NewYesNoUsecase(),
			input: "yes",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewYesNoUsecase(),
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

func TestYesNoUsecaseJudgeYesNoJapanese(t *testing.T) {
	tests := []struct {
		name  string
		svc   YesNoUsecase
		input string
		want  bool
	}{
		{
			name:  "正常系: はいのときは true",
			svc:   NewYesNoUsecase(),
			input: "はい",
			want:  true,
		},
		{
			name:  "正常系: はのときは true",
			svc:   NewYesNoUsecase(),
			input: "は",
			want:  true,
		},
		{
			name:  "正常系: はいでないときは false",
			svc:   NewYesNoUsecase(),
			input: "いいえ",
			want:  false,
		},
		{
			name:  "正常系: 空文字列は false",
			svc:   NewYesNoUsecase(),
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
