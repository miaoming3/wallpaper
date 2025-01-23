package initialization

import (
	"fmt"
	"github.com/miaoming3/wallpaper/app/global"
	"github.com/miaoming3/wallpaper/app/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func RetryableOpenDatabase(dsn string, config *gorm.Config, retries int, retryInterval time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	for i := 0; i < retries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), config)
		if err == nil {
			return db, nil // 成功连接，返回数据库实例和nil错误
		}
		time.Sleep(retryInterval)                   // 等待指定的重试间隔
		log.Printf("尝试第 %d 次连接数据库失败: %v", i+1, err) // 使用Printf记录重试尝试（可以根据需要调整日志级别）
	}
	return nil, err // 重试失败，返回nil数据库实例和最后的错误
}
func InitDataBases() {
	baseConfig := global.SysConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		baseConfig.Username, baseConfig.Password, baseConfig.Host, baseConfig.Port, baseConfig.DbName, baseConfig.Charset)
	Db, err := RetryableOpenDatabase(dsn, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,       // Don't include params in the SQL log
				Colorful:                  false,       // Disable color
			})}, 3, 2*time.Second) // 尝试3次，每次间隔2秒
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err) // 在无法连接时，使用Fatalf记录致命错误并退出程序
	}
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
	_ = global.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.AdminModel{},
	)
	pwdbyte, _ := bcrypt.GenerateFromPassword([]byte("dx067870"), bcrypt.DefaultCost)
	global.DbClient.Model(&models.AdminModel{}).Where("username =? and status= ?", "admin", 1).FirstOrCreate(&models.AdminModel{
		Username: "admin",
		Password: string(pwdbyte),
	})
}
