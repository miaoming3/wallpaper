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

type AdminSearch struct {
	Email    string `json:"email" form:"email" uri:"email"`
	Phone    string `json:"phone" form:"phone" uri:"phone"`
	Username string `json:"username" form:"username" uri:"username" `
	Status   string `json:"status" form:"status" uri:"status"`
}

type AdminDel struct {
	ID int `json:"id" form:"id" uri:"id" binding:"required"`
}
