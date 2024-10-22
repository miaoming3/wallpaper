package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"gorm.io/gorm"
)

type CategoryStatus uint

const (
	StatusOn  CategoryStatus = iota + 1
	StatusOff CategoryStatus = 2
)

type Category struct {
	gorm.Model
	Name   string `gorm:"name;type:varchar(255) "`
	Pid    uint
	Sort   uint
	Status CategoryStatus `gorm:"status;type:tinyint(3);default:1"`
}

func (ca *Category) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "catgegory")
}
