package middleware

import (
	"github.com/gin-gonic/gin"
)

func CoreMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有源访问，也可以指定具体的源如 "http://example.com"
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, uid, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 对于OPTIONS请求，返回状态码204 No Content，不处理请求体。
			return
		}
		c.Next() // 继续处理请求。重要！否则请求将被拦截。
	}
}
