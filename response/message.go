package response

const (
	ACCESSERROR uint = 500
	CLIENTERROR uint = 400
	//1100- 1199 验证码的错误码
	CAPTCHAERROR uint = 1100
	//管理员错误码 1300 - 1600
	ADMINORPASSWORD uint = 1300
)

var messageString = map[uint]string{
	ACCESSERROR:     "服务器内部错误",
	CLIENTERROR:     "客户端错误",
	CAPTCHAERROR:    "验证码服务",
	ADMINORPASSWORD: "用户名或密码错误",
}
