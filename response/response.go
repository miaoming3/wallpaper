package response

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Err  error       `json:"error,omitempty"`
	Page *PageData   `json:"page,omitempty"`
}

type PageData struct {
	Total int64 `json:"total"`
	Size  int   `json:"size"`
	Page  int   `json:"page"`
	Next  bool  `json:"next"`
}

func (ar ApiResponse) Error() string {
	return ar.Err.Error()
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
		Code: code,
		Data: []string{},
		Msg:  GetMessage(code),
		Err:  err,
	}
}

func ApiPageSuccess(data interface{}, total int64, page int, size int, next bool) *ApiResponse {
	if data == nil || reflect.ValueOf(data).IsNil() {
		data = []string{}
	}
	return &ApiResponse{
		Code: 0,
		Data: data,
		Msg:  "success",
		Page: &PageData{
			Total: total,
			Size:  size,
			Page:  page,
			Next:  next,
		},
	}
}

func Response(c *gin.Context, response *ApiResponse) {
	c.JSON(http.StatusOK, response)
}

func GetMessage(code uint) string {
	if message, exists := messageString[code]; exists {
		return message
	}
	return "服务器内部错误"
}
