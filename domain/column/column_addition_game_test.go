package column

import (
	"testing"

	"github.com/jiro4989/tene3rm/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColumnAdditionGame(t *testing.T) {
	g, err := NewColumnAdditionGame(12, 34)
	require.NoError(t, err)

	x1, err := domain.NewRangeInt(2, 0, 3)
	require.NoError(t, err)
	y1, err := domain.NewRangeInt(0, 0, 2)
	require.NoError(t, err)
	p1 := NewPosition(x1, y1)

	g2 := g
	g2.pos = p1

	got := g.MoveLeft()
	assert.Equal(t, g2, got)

	got = g.MoveRight()
	assert.Equal(t, g, got)

	got = g.MoveUp()
	assert.Equal(t, g, got)

	x2, err := domain.NewRangeInt(3, 0, 3)
	require.NoError(t, err)
	y2, err := domain.NewRangeInt(1, 0, 2)
	require.NoError(t, err)
	p2 := NewPosition(x2, y2)

	g3 := g
	g3.pos = p2

	got = g.MoveDown()
	assert.Equal(t, g3, got)

	x, y := g.PositionXY()
	assert.Equal(t, 3, x)
	assert.Equal(t, 0, y)

	g4 := g.SetString("9")
	assert.Equal(t, "9", g4.CurrentPositionCellValue())
	assert.Equal(t, []rune("9")[0], g4.CurrentPositionCellValueRune())
}

func TestColumnAdditionGame_ResultStringLines(t *testing.T) {
	g, err := NewColumnAdditionGame(12, 34)
	require.NoError(t, err)

	g = g.SetString("8")
	g = g.MoveLeft()
	g = g.SetString("4")

	g = g.MoveDown()

	g = g.SetString("6")
	g = g.MoveLeft()
	g = g.SetString("3")

	g = g.MoveDown()
	g = g.MoveRight()
	g = g.MoveRight()

	g = g.SetString("8")
	g = g.MoveLeft()
	g = g.SetString("0")
	g = g.MoveLeft()
	g = g.SetString("4")

	want := []string{
		"    1 2",
		"  x 3 4",
		"-------",
		"    4 8",
		"  3 6  ",
		"-------",
		"  4 0 8",
	}

	got := g.ResultStringLines()
	assert.Equal(t, want, got)

	b, err := g.Evaluate()
	assert.NoError(t, err)
	assert.Equal(t, true, b)

	// 間違った値に変更
	g = g.SetString("5")
	b, err = g.Evaluate()
	assert.NoError(t, err)
	assert.Equal(t, false, b)
}
