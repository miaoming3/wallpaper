package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/core/server"
	response2 "github.com/miaoming3/wallpaper/http/response"
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
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}
	response2.Response(c, uc.UploadOneFile(c, &file))
}

func (uc *UploadController) UploadFileMust(c *gin.Context) {
	var file dto.UploadFileMust
	if err := c.ShouldBind(&file); err != nil {
		response2.Response(c, response2.ApiError(response2.CLIENTERROR, err))
		return
	}
	response2.Response(c, uc.UploadsFileMust(c, &file))
}
