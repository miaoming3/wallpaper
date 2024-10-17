package initialization

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/global"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery(), sessions.Sessions("session", cookie.NewStore([]byte("1245"))))
	r.Static("/static", "./static")
	r.LoadHTMLGlob(global.SysConfig.Template)
	return r
}
