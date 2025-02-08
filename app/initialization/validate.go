package initialization

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var customizeValidate = map[string]validator.Func{
	"isTags":  IsTage,
	"isPhone": IsPhone,
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

func IsPhone(fl validator.FieldLevel) bool {
	return true
}
