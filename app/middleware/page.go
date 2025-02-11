package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type page struct {
	Page     int `json:"page" form:"page" uri:"page"`
	PageSize int `json:"page_size" form:"page_size" uri:"page_size"`
}

func Page() gin.HandlerFunc {

	return func(context *gin.Context) {
		var pageData page
		if err := context.ShouldBind(&pageData); err != nil {
			context.AbortWithStatusJSON(http.StatusNoContent, nil)
			return
		}
		if pageData.PageSize == 0 {
			pageData.PageSize = 15
		}
		if pageData.Page == 0 {
			pageData.Page = 1
		}
		context.Set("page", pageData.Page)
		context.Set("pageSize", pageData.PageSize)

	}
}
