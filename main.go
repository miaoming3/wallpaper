package main

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/initialization"
)

func init() {
	initialization.InitConfig("")
	initialization.InitCaptcha()
	initialization.InitDataBases()
	initialization.InitRedis()
	initialization.InitValidate()
}

func main() {

	v := initialization.InitRoute()
	initialization.InitLoggers()
	_ = v.Run(fmt.Sprintf(":%v", global.SysConfig.SysPort))
}
