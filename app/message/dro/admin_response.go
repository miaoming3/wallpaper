package dro

type LoginResponse struct {
	Uid   uint   `json:"uid"`
	Token string `json:"token"`
}
