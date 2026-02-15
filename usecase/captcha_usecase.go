package usecase

import (
	"io"

	"github.com/steambap/captcha"
)

type CaptchaUsecase struct{}

func NewCaptchaUsecase() CaptchaUsecase {
	return CaptchaUsecase{}
}

func (c CaptchaUsecase) Execute(w io.Writer) (string, error) {
	// TODO: writer を渡すのと、Captcha を内包してるの良くない気がする
	data, _ := captcha.New(150, 50)
	if err := data.WriteImage(w); err != nil {
		return "", err
	}
	return data.Text, nil
}
