package config

import (
	"math/big"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	GrpcPort string
	HttpPort string

	LocalUrl string
	Node1Url string
	Node2Url string
	Node3Url string
	ChainID  *big.Int

	RedisAddr     string
	RedisPassword string
	RedisDB       int

	DbHost     string
	DbUsername string
	DbPassword string
	DbName     string
	DbPort     string
)

// Init 初始化配置，从 .env 文件读取
func Init() {
	// 加载 .env 文件（如果存在）
	if err := godotenv.Load(); err != nil {
		panic("未找到 .env 文件: " + err.Error())
	}

	// 读取配置，所有配置项必须从 .env 文件读取
	GrpcPort = getEnv("GRPC_PORT")
	HttpPort = getEnv("HTTP_PORT")

	LocalUrl = getEnv("LOCAL_URL")
	Node1Url = getEnv("NODE1_URL")
	Node2Url = getEnv("NODE2_URL")
	Node3Url = getEnv("NODE3_URL")
	chainIDStr := getEnv("CHAIN_ID")
	chainIDInt, err := strconv.ParseInt(chainIDStr, 10, 64)
	if err != nil {
		panic("CHAIN_ID 解析失败: " + err.Error())
	}
	ChainID = big.NewInt(chainIDInt)

	RedisAddr = getEnv("REDIS_ADDR")
	RedisPassword = getEnv("REDIS_PASSWORD")
	redisDBStr := getEnv("REDIS_DB")
	redisDBInt, err := strconv.Atoi(redisDBStr)
	if err != nil {
		panic("REDIS_DB 解析失败: " + err.Error())
	}
	RedisDB = redisDBInt

	DbHost = getEnv("DB_HOST")
	DbUsername = getEnv("DB_USERNAME")
	DbPassword = getEnv("DB_PASSWORD")
	DbName = getEnv("DB_NAME")
	DbPort = getEnv("DB_PORT")
}

// getEnv 获取环境变量，如果不存在则 panic
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("环境变量 " + key + " 未设置，请在 .env 文件中配置")
	}
	return value
}
