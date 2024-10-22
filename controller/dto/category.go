package dto

import "github.com/miaoming3/wallpaper/models"

type CategoryIndex struct {
	Name   string `uri:"name" form:"name" query:"page"`
	Status string ` uri:"status" form:"status" query:"status"`
	Pid    string `uri:"pid" form:"pid" query:"pid"`
}

type SaveCategory struct {
	Name   string                `json:"name" form:"name" binding:"required"`
	Pid    uint                  `json:"pid" form:"pid" default:"0"`
	Sort   uint                  `json:"sort" form:"sort" default:"50"`
	Status models.CategoryStatus `json:"status" form:"Status" default:"1"`
}

type UpdateCategory struct {
	ID uint `json:"id" form:"id" binding:"required"`
	SaveCategory
}

type DeleteCategory struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
