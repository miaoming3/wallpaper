package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/app/controller/v1"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/middleware"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path/filepath"
)

func CommonRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, response.ApiError(http.StatusNotFound, nil))
	})
	r.Static(global.SysConfig.Dir, filepath.Join("./", global.SysConfig.Dir))
	r.Use(middleware.CoreMiddleware(), middleware.LoggerMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"
	BaseRouters(r)
	UploadRouter(r)

}

func BaseRouters(r *gin.Engine) {
	commonController := v1.NewCommonController()

	r.GET("captcha", commonController.Captcha)

}
