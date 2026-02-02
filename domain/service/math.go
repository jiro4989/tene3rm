package service

import (
	"github.com/jiro4989/tene3rm/domain/model"
	"github.com/jiro4989/tene3rm/infra"
)

type MathService struct{}

func NewMathService() MathService {
	return MathService{}
}

func (s MathService) SimpleOperations(oprg, arg, brg infra.RandomGenerator) (model.Number, model.Number, model.Number, string) {
	ops := []model.Operator{
		&model.PlusOperator{},
		&model.MinusOperator{},
		&model.MultiOperator{},
	}
	opi := oprg.Intn(len(ops))
	op := ops[opi]

	a := arg.Intn(10)
	b := brg.Intn(10)

	// 0にならないようにする
	if a < 1 {
		a = 1
	}
	if b < 1 {
		b = 1
	}

	nr := model.NewNumber(op.Do(a, b))
	na := model.NewNumber(a)
	nb := model.NewNumber(b)
	return nr, na, nb, op.Op()
}
