package bsdb

import (
	"blockchain_services/config"
	"fmt"
	"log"
	"strconv"
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

func InitPgConn() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, username, password, database, port,
	)

	// 配置GORM，启用日志
	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	Db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 配置连接池
	sqlDB, err := Db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)                  // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour)        // 设置了连接可复用的最大时间
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 设置连接的最大空闲时间

	// 测试连接
	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}

	// 统一执行数据库迁移
	tables := []interface{}{
		&Block{},
		&Transaction{},
		&Receipt{},
		&Log{},
		&Account{},
		&Token{},
		&NFT{},
	}

	for _, table := range tables {
		if err = Db.AutoMigrate(table); err != nil {
			return fmt.Errorf("数据库迁移失败: %w", err)
		}
	}

	// 创建索引以优化查询性能
	if err = createIndexes(); err != nil {
		log.Printf("创建索引警告: %v", err)
		// 索引创建失败不影响主流程，只记录警告
	}

	log.Println("数据库初始化成功")
	return nil
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
		"CREATE INDEX IF NOT EXISTS idx_transactions_from_address ON transactions(from_address)",
		"CREATE INDEX IF NOT EXISTS idx_transactions_to_address ON transactions(to_address)",

		// 收据表索引
		"CREATE INDEX IF NOT EXISTS idx_receipts_transaction_hash ON receipts(transaction_hash)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_block_hash ON receipts(block_hash)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_from_address ON receipts(from_address)",
		"CREATE INDEX IF NOT EXISTS idx_receipts_to_address ON receipts(to_address)",

		// 日志表索引
		"CREATE INDEX IF NOT EXISTS idx_logs_transaction_hash ON logs(transaction_hash)",
		"CREATE INDEX IF NOT EXISTS idx_logs_block_number ON logs(block_number)",
		"CREATE INDEX IF NOT EXISTS idx_logs_address ON logs(address)",

		// 账户表索引（已在struct中定义uniqueIndex，这里添加其他索引）
		"CREATE INDEX IF NOT EXISTS idx_account_tx_count ON accounts(total_transaction)",
		"CREATE INDEX IF NOT EXISTS idx_account_last_tx ON accounts(last_transaction_at)",

		// Token表索引和唯一约束
		"CREATE UNIQUE INDEX IF NOT EXISTS idx_token_account_address_unique ON tokens(account_id, token_address)",
		"CREATE INDEX IF NOT EXISTS idx_token_address ON tokens(token_address)",

		// NFT表索引和唯一约束
		"CREATE UNIQUE INDEX IF NOT EXISTS idx_nft_account_contract_token_unique ON nfts(account_id, contract_address, token_id)",
		"CREATE INDEX IF NOT EXISTS idx_nft_contract_token ON nfts(contract_address, token_id)",
		"CREATE INDEX IF NOT EXISTS idx_nft_standard ON nfts(standard)",
	}

	for _, idxSQL := range indexes {
		if err := Db.Exec(idxSQL).Error; err != nil {
			return fmt.Errorf("创建索引失败 [%s]: %w", idxSQL, err)
		}
	}
	return nil
}

// GetBlockByNumber 1.区块详情
func GetBlockByNumber(number string) (Block, error) {
	var block Block
	err := Db.Where("number = ?", number).First(&block).Error
	if err != nil {
		return block, fmt.Errorf("查询区块失败: %w", err)
	}
	return block, nil
}

// GetTransactionByHash 2.交易详情
func GetTransactionByHash(hash string) (Transaction, error) {
	var transaction Transaction
	err := Db.Where("hash = ?", hash).First(&transaction).Error
	if err != nil {
		return transaction, fmt.Errorf("查询交易失败: %w", err)
	}
	return transaction, nil
}

// GetTransactions 3.批量交易查询
func GetTransactions(start, end int) ([]Transaction, error) {
	var transactions []Transaction
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 1000 {
		limit = 1000 // 限制最大查询数量
	}

	err := Db.
		Order("id ASC").
		Offset(start).
		Limit(limit).
		Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("查询交易列表失败: %w", err)
	}
	return transactions, nil
}

// GetBlocks 4.批量区块查询
func GetBlocks(start, end int) ([]Block, error) {
	var blocks []Block
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 1000 {
		limit = 1000 // 限制最大查询数量
	}

	err := Db.
		Order("id ASC").
		Offset(start).
		Limit(limit).
		Find(&blocks).Error
	if err != nil {
		return nil, fmt.Errorf("查询区块列表失败: %w", err)
	}
	return blocks, nil
}

// GetAddressCount 5.地址数量
func GetAddressCount() (int64, error) {
	var count int64
	err := Db.Raw("SELECT COUNT(DISTINCT address) FROM (SELECT from_address AS address FROM transactions WHERE from_address IS NOT NULL AND from_address != '' UNION SELECT to_address AS address FROM transactions WHERE to_address IS NOT NULL AND to_address != '') AS all_addresses").Scan(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询地址数量失败: %w", err)
	}
	return count, nil
}

