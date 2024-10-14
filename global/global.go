package global

import (
	"github.com/miaoming3/wallpaper/config"
	"gorm.io/gorm"
)

var (
	SysConfig *config.Config
	DbClient  *gorm.DB
)
