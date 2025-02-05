package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/response"
)

type CommonController struct {
}

func NewCommonController() *CommonController {

	return &CommonController{}
}

func (base *CommonController) Captcha(c *gin.Context) {
	id, image, _, err := global.Captcha.Generate()
	if err != nil {
		response.Response(c, response.ApiError(message.ACCESSERROR, err))
		return
	}
	response.Response(c, response.ApiSuccess(map[string]string{
		"id":      id,
		"captcha": image,
	}))
}
