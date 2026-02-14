package usecase

import (
	"math"

	"github.com/jiro4989/tene3rm/domain"
	"github.com/jiro4989/tene3rm/infra"
)

type SimpleOperationUsecase struct{}

func NewSimpleOperationUsecase() SimpleOperationUsecase {
	return SimpleOperationUsecase{}
}

func (s SimpleOperationUsecase) Execute(oprg, arg, brg infra.RandomGenerator) (int, int, int, string) {
	ops := []domain.Operator{
		&domain.PlusOperator{},
		&domain.MinusOperator{},
		&domain.MultiOperator{},
	}
	opi := oprg.Intn(len(ops))
	op := ops[opi]

	// 0にならないようにする
	a := int(math.Max(float64(arg.Intn(10)), 1))
	b := int(math.Max(float64(brg.Intn(10)), 1))

	return op.Do(a, b), a, b, op.Op()
}
