package domain

import "fmt"

// TwoDigitMultiplyColumnAddition は2桁同士の乗算の筆算を管理する。
type TwoDigitMultiplyColumnAddition struct {
	a, b int
	op   string
}

func NewTwoDigitMultiplyColumnAddition(a, b int) (TwoDigitMultiplyColumnAddition, error) {
	if a < 10 {
		return TwoDigitMultiplyColumnAddition{}, fmt.Errorf("must be 10 <= a: %d", a)
	}
	if 99 < a {
		return TwoDigitMultiplyColumnAddition{}, fmt.Errorf("must be a <= 99: %d", a)
	}
	if b < 10 {
		return TwoDigitMultiplyColumnAddition{}, fmt.Errorf("must be 10 <= b: %d", b)
	}
	if 99 < b {
		return TwoDigitMultiplyColumnAddition{}, fmt.Errorf("must be b <= 99: %d", b)
	}

	return TwoDigitMultiplyColumnAddition{
		a:  a,
		b:  b,
		op: "x",
	}, nil
}

// MultiplyOnesPlace は b の1の位の計算結果を返す。
func (c TwoDigitMultiplyColumnAddition) MultiplyOnesPlace() int {
	return c.a * int(c.b%10)
}

// MultiplyTensPlace は b の10の位の計算結果を返す。
func (c TwoDigitMultiplyColumnAddition) MultiplyTensPlace() int {
	return c.a * int(c.b/10)
}

// Multiply は計算結果を返す。
func (c TwoDigitMultiplyColumnAddition) Multiply() int {
	return c.a * c.b
}

// Equal は途中計算も含めて正しいか判定する。
func (c TwoDigitMultiplyColumnAddition) Equal(n1, n2, n3 int) bool {
	r1 := n1 == c.MultiplyOnesPlace()
	r2 := n2 == c.MultiplyTensPlace()
	r3 := n3 == c.Multiply()
	return r1 && r2 && r3
}
