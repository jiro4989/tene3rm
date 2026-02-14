package usecase

import "github.com/jiro4989/tene3rm/domain"

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
	i := domain.NewUserInputText(input)
	w := domain.NewUserInputText(want).Prefixes()
	return i.In(w)
}
