package initialization

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/app/controller/v1"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/response"
	"github.com/miaoming3/wallpaper/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path/filepath"
)

func InitRoute() *gin.Engine {
	gin.SetMode(global.SysConfig.Model)
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, response.ApiError(http.StatusNotFound, nil))
	})
	r.Static(global.SysConfig.Dir, filepath.Join("./", global.SysConfig.Dir))

	uploadController := v1.NewUploadController()
	captchaController := v1.NewCaptchaController()
	r.GET("/captcha", captchaController.Captcha)
	r.POST("uploads/file", uploadController.UploadFile)
	r.POST("uploads/must", uploadController.UploadFileMust)
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
