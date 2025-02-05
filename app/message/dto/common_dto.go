package dto

type Page struct {
	Page     int `uri:"page" form:"page" query:"page" json:"page"`
	PageSize int `uri:"pageSize" form:"pageSize" query:"pageSize" json:"pageSize"`
}
