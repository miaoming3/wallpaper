package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/core/server"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/message/dto"
	"github.com/miaoming3/wallpaper/app/response"
)

type UploadController struct {
}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (uc *UploadController) UploadFile(c *gin.Context) {
	var file dto.UploadFile
	if err := c.ShouldBind(&file); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewUploadServer().UploadOneFile(c, &file))
}

func (uc *UploadController) UploadFileMust(c *gin.Context) {
	var file dto.UploadFileMust
	if err := c.ShouldBind(&file); err != nil {
		response.Response(c, response.ApiError(message.CLIENTERROR, err))
		return
	}
	response.Response(c, server.NewUploadServer().UploadsFileMust(c, &file))
}
