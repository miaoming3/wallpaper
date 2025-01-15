package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;type:varchar(255)"` // 假设用户名是唯一的
	Uid      uint   `gorm:"unique:idx_uid;varchar(128)"`
}

func (model *Users) TableName() string {
	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "users")
}
