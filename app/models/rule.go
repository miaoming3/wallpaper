package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type Rule struct {
	*gorm.Model
}

func (model *Rule) TableName() string {

	return fmt.Sprintf("%v%v", global.SysConfig.DataBaseConfig.Prefix, "rule")
}
