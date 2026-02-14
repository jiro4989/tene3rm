package column

import (
	"testing"

	"github.com/jiro4989/tene3rm/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPosition_MoveLeft(t *testing.T) {
	x, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p := NewPosition(x, y)

	x2, err := domain.NewRangeInt(2, 1, 5)
	require.NoError(t, err)
	y2, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p2 := NewPosition(x2, y2)

	x3, err := domain.NewRangeInt(1, 1, 5)
	require.NoError(t, err)
	y3, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p3 := NewPosition(x3, y3)

	tests := []struct {
		name string
		p    Position
		want Position
	}{
		{
			name: "正常系: 値が範囲内",
			p:    p,
			want: p2,
		},
		{
			name: "正常系: 値が範囲を超えることはない",
			p:    p3,
			want: p3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.p.MoveLeft()
			a.Equal(tt.want, got)
		})
	}
}

func TestPosition_MoveRight(t *testing.T) {
	x, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p := NewPosition(x, y)

	x2, err := domain.NewRangeInt(4, 1, 5)
	require.NoError(t, err)
	y2, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p2 := NewPosition(x2, y2)

	x3, err := domain.NewRangeInt(5, 1, 5)
	require.NoError(t, err)
	y3, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p3 := NewPosition(x3, y3)

	tests := []struct {
		name string
		p    Position
		want Position
	}{
		{
			name: "正常系: 値が範囲内",
			p:    p,
			want: p2,
		},
		{
			name: "正常系: 値が範囲を超えることはない",
			p:    p3,
			want: p3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.p.MoveRight()
			a.Equal(tt.want, got)
		})
	}
}

func TestPosition_MoveDown(t *testing.T) {
	x, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p := NewPosition(x, y)

	x2, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y2, err := domain.NewRangeInt(4, 1, 5)
	require.NoError(t, err)
	p2 := NewPosition(x2, y2)

	x3, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y3, err := domain.NewRangeInt(5, 1, 5)
	require.NoError(t, err)
	p3 := NewPosition(x3, y3)

	tests := []struct {
		name string
		p    Position
		want Position
	}{
		{
			name: "正常系: 値が範囲内",
			p:    p,
			want: p2,
		},
		{
			name: "正常系: 値が範囲を超えることはない",
			p:    p3,
			want: p3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.p.MoveDown()
			a.Equal(tt.want, got)
		})
	}
}

func TestPosition_MoveUp(t *testing.T) {
	x, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	p := NewPosition(x, y)

	x2, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y2, err := domain.NewRangeInt(2, 1, 5)
	require.NoError(t, err)
	p2 := NewPosition(x2, y2)

	x3, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y3, err := domain.NewRangeInt(1, 1, 5)
	require.NoError(t, err)
	p3 := NewPosition(x3, y3)

	tests := []struct {
		name string
		p    Position
		want Position
	}{
		{
			name: "正常系: 値が範囲内",
			p:    p,
			want: p2,
		},
		{
			name: "正常系: 値が範囲を超えることはない",
			p:    p3,
			want: p3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got := tt.p.MoveUp()
			a.Equal(tt.want, got)
		})
	}
}

func TestPosition_XY(t *testing.T) {
	x, err := domain.NewRangeInt(3, 1, 5)
	require.NoError(t, err)
	y, err := domain.NewRangeInt(4, 1, 5)
	require.NoError(t, err)
	p := NewPosition(x, y)

	assert.Equal(t, 3, p.X())
	assert.Equal(t, 4, p.Y())
}
