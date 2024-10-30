package models

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
}
type ShowStatus uint

const (
	ShowStatusOn  ShowStatus = iota + 1
	ShowStatusOff ShowStatus = 2
)

func (m *Menu) TableName() string {
	return fmt.Sprintf("%smenu", global.SysConfig.Prefix)
}
