package main

import (
	"blockchain_services/config"
	"blockchain_services/download"
	server "blockchain_services/http"
	db "blockchain_services/postgres"
	"blockchain_services/redis"
)

func main() {
	// 初始化配置（从 .env 文件读取）
	config.Init()

	// 初始化Redis客户端
	redis.NewRedisClient()

	// 初始化数据库连接
	db.InitPgConn()

	// 启动区块同步服务
	go download.Sync()

	// 启动HTTP服务
	server.StartHttp()
}
