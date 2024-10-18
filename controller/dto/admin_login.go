package dto

type AdminLoginData struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	CaptchaID string `json:"captcha_id" form:"captcha_id" binding:"required"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required;"`
}
