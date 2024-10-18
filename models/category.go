package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string
	Pid    uint
	Sort   uint
	Status uint
}

func (ca *Category) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "catgegory")
}
