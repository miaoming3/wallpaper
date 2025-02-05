package global

import (
	"github.com/miaoming3/wallpaper/app/config"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SysConfig   *config.Config
	DbClient    *gorm.DB
	Captcha     *base64Captcha.Captcha
	SugarLogger *zap.SugaredLogger
)
