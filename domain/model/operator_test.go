package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperator(t *testing.T) {
	a := assert.New(t)

	n1 := NewNumber(1)
	n3 := NewNumber(3)
	n4 := NewNumber(4)
	nm2 := NewNumber(-2)

	p := &PlusOperator{}
	a.Equal(n4, p.Do(n1, n3))
	a.Equal("+", p.Op())

	m := &MinusOperator{}
	a.Equal(nm2, m.Do(n1, n3))
	a.Equal("-", m.Op())

	m2 := &MultiOperator{}
	a.Equal(n3, m2.Do(n1, n3))
	a.Equal("*", m2.Op())

	a.Equal(n3, n3.Max(n1))
	a.Equal(n4, n3.Max(n4))
}
