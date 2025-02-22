package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model
}

func (model *Role) TableName() string {

	return fmt.Sprintf("%v%v", global.SysConfig.DataBaseConfig.Prefix, "role")
}
