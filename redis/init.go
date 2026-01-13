package redis

import (
	"blockchain_services/config"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         config.RedisAddr,
		Password:     config.RedisPassword,
		DB:           config.RedisDB,
		DialTimeout:  5 * time.Second, // 连接超时时间
		ReadTimeout:  3 * time.Second, // 读取超时时间
		WriteTimeout: 3 * time.Second, // 写入超时时间
		PoolSize:     10,              // 连接池大小
		MinIdleConns: 5,               // 最小空闲连接数
	})

	return RedisClient
}
