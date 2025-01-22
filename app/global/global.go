package global

import (
	"github.com/miaoming3/wallpaper/app/config"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
)

var (
	SysConfig *config.Config
	DbClient  *gorm.DB
	Captcha   *base64Captcha.Captcha
)
