package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperator(t *testing.T) {
	a := assert.New(t)

	p := &PlusOperator{}
	a.Equal(4, p.Do(1, 3))
	a.Equal("+", p.Op())

	m := &MinusOperator{}
	a.Equal(-2, m.Do(1, 3))
	a.Equal("-", m.Op())

	m2 := &MultiOperator{}
	a.Equal(3, m2.Do(1, 3))
	a.Equal("*", m2.Op())
}
