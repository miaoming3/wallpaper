package main

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/initialization"
)

func init() {
	initialization.InitConfig("")
	initialization.InitDataBases()
	initialization.InitValidate()

}

func main() {

	v := initialization.InitRoute()
	_ = v.Run(fmt.Sprintf(":%v", global.SysConfig.SysPort))
}
