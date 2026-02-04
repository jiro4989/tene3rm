package model

import "fmt"

// ColumnAddition は筆算を管理する。
type ColumnAddition struct {
	a, b int
	op   string
}

func NewColumnAddition(a, b int) (ColumnAddition, error) {
	if a < 10 {
		return ColumnAddition{}, fmt.Errorf("must be 10 <= a: %d", a)
	}
	if 99 < a {
		return ColumnAddition{}, fmt.Errorf("must be a <= 99: %d", a)
	}
	if b < 10 {
		return ColumnAddition{}, fmt.Errorf("must be 10 <= b: %d", b)
	}
	if 99 < b {
		return ColumnAddition{}, fmt.Errorf("must be b <= 99: %d", b)
	}

	return ColumnAddition{
		a:  a,
		b:  b,
		op: "x",
	}, nil
}

// MultiplyOnesPlace は b の1の位の計算結果を返す。
func (c ColumnAddition) MultiplyOnesPlace() int {
	return c.a * int(c.b%10)
}

// MultiplyTensPlace は b の10の位の計算結果を返す。
func (c ColumnAddition) MultiplyTensPlace() int {
	return c.a * int(c.b/10)
}

// Multiply は計算結果を返す。
func (c ColumnAddition) Multiply() int {
	return c.a * c.b
}
