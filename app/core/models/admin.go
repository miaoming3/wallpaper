package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Password string
	Username string
	Status   uint
	Uid      string
	UserUID  uint
}

func (admin *Admin) TableName() string {

	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "admin")
}
