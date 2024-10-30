package dto

type MenuIndex struct {
	Name   string `uri:"name" form:"name" query:"page"`
	Status string ` uri:"status" form:"status" query:"status"`
	Pid    string `uri:"pid" form:"pid" query:"pid"`
}
