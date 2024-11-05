package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func LoadPageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		if pageString, exists := c.GetQuery("page"); exists {
			page, _ = strconv.Atoi(pageString)
			if page <= 0 {
				page = 1
			}

		}
		pageSize := 20
		if pageString, exists := c.GetQuery("pageSize"); exists {
			pageSize, _ = strconv.Atoi(pageString)
			if pageSize <= 0 {
				pageSize = 20
			}

		}
		c.Set("page", page)
		c.Set("pageSize", pageSize)
		c.Next()
	}

}
