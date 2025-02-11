package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/app/controller/v1"
	"github.com/miaoming3/wallpaper/app/middleware"
)

func AdminRouter(admin *gin.RouterGroup) {
	adminController := v1.NewAdminController()
	admin.POST("login", adminController.Login)
	adminAuth := admin.Group("admin")
	{
		adminAuth.Use(middleware.AuthMiddleware())
		adminAuth.GET("info", adminController.Info)
		adminAuth.POST("change/password", adminController.ChangePassword)
		adminAuth.GET("index", middleware.Page(), adminController.Index)
		adminAuth.POST("created", adminController.Created)
		adminAuth.GET("detail", adminController.GetAdminInfo)
		adminAuth.POST("updated", adminController.Updated)
		adminAuth.DELETE("delete", adminController.Delete)
		adminAuth.POST("logout", adminController.Logout)
		adminAuth.POST("change/info", adminController.ChangeInfo)
	}

}
