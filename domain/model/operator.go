package model

type Operator interface {
	Do(int, int) int
	Op() string
}

var (
	_ Operator = &PlusOperator{}
	_ Operator = &MinusOperator{}
	_ Operator = &MultiOperator{}
)

type PlusOperator struct{}

func (o *PlusOperator) Do(a, b int) int {
	return a + b
}

func (o *PlusOperator) Op() string {
	return "+"
}

type MinusOperator struct{}

func (o *MinusOperator) Do(a, b int) int {
	return a - b
}

func (o *MinusOperator) Op() string {
	return "-"
}

type MultiOperator struct{}

func (o *MultiOperator) Do(a, b int) int {
	return a * b
}

func (o *MultiOperator) Op() string {
	return "*"
}
