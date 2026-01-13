package tests

import (
	"blockchain_services/config"
	bsdb "blockchain_services/postgres"
	redisCache "blockchain_services/redis"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDatabaseConnection 测试数据库连接
func TestDatabaseConnection(t *testing.T) {
	// 初始化配置
	config.Init()

	// 初始化数据库连接
	bsdb.InitPgConn()

}

// TestDatabaseQuery 测试数据库查询功能
func TestDatabaseQuery(t *testing.T) {
	// 初始化配置
	config.Init()

	// 初始化数据库连接
	bsdb.InitPgConn()

	// 测试简单查询
	var result int64
	err := bsdb.Db.Raw("SELECT 1").Scan(&result).Error
	assert.NoError(t, err, "数据库查询失败")
	assert.Equal(t, int64(1), result, "查询结果应该为 1")

	// 测试获取当前时间
	var currentTime time.Time
	err = bsdb.Db.Raw("SELECT NOW()").Scan(&currentTime).Error
	assert.NoError(t, err, "获取当前时间失败")
	assert.False(t, currentTime.IsZero(), "当前时间不应该为零值")
	t.Logf("数据库当前时间: %s", currentTime.Format(time.RFC3339))

	// 测试获取数据库版本
	var version string
	err = bsdb.Db.Raw("SELECT version()").Scan(&version).Error
	assert.NoError(t, err, "获取数据库版本失败")
	assert.NotEmpty(t, version, "数据库版本不应该为空")
	t.Logf("数据库版本: %s", version)
}

// TestRedisConnection 测试 Redis 连接
func TestRedisConnection(t *testing.T) {
	// 初始化配置
	config.Init()

	// 创建 Redis 客户端
	client := redisCache.NewRedisClient()
	require.NotNil(t, client, "Redis 客户端不应该为 nil")

	// 测试 Ping，使用更长的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		// 提供详细的错误信息，帮助诊断问题
		t.Logf("Redis 连接失败详情:")
		t.Logf("  地址: %s", config.RedisAddr)
		t.Logf("  错误: %v", err)
		t.Logf("  可能的原因:")
		t.Logf("    1. Redis 服务器未运行")
		t.Logf("    2. 网络连接问题（防火墙、网络不通）")
		t.Logf("    3. Redis 服务器地址或端口配置错误")
		t.Logf("    4. Redis 密码配置错误")
		t.Logf("  建议: 检查 Redis 服务器状态和网络连接")
	}
	assert.NoError(t, err, "Redis Ping 失败，请检查 Redis 服务器是否运行以及网络连接是否正常")
	if err == nil {
		assert.Equal(t, "PONG", pong, "Redis Ping 应该返回 PONG")
		t.Logf("Redis 连接成功: %s", pong)
	}
}

// TestRedisBasicOperations 测试 Redis 基本操作
func TestRedisBasicOperations(t *testing.T) {
	// 初始化配置
	config.Init()

	// 创建 Redis 客户端
	client := redisCache.NewRedisClient()
	require.NotNil(t, client, "Redis 客户端不应该为 nil")

	ctx := context.Background()
	testKey := "test:connection:key"
	testValue := "test_value"

	// 清理测试数据
	defer func() {
		client.Del(ctx, testKey)
	}()

	// 测试 Set
	err := client.Set(ctx, testKey, testValue, 10*time.Second).Err()
	assert.NoError(t, err, "Redis Set 失败")

	// 测试 Get
	value, err := client.Get(ctx, testKey).Result()
	assert.NoError(t, err, "Redis Get 失败")
	assert.Equal(t, testValue, value, "获取的值应该与设置的值相同")

	// 测试 Exists
	exists, err := client.Exists(ctx, testKey).Result()
	assert.NoError(t, err, "Redis Exists 失败")
	assert.Equal(t, int64(1), exists, "键应该存在")

	// 测试 TTL
	ttl, err := client.TTL(ctx, testKey).Result()
	assert.NoError(t, err, "Redis TTL 失败")
	assert.Greater(t, ttl, time.Duration(0), "TTL 应该大于 0")
	t.Logf("键的 TTL: %v", ttl)

	// 测试 Delete
	err = client.Del(ctx, testKey).Err()
	assert.NoError(t, err, "Redis Delete 失败")

	// 验证删除成功
	exists, err = client.Exists(ctx, testKey).Result()
	assert.NoError(t, err, "Redis Exists 失败")
	assert.Equal(t, int64(0), exists, "键应该已被删除")
}

