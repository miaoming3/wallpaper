package main

import (
	"fmt"
	"github.com/miaoming3/wallpaper/http/global"
	initialization2 "github.com/miaoming3/wallpaper/http/initialization"

	"github.com/miaoming3/wallpaper/initialization"
)

func init() {
	initialization2.InitConfig("")
	initialization2.InitDataBases()

}

func main() {

	v := initialization.InitRoutes()
	_ = v.Run(fmt.Sprintf(":%v", global.SysConfig.SysPort))
}
