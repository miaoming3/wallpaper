package main

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"github.com/miaoming3/wallpaper/initialization"
)

func init() {
	initialization.InitConfig("")
	initialization.InitDataBases()

}

func main() {
	v := initialization.InitRoutes()
	_ = v.Run(fmt.Sprintf(":%v", global.SysConfig.SysPort))
}