// TestRedisCacheFunctions 测试 Redis 缓存函数
func TestRedisCacheFunctions(t *testing.T) {
	// 初始化配置
	config.Init()

	// 创建 Redis 客户端
	client := redisCache.NewRedisClient()
	require.NotNil(t, client, "Redis 客户端不应该为 nil")

	testKey := "test:cache:key"
	testValue := "test_cache_value"

	// 清理测试数据
	defer func() {
		redisCache.CacheDelete(testKey)
	}()

	// 测试 CacheSetString
	err := redisCache.CacheSetString(testKey, testValue, 10*time.Second)
	assert.NoError(t, err, "CacheSetString 失败")

	// 测试 CacheGetString
	value, err := redisCache.CacheGetString(testKey)
	assert.NoError(t, err, "CacheGetString 失败")
	assert.Equal(t, testValue, value, "获取的值应该与设置的值相同")

	// 测试 CacheExists
	exists, err := redisCache.CacheExists(testKey)
	assert.NoError(t, err, "CacheExists 失败")
	assert.True(t, exists, "键应该存在")

	// 测试 CacheDelete
	err = redisCache.CacheDelete(testKey)
	assert.NoError(t, err, "CacheDelete 失败")

	// 验证删除成功
	exists, err = redisCache.CacheExists(testKey)
	assert.NoError(t, err, "CacheExists 失败")
	assert.False(t, exists, "键应该已被删除")
}

// TestRedisIncr 测试 Redis 递增操作
func TestRedisIncr(t *testing.T) {
	// 初始化配置
	config.Init()

	// 创建 Redis 客户端
	client := redisCache.NewRedisClient()
	require.NotNil(t, client, "Redis 客户端不应该为 nil")

	testKey := "test:incr:key"

	// 清理测试数据
	defer func() {
		redisCache.CacheDelete(testKey)
	}()

	// 测试 CacheIncr
	count, err := redisCache.CacheIncr(testKey)
	assert.NoError(t, err, "CacheIncr 失败")
	assert.Equal(t, int64(1), count, "第一次递增应该返回 1")

	count, err = redisCache.CacheIncr(testKey)
	assert.NoError(t, err, "CacheIncr 失败")
	assert.Equal(t, int64(2), count, "第二次递增应该返回 2")

	t.Logf("递增后的值: %d", count)
}

// TestDatabaseAndRedisIntegration 测试数据库和 Redis 集成
func TestDatabaseAndRedisIntegration(t *testing.T) {
	// 初始化配置
	config.Init()

	// 初始化数据库
	bsdb.InitPgConn()
	sqlDB, err := bsdb.Db.DB()
	require.NoError(t, err, "获取数据库连接失败")
	defer sqlDB.Close()

	// 初始化 Redis
	client := redisCache.NewRedisClient()
	require.NotNil(t, client, "Redis 客户端不应该为 nil")

	// 测试数据库连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = sqlDB.PingContext(ctx)
	assert.NoError(t, err, "数据库连接失败")

	// 测试 Redis 连接
	pong, err := client.Ping(ctx).Result()
	assert.NoError(t, err, "Redis 连接失败")
	assert.Equal(t, "PONG", pong, "Redis Ping 应该返回 PONG")

	t.Logf("数据库和 Redis 连接都正常")
}
