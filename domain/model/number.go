package model

import (
	"strconv"
)

type Number struct {
	value int
}

func NewNumber(n int) Number {
	return Number{
		value: n,
	}
}

func NewNumberWithUserInputText(t UserInputText) (Number, error) {
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

func (n Number) Plus(n2 Number) Number {
	return NewNumber(n.value + n2.value)
}

func (n Number) Minus(n2 Number) Number {
	return NewNumber(n.value - n2.value)
}

func (n Number) Multi(n2 Number) Number {
	return NewNumber(n.value * n2.value)
}

func (n Number) Max(n2 Number) Number {
	if n.value < n2.value {
		return n2
	}
	return n
}
