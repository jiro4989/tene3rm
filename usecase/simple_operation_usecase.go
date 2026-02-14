package usecase

import (
	"math"

	"github.com/jiro4989/tene3rm/domain"
	"github.com/jiro4989/tene3rm/infra"
)

type SimpleOperationUsecase struct {
	randGen infra.RandomGenerator
}

func NewSimpleOperationUsecase(rg infra.RandomGenerator) SimpleOperationUsecase {
	return SimpleOperationUsecase{
		randGen: rg,
	}
}

func (s SimpleOperationUsecase) Execute() (int, int, int, string) {
	ops := []domain.Operator{
		&domain.PlusOperator{},
		&domain.MinusOperator{},
		&domain.MultiOperator{},
	}
	opi := s.randGen.Intn(len(ops))
	op := ops[opi]

	// 0にならないようにする
	a := int(math.Max(float64(s.randGen.Intn(10)), 1))
	b := int(math.Max(float64(s.randGen.Intn(10)), 1))

	return op.Do(a, b), a, b, op.Op()
}
