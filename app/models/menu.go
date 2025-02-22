package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

type MenuType uint8

const (
	MenuTypeInterface MenuType = iota
	MenuTypeUrl
)

type Menu struct {
	*gorm.Model
	Name    string `gorm:"name;type:varchar(255);unique;comment:菜单名"`
	Url     string `gorm:"url;type:varchar(255);comment:链接"`
	Icon    string `gorm:"icon;type:varchar(500);comment:图标"`
	Pid     uint   `gorm:"pid;default:0;comment:父级id"`
	Type    uint8  `gorm:"type;default:0;type:tinyint(2);comment:类型 0 接口 1 链接"`
	IsAdmin uint   `gorm:"is_admin;default:0;type:tinyint(2);comment: 后台 0 前台 1 后台"`
	Method  string `gorm:"method;type:varchar(255);default:0;comment:请求方式 0 get 1 post 2 put 3 delete 4 header "`
	Key     string `gorm:"key;type:varchar(255);unique;comment:标识"`
	Status  uint   `gorm:"status;type:tinyint(2);default:1;comment: 状态 1 正常 2 禁用"`
}

func (model *Menu) TableName() string {

	return fmt.Sprintf("%v%v", global.SysConfig.DataBaseConfig.Prefix, "menu")
}
