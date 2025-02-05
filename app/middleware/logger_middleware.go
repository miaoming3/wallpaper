package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/global"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.SugarLogger.Info()
		c.Next()
	}

}
