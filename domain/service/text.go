package service

import "github.com/jiro4989/tene3rm/domain/model"

type TextService struct{}

func NewTextService() TextService {
	return TextService{}
}

func (s TextService) JudgeYesNo(input string) bool {
	i := model.NewText(input)
	want := model.NewText("yes").Prefixes()
	return want.Contain(i)
}

func (s TextService) JudgeYesNoDenial(input string) bool {
	i := model.NewText(input)
	want := model.NewText("no").Prefixes()
	return want.Contain(i)
}

func (s TextService) JudgeYesNoJapanese(input string) bool {
	i := model.NewText(input)
	want := model.NewText("はい").Prefixes()
	return want.Contain(i)
}
