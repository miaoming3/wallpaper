package dto

type AdminLoginResponse struct {
	Token      string `json:"token"`
	Expression string `json:"expression"`
	Username   string `json:"username"`
	UID        string `json:"uid"`
}
