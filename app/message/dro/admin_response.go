package dro

type LoginResponse struct {
	Uid   uint   `json:"uid"`
	Token string `json:"token"`
}

type AdminInfo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	ImgUrl    string `json:"img_url"`
	CreatedAt string `json:"created_at"`
}
