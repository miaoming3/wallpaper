package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/controller/dto"
	"github.com/miaoming3/wallpaper/response"
)

func LoadPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commonPage dto.Page
		if err := c.ShouldBind(&commonPage); err != nil {
			response.Response(c, response.ApiError(response.CLIENTERROR, err))
			return
		}
		if commonPage.PageSize == 0 {
			commonPage.PageSize = 20
		}
		if commonPage.Page < 1 {
			commonPage.Page = 1
		}
		c.Set("page", commonPage.Page)
		c.Set("pageSize", commonPage.PageSize)
		c.Next()
	}

}
