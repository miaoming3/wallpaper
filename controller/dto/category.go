package dto

type CategoryIndex struct {
	Page     int    `uri:"page" form:"page" query:"page"`
	Name     string `uri:"name" form:"name" query:"page"`
	Status   string ` uri:"status" form:"status" query:"status"`
	Pid      string `uri:"pid" form:"pid" query:"pid"`
	PageSize int    `uri:"pageSize" form:"pageSize" query:"pageSize"`
}
