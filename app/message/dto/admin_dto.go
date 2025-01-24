package dto

type AdminLogin struct {
	Username  string `json:"username" form:"username" binding:"required,isTags"`
	Password  string `json:"password" form:"password" binding:"required"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	CaptchaId string `json:"id" form:"id" binding:"required"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password" form:"old_password" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"confirm_password,eqfield:confirm_password"`
}
