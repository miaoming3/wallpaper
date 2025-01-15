package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/http/global"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string `gorm:"name;type:varchar(255) "`
	Pid    uint
	Sort   uint
	Status uint `gorm:"status;type:tinyint(3);default:1"`
}

func (ca *Category) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "catgegory")
}
