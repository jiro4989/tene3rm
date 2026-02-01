package service

import "github.com/jiro4989/tene3rm/domain/model"

type TextService struct{}

func NewTextService() TextService {
	return TextService{}
}

func (s TextService) JudgeYesNo(input string) bool {
	return judge(input, "yes")
}

func (s TextService) JudgeYesNoDenial(input string) bool {
	return judge(input, "no")
}

func (s TextService) JudgeYesNoJapanese(input string) bool {
	return judge(input, "はい")
}

func judge(input string, want string) bool {
	i := model.NewText(input)
	w := model.NewText(want).Prefixes()
	return w.Contains(i)
}
