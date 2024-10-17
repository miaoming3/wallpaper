package response

import (
	"github.com/gin-gonic/gin"
	"github.com/miaoming3/wallpaper/global"
	"net/http"
)

type ApiResponse struct {
	Code  uint        `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
	Page  *PageData   `json:"page,omitempty"`
}

type PageData struct {
	Total int64 `json:"total"`
	Size  uint  `json:"size"`
	Next  bool  `json:"next"`
}

func ApiSuccess(data interface{}) *ApiResponse {
	if data == nil {
		data = []string{}
	}

	return &ApiResponse{
		Code: 0,
		Data: data,
		Msg:  "success",
	}
}

func ApiError(code uint, err error) *ApiResponse {

	return &ApiResponse{
		Code:  code,
		Data:  []string{},
		Msg:   getMessage(code),
		Error: ShowError(err),
	}
}
func ShowError(err error) string {

	if global.SysConfig.Dev && err != nil {
		return err.Error()
	}
	return ""
}

func ApiPageSuccess(data interface{}, total int64, next bool) *ApiResponse {
	if data == nil {
		data = []string{}
	}
	return &ApiResponse{
		Code: 0,
		Data: data,
		Msg:  "success",
		Page: &PageData{
			Total: total,
			Size:  20,
			Next:  next,
		},
	}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiSuccess(data))
}
func Error(c *gin.Context, code uint, err error) {
	c.JSON(http.StatusOK, ApiError(code, err))
}
func PageSuccess(c *gin.Context, data interface{}, total int64, next bool) {
	c.JSON(http.StatusOK, ApiPageSuccess(data, total, next))
}

func getMessage(code uint) string {
	stringMessage := messageString[code]
	if stringMessage != "" {
		return stringMessage
	}
	return "服务器内部错误"
}
