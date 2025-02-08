package response

import (
	"github.com/miaoming3/wallpaper/app/message"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type APi struct {
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

func (ar *APi) Error() string {
	return ar.Err.Error()
}
func ApiSuccess(data interface{}) *APi {
	if data == nil {
		data = []string{}
	}

	return &APi{
		Code: 0,
		Data: data,
		Msg:  "success",
	}
}

func ApiError(code uint, err error) *APi {
	return &APi{
		Code: code,
		Data: []string{},
		Msg:  message.GetMessage(code),
		Err:  err,
	}
}

func ApiPageSuccess(data interface{}, total int64, page int, size int, next bool) *APi {
	if data == nil || reflect.ValueOf(data).IsNil() {
		data = []string{}
	}
	return &APi{
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

func Response(c *gin.Context, response *APi) {
	c.JSON(http.StatusOK, response)
}
