package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/app/controller/v1"
)

func AdminRouter(admin *gin.RouterGroup) {
	adminController := v1.NewAdminController()
	admin.POST("login", adminController.Login)

}
