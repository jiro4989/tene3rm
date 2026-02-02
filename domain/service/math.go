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

	// 0にならないようにする
	n1 := model.NewNumber(1)
	a := model.NewNumber(arg.Intn(10)).Max(n1)
	b := model.NewNumber(brg.Intn(10)).Max(n1)

	return op.Do(a, b), a, b, op.Op()
}
