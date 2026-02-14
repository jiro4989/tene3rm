package usecase

import "github.com/jiro4989/tene3rm/domain"

type YesNoUsecase struct{}

func NewYesNoUsecase() YesNoUsecase {
	return YesNoUsecase{}
}

func (s YesNoUsecase) JudgeYesNo(input string) bool {
	return judge(input, "yes")
}

func (s YesNoUsecase) JudgeYesNoDenial(input string) bool {
	return judge(input, "no")
}

func (s YesNoUsecase) JudgeYesNoJapanese(input string) bool {
	return judge(input, "はい")
}

func judge(input string, want string) bool {
	i := domain.NewUserInputText(input)
	w := domain.NewUserInputText(want).Prefixes()
	return i.In(w)
}
