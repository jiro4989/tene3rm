package model

type Operator interface {
	Do(Number, Number) Number
	Op() string
}

var (
	_ Operator = &PlusOperator{}
	_ Operator = &MinusOperator{}
	_ Operator = &MultiOperator{}
)

type PlusOperator struct{}

func (o *PlusOperator) Do(a, b Number) Number {
	return a.Plus(b)
}

func (o *PlusOperator) Op() string {
	return "+"
}

type MinusOperator struct{}

func (o *MinusOperator) Do(a, b Number) Number {
	return a.Minus(b)
}

func (o *MinusOperator) Op() string {
	return "-"
}

type MultiOperator struct{}

func (o *MultiOperator) Do(a, b Number) Number {
	return a.Multi(b)
}

func (o *MultiOperator) Op() string {
	return "*"
}
