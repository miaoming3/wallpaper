package initialization

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var customizeValidate = map[string]validator.Func{
	"isTags": IsTage,
}

func InitValidate() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for k, val := range customizeValidate {
			_ = v.RegisterValidation(k, val)
		}
	}

}

func IsTage(fl validator.FieldLevel) bool {
	return false
}
