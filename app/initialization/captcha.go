package initialization

import (
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type configJsonBody struct {
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func toRGBA(bgcolor []uint8) *color.RGBA {
	if len(bgcolor) == 0 {
		return nil
	}
	var colorRgba *color.RGBA
	for k, v := range bgcolor {
		switch k {
		case 0:
			colorRgba.R = v
		case 1:
			colorRgba.G = v
		case 2:
			colorRgba.B = v
		case 3:
			colorRgba.A = v
		}
	}
	return colorRgba
}

func InitCaptcha() {

	if global.SysConfig.Source == "" {
		if global.SysConfig.CaptchaType == "string" || global.SysConfig.CaptchaType == "chinese" {
			global.SysConfig.Source = "abcdefghijklmnopqrstvuwxyz1234567890"
		}
		if global.SysConfig.CaptchaType == "math" {
			global.SysConfig.Source = "1234567890"
		}
	}
	var config = configJsonBody{
		DriverAudio:   base64Captcha.DefaultDriverAudio,
		DriverString:  base64Captcha.NewDriverString(global.SysConfig.Height, global.SysConfig.Width, global.SysConfig.NoiseCount, global.SysConfig.ShowLineOptions, global.SysConfig.Length, global.SysConfig.Source, toRGBA(global.SysConfig.BgColor), nil, global.SysConfig.Fonts),
		DriverChinese: base64Captcha.NewDriverChinese(global.SysConfig.Height, global.SysConfig.Width, global.SysConfig.NoiseCount, global.SysConfig.ShowLineOptions, global.SysConfig.Length, global.SysConfig.Source, toRGBA(global.SysConfig.BgColor), nil, global.SysConfig.Fonts),
		DriverMath:    base64Captcha.NewDriverMath(global.SysConfig.Height, global.SysConfig.Width, global.SysConfig.NoiseCount, global.SysConfig.ShowLineOptions, toRGBA(global.SysConfig.BgColor), nil, global.SysConfig.Fonts),
		DriverDigit:   base64Captcha.DefaultDriverDigit,
	}
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch global.SysConfig.CaptchaType {
	case "audio":
		driver = config.DriverAudio
	case "string":
		driver = config.DriverString.ConvertFonts()
	case "math":
		driver = config.DriverMath.ConvertFonts()
	case "chinese":
		driver = config.DriverChinese.ConvertFonts()
	default:
		driver = config.DriverDigit
	}
	global.Captcha = base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

}
