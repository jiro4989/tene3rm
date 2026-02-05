package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNumber(t *testing.T) {
	a := assert.New(t)

	a.True(NewNumber(5).Equal(NewNumber(5)))
	a.False(NewNumber(6).Equal(NewNumber(5)))
	a.Equal(5, NewNumber(5).Value())
}

func TestNewNumberWithUserInputText(t *testing.T) {
	tests := []struct {
		name    string
		input   UserInputText
		want    Number
		wantErr bool
	}{
		{
			name:    "正常系: 5はNumberに変換できる",
			input:   NewUserInputText("5"),
			want:    NewNumber(5),
			wantErr: false,
		},
		{
			name:    "正常系: 空白は無視される",
			input:   NewUserInputText(" 5 "),
			want:    NewNumber(5),
			wantErr: false,
		},
		{
			name:    "異常系: 数値変換出来ない場合はエラー",
			input:   NewUserInputText("hoge"),
			want:    NewNumber(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := NewNumberWithUserInputText(tt.input)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, got)
		})
	}
}
