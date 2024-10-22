package dto

type CategoryIndex struct {
	Name   string `uri:"name" form:"name" query:"page"`
	Status string ` uri:"status" form:"status" query:"status"`
	Pid    string `uri:"pid" form:"pid" query:"pid"`
}

type SaveCategory struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Pid    uint   `json:"pid" form:"pid"`
	Sort   uint   `json:"sort" form:"sort"`
	Status uint   `json:"status" form:"Status"`
}
