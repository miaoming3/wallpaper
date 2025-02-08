package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/global"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			global.SugarLogger.Info(
				fmt.Sprintf("请求方式: %v ", c.Request.Method),
				fmt.Sprintf("请求地址: %v ", c.Request.RequestURI),
				fmt.Sprintf("请求路劲: %v ", c.Request.URL.Path),
				fmt.Sprintf("客户端ip: %v ", c.ClientIP()),
				fmt.Sprintf("请求参数: %v ", c.Request.PostForm))
		}
		c.Next()
	}

}
