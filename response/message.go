package response

var (
	ACCESSERROR uint = 500
	CLIENTERROR uint = 400
)

var messageString = map[uint]string{
	ACCESSERROR: "服务器内部错误",
}
