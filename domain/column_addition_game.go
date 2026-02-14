package domain

import (
	"fmt"
	"strconv"
	"strings"
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
	return ColumnAdditionGame{
		ca:    ca,
		pos:   NewPosition(3, 0, 0, 0, 3, 2),
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
	c.cells.value[c.pos.y].value[c.pos.x] = cell
	return c
}

func (c ColumnAdditionGame) PositionXY() (int, int) {
	return c.pos.x, c.pos.y
}

func (c ColumnAdditionGame) CurrentPositionCellValue() string {
	return c.cells.value[c.pos.y].value[c.pos.x].value
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

type Cell struct {
	value string
}

func NewCell(s string) Cell {
	return Cell{value: s}
}

func NewEmptyCell() Cell {
	return NewCell(" ")
}

type Row struct {
	value []Cell
}

func NewRow() Row {
	r := Row{}
	for i := 0; i < 4; i++ {
		r.value = append(r.value, NewEmptyCell())
	}
	return r
}

func (r Row) ToStrings() []string {
	result := make([]string, 0)
	for _, cell := range r.value {
		v := cell.value
		result = append(result, v)
	}
	return result
}

func (r Row) ResultStringLine() string {
	return strings.Join(r.ToStrings(), " ")
}

func (r Row) ToInt() (int, error) {
	// TODO: ここのバリデーションがあまい
	s := strings.Join(r.ToStrings(), "")
	s = strings.TrimSpace(s)
	return strconv.Atoi(s)
}

type Cells struct {
	value []Row
}

func NewCells() Cells {
	c := Cells{}
	for i := 0; i < 3; i++ {
		c.value = append(c.value, NewRow())
	}
	return c
}
