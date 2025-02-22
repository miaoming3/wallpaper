package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"gorm.io/gorm"
)

const (
	StatusDisabled uint8 = 0
	StatusActive   uint8 = 1
)

type AdminModel struct {
	*gorm.Model
	Username string `gorm:"username;unique:username;type:varchar(150);comment:用户名"`
	Password string `gorm:"password;type:varchar(255);comment:密码"`
	Status   uint8  `gorm:"status;type:tinyint(3);default:1;comment:状态"`
	Email    string `gorm:"email;unique;type:varchar(255);comment:邮件"`
	Phone    string `gorm:"phone;unique;type:char(12);comment:手机号"`
	Avatar   string `gorm:"avatar;type:varchar(255);comment:头像"`
}

func (model *AdminModel) TableName() string {
	return fmt.Sprintf("%v%v", global.SysConfig.DataBaseConfig.Prefix, "admin")
}

func (model *AdminModel) StatusString(status uint8) (statusTxt string) {
	switch status {
	case StatusDisabled:
		statusTxt = "禁用"
	case StatusActive:
		statusTxt = "正常"
	}
	return
}
func (model *AdminModel) StatusUint(status string) (statusTxt uint8) {
	switch status {
	case "0":
	case "禁用":
		statusTxt = StatusDisabled
	case "1":
	case "正常":
	default:
		statusTxt = StatusActive
	}
	return
}
