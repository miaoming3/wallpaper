package dto

type CommonAdmin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type AdminLogin struct {
	*CommonAdmin
	*AdminCaptcha
}

type AdminCaptcha struct {
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	CaptchaId string `json:"id" form:"id" binding:"required"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password" form:"old_password" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

type ChangeAdminInfo struct {
	*CommonAdmin
	Email  string `json:"email" form:"email" binding:"required,email"`
	Phone  string `json:"phone" form:"phone" binding:"required,isPhone"`
	Avatar string `json:"avatar" form:"avatar"`
}
