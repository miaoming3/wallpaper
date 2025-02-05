package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/controller/v1"
)

func UploadRouter(r *gin.Engine) {
	uploadController := v1.NewUploadController()
	r.POST("uploads/file", uploadController.UploadFile)
	r.POST("uploads/must", uploadController.UploadFileMust)
}
