package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
)

type Admin struct {
}

func (admin *Admin) TableName() string {

	return fmt.Sprintf("%s%s", global.SysConfig.Prefix, "admin")
}
