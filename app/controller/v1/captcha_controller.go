package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController struct {
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

func (ca *CaptchaController) Captcha(c *gin.Context) {
	digit := base64Captcha.NewDriverDigit(40, 100, 5, 0.7, 80)
	global.Captcha = base64Captcha.NewCaptcha(digit, base64Captcha.DefaultMemStore)
	id, image, _, err := global.Captcha.Generate()
	if err != nil {
		response.ApiError(message.ACCESSERROR, err)
		return
	}
	response.ApiSuccess(map[string]string{
		"id":      id,
		"captcha": image,
	})
	return
}
