package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/message"
	"github.com/miaoming3/wallpaper/app/response"
	"net/http"
	"strconv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		uid := context.Request.Header.Get("uid")
		if token == "" || uid == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ApiError(message.Unauthorized, nil))
			return
		}
		adminToken, err := global.RedisClien.Get(context, fmt.Sprintf("admin:token:%v", uid)).Result()
		if err != nil || adminToken != token {
			context.AbortWithStatusJSON(http.StatusUnauthorized, response.ApiError(message.Unauthorized, nil))
			return
		}
		uidInt, _ := strconv.Atoi(uid)
		context.Set("uid", uidInt)
		context.Next()

	}
}