// GetTransactionCount 6.交易总数
func GetTransactionCount() (uint64, error) {
	var count int64
	err := Db.Model(&Transaction{}).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询交易总数失败: %w", err)
	}
	return uint64(count), nil
}

// GetBlockCount 7.区块总数
func GetBlockCount() (string, error) {
	var count int64
	err := Db.Model(&Block{}).Count(&count).Error
	if err != nil {
		return "0", fmt.Errorf("查询区块总数失败: %w", err)
	}
	return strconv.FormatInt(count, 10), nil
}

// GetReceiptByHash 8.查询收据
func GetReceiptByHash(hash string) (Receipt, error) {
	var receipt Receipt
	err := Db.Where("transaction_hash = ?", hash).First(&receipt).Error
	if err != nil {
		return receipt, fmt.Errorf("查询收据失败: %w", err)
	}
	return receipt, nil
}

// GetReceipts 9.批量查询收据
func GetReceipts(start, end int) ([]Receipt, error) {
	var receipts []Receipt
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 1000 {
		limit = 1000 // 限制最大查询数量
	}

	err := Db.
		Order("id ASC").
		Offset(start).
		Limit(limit).
		Find(&receipts).Error
	if err != nil {
		return nil, fmt.Errorf("查询收据列表失败: %w", err)
	}
	return receipts, nil
}

// GetReceiptsByLast 10.查询最近7天收据
func GetReceiptsByLast() map[string]int {
	now := time.Now().Truncate(24 * time.Hour)
	sevenDaysAgo := now.AddDate(0, 0, -6) // 包含今天共7天：-6 到 0

	counts := make(map[string]int)
	for i := 0; i < 7; i++ {
		date := sevenDaysAgo.AddDate(0, 0, i)
		dateStr := date.Format("01-02")
		counts[dateStr] = 0
	}

	// 优化：使用SQL聚合查询替代加载所有区块数据
	type DateCount struct {
		Date  string
		Count int64
	}

	var results []DateCount
	err := Db.Raw(`
		SELECT 
			TO_CHAR(TO_TIMESTAMP(timestamp), 'MM-DD') AS date,
			COUNT(*) AS count
		FROM blocks
		WHERE timestamp >= ?
		GROUP BY TO_CHAR(TO_TIMESTAMP(timestamp), 'MM-DD')
		ORDER BY date
	`, sevenDaysAgo.Unix()).Scan(&results).Error

	if err != nil {
		log.Printf("GetReceiptsByLast查询失败: %v", err)
		// 降级到原来的方法
		var blocks []Block
		Db.Where("timestamp >= ?", sevenDaysAgo.Unix()).Find(&blocks)
		for _, block := range blocks {
			dateStr := time.Unix(int64(block.Timestamp), 0).Format("01-02")
			if _, exists := counts[dateStr]; exists && len(block.Transactions) > 0 {
				counts[dateStr] += len(block.Transactions)
			}
		}
	} else {
		// 使用SQL查询结果
		for _, result := range results {
			if _, exists := counts[result.Date]; exists {
				counts[result.Date] = int(result.Count)
			}
		}
	}

	return counts
}

// GetTransactionsByAddress 11.根据地址查询交易
func GetTransactionsByAddress(address string, start int, end int) ([]Receipt, error) {
	var receipts []Receipt
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 1000 {
		limit = 1000 // 限制最大查询数量
	}

	err := Db.Where("from_address = ? OR to_address = ?", address, address).
		Order("id DESC").
		Offset(start).
		Limit(limit).
		Find(&receipts).Error
	if err != nil {
		return nil, fmt.Errorf("根据地址查询交易失败: %w", err)
	}
	return receipts, nil
}

// GetReceiptsByAddress 12.根据地址查询收据
func GetReceiptsByAddress(address string) ([]Receipt, error) {
	var receipts []Receipt
	err := Db.Where("from_address = ? OR to_address = ?", address, address).
		Order("id DESC").
		Find(&receipts).Error
	if err != nil {
		return nil, fmt.Errorf("根据地址查询收据失败: %w", err)
	}
	return receipts, nil
}

// GetTransactionCountByAddress 13. 根据地址查询地址交易总数
func GetTransactionCountByAddress(address string) (int64, error) {
	var count int64
	err := Db.Model(&Receipt{}).
		Where("from_address = ? OR to_address = ?", address, address).
		Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询地址交易总数失败: %w", err)
	}
	return count, nil
}

// GetAddresses 获取所有唯一地址
func GetAddresses() ([]string, error) {
	var addresses []string
	// 使用UNION查询获取所有唯一的from_address和to_address
	err := Db.Raw(`
		SELECT DISTINCT address FROM (
			SELECT from_address AS address FROM transactions WHERE from_address IS NOT NULL AND from_address != ''
			UNION
			SELECT to_address AS address FROM transactions WHERE to_address IS NOT NULL AND to_address != ''
		) AS all_addresses
	`).Scan(&addresses).Error
	if err != nil {
		return nil, fmt.Errorf("查询唯一地址失败: %w", err)
	}
	return addresses, nil
}
