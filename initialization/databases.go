package initialization

import (
	"fmt"
	"github.com/miaoming3/wallpaper/global"
	"github.com/miaoming3/wallpaper/models"
	"gorm.io/driver/mysql"
	"log"
	"time"

	"gorm.io/gorm"
)

func InitDataBases() {
	baseConfig := global.SysConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		baseConfig.Username, baseConfig.Password, baseConfig.Host, baseConfig.Port, baseConfig.DbName, baseConfig.Charset)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("databases content err : %v", err)
	}
	sqlDB, _ := Db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.DbClient = Db
	autoMigrate()
}

func autoMigrate() {
	global.DbClient.AutoMigrate(&models.Admin{})
}
