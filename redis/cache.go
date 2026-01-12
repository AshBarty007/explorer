package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var ctx = context.Background()

// CacheGet 从缓存获取数据
func CacheGet(key string, dest interface{}) error {
	if RedisClient == nil {
		return fmt.Errorf("Redis客户端未初始化")
	}

	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), dest)
}

// CacheSet 设置缓存数据
func CacheSet(key string, value interface{}, expiration time.Duration) error {
	if RedisClient == nil {
		return fmt.Errorf("Redis客户端未初始化")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("序列化失败: %w", err)
	}

	return RedisClient.Set(ctx, key, data, expiration).Err()
}

// CacheDelete 删除缓存
func CacheDelete(key string) error {
	if RedisClient == nil {
		return fmt.Errorf("Redis客户端未初始化")
	}
	return RedisClient.Del(ctx, key).Err()
}

// CacheExists 检查缓存是否存在
func CacheExists(key string) (bool, error) {
	if RedisClient == nil {
		return false, fmt.Errorf("Redis客户端未初始化")
	}
	count, err := RedisClient.Exists(ctx, key).Result()
	return count > 0, err
}

// CacheGetOrSet 获取缓存，如果不存在则执行函数并设置缓存
func CacheGetOrSet(key string, dest interface{}, expiration time.Duration, fn func() (interface{}, error)) error {
	// 尝试从缓存获取
	err := CacheGet(key, dest)
	if err == nil {
		return nil
	}

	// 缓存不存在，执行函数获取数据
	data, err := fn()
	if err != nil {
		return err
	}

	// 设置缓存
	if err := CacheSet(key, data, expiration); err != nil {
		// 缓存设置失败不影响返回数据
		fmt.Printf("警告: 设置缓存失败 %s: %v\n", key, err)
	}

	// 将数据赋值给dest
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataBytes, dest)
}

// CacheIncr 递增计数器
func CacheIncr(key string) (int64, error) {
	if RedisClient == nil {
		return 0, fmt.Errorf("Redis客户端未初始化")
	}
	return RedisClient.Incr(ctx, key).Result()
}

// CacheSetString 设置字符串缓存
func CacheSetString(key string, value string, expiration time.Duration) error {
	if RedisClient == nil {
		return fmt.Errorf("Redis客户端未初始化")
	}
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// CacheGetString 获取字符串缓存
func CacheGetString(key string) (string, error) {
	if RedisClient == nil {
		return "", fmt.Errorf("Redis客户端未初始化")
	}
	return RedisClient.Get(ctx, key).Result()
}
