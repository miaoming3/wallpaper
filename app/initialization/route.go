package initialization

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/app/controller/v1"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/initialization/routers"
)

type Controller struct {
	Upload v1.UploadController
	Admin  v1.AdminController
}

func InitRoute() *gin.Engine {
	gin.SetMode(global.SysConfig.Model)
	r := gin.Default()
	routers.CommonRouter(r)
	api := r.Group("/api/v1")
	{
		routers.AdminRouter(api)
	}

	return r
}
