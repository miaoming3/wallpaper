package initialization

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/miaoming3/wallpaper/app/global"
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
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(&global.SysConfig); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}
	})
	if err := viper.Unmarshal(&global.SysConfig); err != nil {
		log.Panicf("config unmarshal  err: %v", err)
	}

}
