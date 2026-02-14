package column

import (
	"strconv"
	"strings"
)

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
