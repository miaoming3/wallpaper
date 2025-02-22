package message

import "github.com/go-playground/validator/v10"

// 存放GetMessages()方法
type Validator interface {
	GetMessages() ValidatorMessages
}

// 校验信息
type ValidatorMessages = map[string]string

var ValidateString = ValidatorMessages{
	"Password.eqfield":         "两次密码不一致",
	"Password.required":        "密码不能为空",
	"Password.minlen":          "密码长度在6-32位之间",
	"Password.maxlen":          "密码长度在6-32位之间",
	"Password.regex":           "密码验证规则错误",
	"ConfirmPassword.required": "确认密码不能为空",
	"ConfirmPassword.minlen":   "确认密码长度在6-32位之间",
	"ConfirmPassword.maxlen":   "确认密码长度在6-32位之间",
	"ConfirmPassword.regex":    "确认密码验证规则错误",
	"ConfirmPassword.eqfield":  "两次密码不一致",
	"OldPassword.required":     "旧密码不能为空",
	"OldPassword.min":          "旧密码长度在6-32位之间",
	"OldPassword.max":          "旧密码长度在6-32位之间",
	"OldPassword.regex":        "旧密码验证规则错误",
	"Captcha.required":         "验证码不能为空",
	"Captcha.len":              "验证码的长度等于4",
	"Email.required":           "邮箱不能为空",
	"Email.email":              "邮箱规则错误",
	"Username.required":        "用户名不能为空",
	"Username.regex":           "用户名验证规则错误",
	"Username.minlen":          "用户名最小长度为6位",
	"Username.maxlen":          "用户名最大长度为32位",
	"Photo.required":           "电话号码不能为空",
	"Photo.len":                "电话号码的长度为11位",
	"Photo.numeric":            "电话号码位纯数字",
}

// GetErrorMsg方法， 获取错误信息
func GetErrorMsg(err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		for _, v := range err.(validator.ValidationErrors) {
			if message, exists := ValidateString[v.Field()+"."+v.Tag()]; exists {
				return message
			}
			return v.Error()
		}
	}
	return "Parameter error"
}
