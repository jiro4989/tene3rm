package service

import (
	"math"

	"github.com/jiro4989/tene3rm/domain/model"
	"github.com/jiro4989/tene3rm/infra"
)

type MathService struct{}

func NewMathService() MathService {
	return MathService{}
}

func (s MathService) SimpleOperations(oprg, arg, brg infra.RandomGenerator) (int, int, int, string) {
	ops := []model.Operator{
		&model.PlusOperator{},
		&model.MinusOperator{},
		&model.MultiOperator{},
	}
	opi := oprg.Intn(len(ops))
	op := ops[opi]

	// 0にならないようにする
	a := int(math.Max(float64(arg.Intn(10)), 1))
	b := int(math.Max(float64(brg.Intn(10)), 1))

	return op.Do(a, b), a, b, op.Op()
}
