package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTwoDigitMultiplyColumnAddition_MultiplyOnesPlace(t *testing.T) {
	c, err := NewTwoDigitMultiplyColumnAddition(12, 34)
	require.NoError(t, err)

	got := c.MultiplyOnesPlace()
	assert.Equal(t, 48, got)
}

func TestTwoDigitMultiplyColumnAddition_MultiplyTensPlace(t *testing.T) {
	c, err := NewTwoDigitMultiplyColumnAddition(12, 34)
	require.NoError(t, err)

	got := c.MultiplyTensPlace()
	assert.Equal(t, 36, got)
}

func TestTwoDigitMultiplyColumnAddition_Multiply(t *testing.T) {
	c, err := NewTwoDigitMultiplyColumnAddition(12, 34)
	require.NoError(t, err)

	got := c.Multiply()
	assert.Equal(t, 408, got)
}

func TestTwoDigitMultiplyColumnAddition_Equal(t *testing.T) {
	c, err := NewTwoDigitMultiplyColumnAddition(12, 34)
	require.NoError(t, err)

	tests := []struct {
		name       string
		c          TwoDigitMultiplyColumnAddition
		n1, n2, n3 int
		want       bool
	}{
		{
			name: "正常系: 計算結果が一致する",
			c:    c,
			n1:   48,
			n2:   36,
			n3:   408,
			want: true,
		},
		{
			name: "正常系: 計算結果が一致しない",
			c:    c,
			n1:   48,
			n2:   36,
			n3:   409,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.c.Equal(tt.n1, tt.n2, tt.n3)
			a.Equal(tt.want, got)
		})
	}
}
