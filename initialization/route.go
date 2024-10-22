package initialization

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/controller/v1"
	docs "github.com/miaoming3/wallpaper/docs"
	"github.com/miaoming3/wallpaper/global"
	"github.com/miaoming3/wallpaper/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Use(gin.Recovery(), middleware.LoadPageMiddleware(), sessions.Sessions("session", cookie.NewStore([]byte("1245"))))
	r.Static("/static", "./static")
	r.LoadHTMLGlob(global.SysConfig.Template)
	api := r.Group("/api")
	{
		adminRouter(api)
		categoryRouter(api)
	}

	return r
}

func adminRouter(r *gin.RouterGroup) {
	adminController := v1.NewAdminController()
	v1Router := r.Group("/v1")
	{
		v1Router.POST("/admin/login", adminController.Login)
	}

}
func categoryRouter(r *gin.RouterGroup) {
	categoryController := v1.NewCategoryController()
	v1Router := r.Group("/v1")
	{
		v1Router.GET("/category/index", categoryController.Index)
		v1Router.POST("category/save", categoryController.Save)
		v1Router.PUT("/category/update", categoryController.Update)
		v1Router.DELETE("/category/del", categoryController.Delete)
	}
}
