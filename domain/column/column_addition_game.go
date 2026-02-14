package column

import (
	"fmt"

	"github.com/jiro4989/tene3rm/domain"
)

type ColumnAdditionGame struct {
	ca    TwoDigitMultiplyColumnAddition
	pos   Position
	cells Cells
}

func NewColumnAdditionGame(a, b int) (ColumnAdditionGame, error) {
	ca, err := NewTwoDigitMultiplyColumnAddition(a, b)
	if err != nil {
		return ColumnAdditionGame{}, err
	}

	x, err := domain.NewRangeInt(3, 0, 3)
	if err != nil {
		return ColumnAdditionGame{}, err
	}

	y, err := domain.NewRangeInt(0, 0, 2)
	if err != nil {
		return ColumnAdditionGame{}, err
	}

	return ColumnAdditionGame{
		ca:    ca,
		pos:   NewPosition(x, y),
		cells: NewCells(),
	}, nil
}

func (c ColumnAdditionGame) MoveLeft() ColumnAdditionGame {
	c.pos = c.pos.MoveLeft()
	return c
}

func (c ColumnAdditionGame) MoveRight() ColumnAdditionGame {
	c.pos = c.pos.MoveRight()
	return c
}

func (c ColumnAdditionGame) MoveUp() ColumnAdditionGame {
	c.pos = c.pos.MoveUp()
	return c
}

func (c ColumnAdditionGame) MoveDown() ColumnAdditionGame {
	c.pos = c.pos.MoveDown()
	return c
}

func (c ColumnAdditionGame) SetString(s string) ColumnAdditionGame {
	cell := NewCell(s)
	c.cells.value[c.pos.Y()].value[c.pos.X()] = cell
	return c
}

func (c ColumnAdditionGame) PositionXY() (int, int) {
	return c.pos.X(), c.pos.Y()
}

func (c ColumnAdditionGame) CurrentPositionCellValue() string {
	return c.cells.value[c.pos.Y()].value[c.pos.X()].value
}

func (c ColumnAdditionGame) CurrentPositionCellValueRune() rune {
	return []rune(c.CurrentPositionCellValue())[0]
}

func (c ColumnAdditionGame) ResultStringLines() []string {
	result := make([]string, 0)
	a10, a1 := divMod(c.ca.a)
	b10, b1 := divMod(c.ca.b)
	const horizontalLine = "-------"
	result = append(result, fmt.Sprintf("    %d %d", a10, a1))
	result = append(result, fmt.Sprintf("  x %d %d", b10, b1))
	result = append(result, horizontalLine)
	result = append(result, c.cells.value[0].ResultStringLine())
	result = append(result, c.cells.value[1].ResultStringLine())
	result = append(result, horizontalLine)
	result = append(result, c.cells.value[2].ResultStringLine())
	return result
}

func divMod(n int) (int, int) {
	return n / 10, n % 10
}

func (c ColumnAdditionGame) Evaluate() (bool, error) {
	n1, err := c.cells.value[0].ToInt()
	if err != nil {
		return false, err
	}
	n2, err := c.cells.value[1].ToInt()
	if err != nil {
		return false, err
	}
	n3, err := c.cells.value[2].ToInt()
	if err != nil {
		return false, err
	}
	return c.ca.Equal(n1, n2, n3), nil
}
