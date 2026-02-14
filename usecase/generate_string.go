package usecase

import (
	"strings"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/jiro4989/tene3rm/repo"
)

type GenerateStringUsecase struct {
	randGen    infra.RandomGenerator
	stringRepo repo.StringRepo
}

func NewGenerateStringUsecase(rg infra.RandomGenerator, sr repo.StringRepo) GenerateStringUsecase {
	return GenerateStringUsecase{
		randGen:    rg,
		stringRepo: sr,
	}
}

func (g *GenerateStringUsecase) Execute(length int, filename string) (string, error) {
	s := g.randomString(length)
	return s, g.stringRepo.Save(filename, s)
}

func (g *GenerateStringUsecase) randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		n := g.randGen.Intn(len(charset))
		r := charset[n]
		sb.WriteByte(r)
	}
	s := sb.String()

	return s
}
