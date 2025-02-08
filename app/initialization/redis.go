package initialization

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/miaoming3/wallpaper/app/global"
	"log"
	"time"
)

func InitRedis() {
	global.RedisClien = createRedisClient()
	if _, err := global.RedisClien.Ping(context.Background()).Result(); err != nil {
		log.Panic("检查redis 服务是否启动")
	}

	go autoReconnect(global.RedisClien)
}

func createRedisClient() *redis.Client {
	options := &redis.Options{
		Addr: fmt.Sprintf("%v:%v", global.SysConfig.RedisConfig.Host, global.SysConfig.RedisConfig.Port),
	}
	if global.SysConfig.RedisConfig.Username != "" && global.SysConfig.RedisConfig.Password != "" {
		options.Username = global.SysConfig.RedisConfig.Username
		options.Password = global.SysConfig.RedisConfig.Password
	}

	if global.SysConfig.RedisConfig.Db > 0 {
		options.DB = global.SysConfig.RedisConfig.Db
	}

	return redis.NewClient(options)
}

func checkConnection(client *redis.Client) bool {
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return false
	}
	return true
}
func autoReconnect(client *redis.Client) {
	ticker := time.NewTicker(10 * time.Second) // 每10秒检查一次连接状态
	defer ticker.Stop()
	for range ticker.C {
		if !checkConnection(client) { // 检查连接是否有效
			global.RedisClien = createRedisClient() // 重新创建客户端并设置连接池
		}
	}
}
