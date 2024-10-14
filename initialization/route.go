package initialization

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/miaoming3/wallpaper/api/v1"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "admin/login.html", nil)
	})
	api := r.Group("/admin")
	{
		apiRouter(api)
	}

	return r
}

func apiRouter(r *gin.RouterGroup) {
	r.POST("/login", v1.Login)
}
