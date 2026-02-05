package service

import "github.com/jiro4989/tene3rm/domain/model"

type YesNoService struct{}

func NewYesNoService() YesNoService {
	return YesNoService{}
}

func (s YesNoService) JudgeYesNo(input string) bool {
	return judge(input, "yes")
}

func (s YesNoService) JudgeYesNoDenial(input string) bool {
	return judge(input, "no")
}

func (s YesNoService) JudgeYesNoJapanese(input string) bool {
	return judge(input, "はい")
}

func judge(input string, want string) bool {
	i := model.NewUserInputText(input)
	w := model.NewUserInputText(want).Prefixes()
	return w.Contains(i)
}
