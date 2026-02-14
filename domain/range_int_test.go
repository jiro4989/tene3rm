package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRangeInt(t *testing.T) {
	tests := []struct {
		name          string
		value, mn, mx int
		want          RangeInt
		wantErr       bool
	}{
		{
			name:  "正常系: 値が範囲内",
			value: 5,
			mn:    1,
			mx:    10,
			want: RangeInt{
				value:    5,
				minValue: 1,
				maxValue: 10,
			},
			wantErr: false,
		},
		{
			name:    "異常系: min, max が異常",
			value:   5,
			mn:      10,
			mx:      1,
			want:    RangeInt{},
			wantErr: true,
		},
		{
			name:    "異常系: 値が範囲外はエラー",
			value:   99,
			mn:      1,
			mx:      10,
			want:    RangeInt{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := NewRangeInt(tt.value, tt.mn, tt.mx)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.NoError(err)
			a.Equal(tt.want, got)
		})
	}
}

func TestRangeInt_SafePlus(t *testing.T) {
	tests := []struct {
		name string
		r    RangeInt
		n    int
		want RangeInt
	}{
		{
			name: "正常系: 値が範囲内なら加算できる",
			r: RangeInt{
				value:    5,
				minValue: 1,
				maxValue: 10,
			},
			n: 3,
			want: RangeInt{
				value:    8,
				minValue: 1,
				maxValue: 10,
			},
		},
		{
			name: "正常系: 値が範囲を越えようとしても上限以内に収まる",
			r: RangeInt{
				value:    5,
				minValue: 1,
				maxValue: 10,
			},
			n: 20,
			want: RangeInt{
				value:    10,
				minValue: 1,
				maxValue: 10,
			},
		},
		{
			name: "正常系: 値が範囲を越えようとしても下限以内に収まる",
			r: RangeInt{
				value:    5,
				minValue: 1,
				maxValue: 10,
			},
			n: -20,
			want: RangeInt{
				value:    1,
				minValue: 1,
				maxValue: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.r.SafePlus(tt.n)
			a.Equal(tt.want, got)
		})
	}
}

func TestRangeInt_Value(t *testing.T) {
	r := RangeInt{
		value:    4,
		minValue: 1,
		maxValue: 10,
	}
	assert.Equal(t, 4, r.Value())
}
