package bsdb

import (
	"blockchain_services/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host     = config.DbHost
	username = config.DbUsername
	password = config.DbPassword
	database = config.DbName
	port     = config.DbPort
	Db       *gorm.DB
)

func InitPgConn() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, database, port,
	)

	// 配置GORM，启用日志
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	Db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 配置连接池
	sqlDB, err := Db.DB()
	if err != nil {
		panic("获取数据库连接失败: " + err.Error())
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)                  // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour)        // 设置了连接可复用的最大时间
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 设置连接的最大空闲时间

	// 测试连接
	if err = sqlDB.Ping(); err != nil {
		panic("数据库连接测试失败: " + err.Error())
	}

	// 统一执行数据库迁移
	tables := []interface{}{
		&Block{},
		&Transaction{},
		&Receipt{},
		&Log{},
		&Token{},
		&NFT{},
		&Account{},
		&Contract{},
	}

	for _, table := range tables {
		if err = Db.AutoMigrate(table); err != nil {
			panic("数据库迁移失败: " + err.Error())
		}
	}

	// 创建索引以优化查询性能
	if err = createIndexes(); err != nil {
		log.Printf("创建索引警告: %v", err)
		// 索引创建失败不影响主流程，只记录警告
	}

	log.Println("数据库初始化成功")
}

// createIndexes 创建数据库索引以优化查询性能
func createIndexes() error {
	indexes := []string{
		// 区块表索引
		"CREATE INDEX IF NOT EXISTS idx_blocks_number ON blocks(number)",
		"CREATE INDEX IF NOT EXISTS idx_blocks_hash ON blocks(hash)",
		"CREATE INDEX IF NOT EXISTS idx_blocks_timestamp ON blocks(timestamp)",

		// 交易表索引
		"CREATE INDEX IF NOT EXISTS idx_transactions_hash ON transactions(hash)",
		"CREATE INDEX IF NOT EXISTS idx_transactions_block_number ON transactions(block_number)",
		"CREATE INDEX IF NOT EXISTS idx_transactions_transaction_index_block_number ON transactions(transaction_index,block_number)",
		"CREATE INDEX IF NOT EXISTS idx_transactions_from_address ON transactions(from_address)",
		"CREATE INDEX IF NOT EXISTS idx_transactions_to_address ON transactions(to_address)",

		// 收据表索引
		"CREATE INDEX IF NOT EXISTS idx_receipts_transaction_hash ON receipts(transaction_hash)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_block_hash ON receipts(block_hash)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_transaction_index_block_number ON receipts(transaction_index,block_number)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_from_address ON receipts(from_address)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_to_address ON receipts(to_address)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_timestamp ON receipts(timestamp)",

		// 日志表索引
		"CREATE INDEX IF NOT EXISTS idx_logs_transaction_hash ON logs(transaction_hash)",
		"CREATE INDEX IF NOT EXISTS idx_logs_block_number ON logs(block_number)",
		"CREATE INDEX IF NOT EXISTS idx_logs_address ON logs(address)",

		// 账户表索引（已在struct中定义uniqueIndex，这里添加其他索引）
		"CREATE INDEX IF NOT EXISTS idx_account_tx_count ON accounts(total_transaction)",
		"CREATE INDEX IF NOT EXISTS idx_account_last_tx ON accounts(last_transaction_at)",
		"CREATE INDEX IF NOT EXISTS idx_account_balance ON accounts(balance)",

		// Token表索引
		"CREATE INDEX IF NOT EXISTS idx_tokens_address ON tokens(address)",
		"CREATE INDEX IF NOT EXISTS idx_tokens_standard ON tokens(standard)",
		"CREATE INDEX IF NOT EXISTS idx_tokens_creator ON tokens(creator)",
		"CREATE INDEX IF NOT EXISTS idx_tokens_created_block_number ON tokens(created_block_number)",

		// NFT表索引
		"CREATE INDEX IF NOT EXISTS idx_nfts_address ON nfts(address)",
		"CREATE INDEX IF NOT EXISTS idx_nfts_standard ON nfts(standard)",
		"CREATE INDEX IF NOT EXISTS idx_nfts_creator ON nfts(creator)",
		"CREATE INDEX IF NOT EXISTS idx_nfts_created_block_number ON nfts(created_block_number)",

		// Contract表索引（uniqueIndex已在struct中定义，这里添加其他索引）
		"CREATE INDEX IF NOT EXISTS idx_contract_creator ON contracts(creator)",
		"CREATE INDEX IF NOT EXISTS idx_contract_created_time ON contracts(created_time)",
		"CREATE INDEX IF NOT EXISTS idx_contract_created_block ON contracts(created_block_number)",
	}

	for _, idxSQL := range indexes {
		if err := Db.Exec(idxSQL).Error; err != nil {
			return fmt.Errorf("创建索引失败 [%s]: %w", idxSQL, err)
		}
	}
	return nil
}
