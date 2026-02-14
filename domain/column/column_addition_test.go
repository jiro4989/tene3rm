package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTwoDigitMultiplyColumnAddition(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    TwoDigitMultiplyColumnAddition
		wantErr bool
	}{
		{
			name: "正常系: 値が範囲内",
			a:    20,
			b:    30,
			want: TwoDigitMultiplyColumnAddition{
				a:  20,
				b:  30,
				op: "x",
			},
			wantErr: false,
		},
		{
			name:    "異常系: 値は 2 桁の範囲でなければならない",
			a:       9,
			b:       30,
			want:    TwoDigitMultiplyColumnAddition{},
			wantErr: true,
		},
		{
			name:    "異常系: 値は 2 桁の範囲でなければならない",
			a:       100,
			b:       30,
			want:    TwoDigitMultiplyColumnAddition{},
			wantErr: true,
		},
		{
			name:    "異常系: 値は 2 桁の範囲でなければならない",
			a:       20,
			b:       9,
			want:    TwoDigitMultiplyColumnAddition{},
			wantErr: true,
		},
		{
			name:    "異常系: 値は 2 桁の範囲でなければならない",
			a:       20,
			b:       100,
			want:    TwoDigitMultiplyColumnAddition{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := NewTwoDigitMultiplyColumnAddition(tt.a, tt.b)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.Equal(tt.want, got)
		})
	}
}
