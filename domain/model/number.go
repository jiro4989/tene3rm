package model

import "strconv"

type Number struct {
	value int
}

func NewNumber(n int) Number {
	return Number{
		value: n,
	}
}

func NewNumberWithText(s string) (Number, error) {
	t := NewText(s)
	n, err := strconv.Atoi(t.value)
	if err != nil {
		return Number{}, err
	}
	return NewNumber(n), nil
}

func (n Number) Equal(n2 Number) bool {
	return n.value == n2.value
}

func (n Number) Value() int {
	return n.value
}
