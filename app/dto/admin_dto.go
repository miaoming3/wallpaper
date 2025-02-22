package dto

type AdminUserName struct {
	Username string `json:"username" form:"username" binding:"required,min=6,max=32"`
}
type AdminPassword struct {
	Password string `json:"password" form:"password" binding:"required,min=6,max=32"`
}
type AdminCaptcha struct {
	Captcha   string `json:"captcha" form:"captcha" binding:"required,len=4"`
	CaptchaId string `json:"id" form:"id" binding:"required"`
}

type AdminLogin struct {
	*AdminUserName
	*AdminPassword
	*AdminCaptcha
}

type ChangePassword struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required,min=6,max=32"`
	*AdminPassword
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,eqfield=Password"`
}

type ChangeAdminInfo struct {
	*AdminUserName
	Email  string `json:"email" form:"email" binding:"required,email"`
	Phone  string `json:"phone" form:"phone" binding:"required,isPhone,len=11,numeric"`
	Avatar string `json:"avatar" form:"avatar"`
}

type AdminSearch struct {
	Email    string `json:"email" form:"email" uri:"email,omitempty"`
	Phone    string `json:"phone" form:"phone" uri:"phone,omitempty"`
	Username string `json:"username" form:"username" uri:"username,omitempty" `
	Status   string `json:"status" form:"status" uri:"status,omitempty"`
}

type AdminDel struct {
	ID int `json:"id" form:"id" uri:"id" binding:"required"`
}

type AdminUpdate struct {
	*AdminDel
	*AdminPassword
	*ChangeAdminInfo
	Status string `json:"status" form:"status" uri:"status"`
}

type AdminCreated struct {
	*AdminPassword
	*ChangeAdminInfo
	Status string `json:"status" form:"status" uri:"status"`
}
