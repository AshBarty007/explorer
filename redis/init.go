package redis

import (
	"blockchain_services/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	return RedisClient
}
