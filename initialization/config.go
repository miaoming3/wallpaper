package initialization

import (
	"github.com/miaoming3/wallpaper/global"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig(path string) {
	if path == "" {
		path = "./config.yaml"
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("not found config file err :%v", err)
		}
		log.Panicf("access  file err: %v ", err)
	}
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("reading config file err: %v", err)
	}
	if err := viper.Unmarshal(&global.SysConfig); err != nil {
		log.Panicf("config unmarshal  err: %v", err)
	}
	log.Println("=========================SysConfig========================")
	log.Println(global.SysConfig)

}
