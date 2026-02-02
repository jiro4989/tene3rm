package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNumberWithText(t *testing.T) {
	tests := []struct {
		name    string
		input   Text
		want    Number
		wantErr bool
	}{
		{
			name:    "正常系: 5はNumberに変換できる",
			input:   NewText("5"),
			want:    NewNumber(5),
			wantErr: false,
		},
		{
			name:    "正常系: 空白は無視される",
			input:   NewText(" 5 "),
			want:    NewNumber(5),
			wantErr: false,
		},
		{
			name:    "異常系: 数値変換出来ない場合はエラー",
			input:   NewText("hoge"),
			want:    NewNumber(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := NewNumberWithText(tt.input)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, got)
		})
	}
}
