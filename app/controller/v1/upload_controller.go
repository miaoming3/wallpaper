package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/server"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/message/dto"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/mojocn/base64Captcha"
)

type UploadController struct {
	server.UploadServer
}

func NewUploadController() *UploadController {
	return &UploadController{server.NewUploadServer()}
}

func (uc *UploadController) UploadFile(c *gin.Context) {
	var file dto.UploadFile
	if err := c.ShouldBind(&file); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, uc.UploadOneFile(c, &file))
}

func (uc *UploadController) UploadFileMust(c *gin.Context) {
	var file dto.UploadFileMust
	if err := c.ShouldBind(&file); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, uc.UploadsFileMust(c, &file))
}

func (cc *UploadController) Generate(c *gin.Context) {
	global.Captcha = base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)
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
