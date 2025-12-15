package bsdb

import (
	"blockchain_services/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
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
	gormCfg := &gorm.Config{}
	Db, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		return err
	}

	if !Db.Migrator().HasTable(&Block{}) {
		fmt.Println("表 'Block' 不存在，正在创建...")

		// 自动创建表（GORM Migrate）
		if err = Db.Table("blocks").AutoMigrate(&Block{}); err != nil {
			panic(fmt.Sprint("创建Block表失败:", err))
		}

		fmt.Println("表 'Block' 创建成功！")
	}

	// 检查Transaction表是否存在
	if !Db.Migrator().HasTable(&Transaction{}) {
		fmt.Println("表 'Transaction' 不存在，正在创建...")

		// 自动创建表（GORM Migrate）
		if err = Db.Table("transactions").AutoMigrate(&Transaction{}); err != nil {
			panic(fmt.Sprint("创建Transaction表失败:", err))
		}

		fmt.Println("表 'Transaction' 创建成功！")
	}

	// 检查Receipt表是否存在
	if !Db.Migrator().HasTable(&Receipt{}) {
		fmt.Println("表 'Receipt' 不存在，正在创建...")

		// 自动创建表（GORM Migrate）
		if err = Db.Table("receipts").AutoMigrate(&Receipt{}); err != nil {
			panic(fmt.Sprint("创建Receipt表失败:", err))
		}

		fmt.Println("表 'Receipt' 创建成功！")
	}

	// 检查Log表是否存在
	if !Db.Migrator().HasTable(&Log{}) {
		fmt.Println("表 'Log' 不存在，正在创建...")

		// 自动创建表（GORM Migrate）
		if err = Db.Table("logs").AutoMigrate(&Log{}); err != nil {
			panic(fmt.Sprint("创建Log表失败:", err))
		}

		fmt.Println("表 'Log' 创建成功！")
	}

	return nil
}

// GetBlockByNumber 1.区块详情
func GetBlockByNumber(number string) Block {
	var block Block
	Db.Where("number = ?", number).First(&block)
	return block
}

// GetTransactionByHash 2.交易详情
func GetTransactionByHash(hash string) Transaction {
	var transaction Transaction
	Db.Where("hash=?", hash).First(&transaction)
	return transaction
}

// GetTransactions 3.批量交易查询
func GetTransactions(start, end int) []Transaction {
	var transactions []Transaction
	Db.
		Order("id ASC").
		Offset(start).
		Limit(end).
		Find(&transactions)
	return transactions
}

// GetBlocks 4.批量区块查询
func GetBlocks(start, end int) []Block {
	var blocks []Block
	Db.
		Order("id ASC").
		Offset(start).
		Limit(end).
		Find(&blocks)
	return blocks
}

// GetAddressCount 5.地址数量
func GetAddressCount() int64 {
	var count int64
	err := Db.Raw("SELECT COUNT(DISTINCT from_address) FROM transactions").Scan(&count).Error
	if err != nil {
		log.Println("GetAddressCount 查询出错", err)
	}
	return count
}

// GetTransactionCount 6.交易总数
func GetTransactionCount() uint {
	var transaction Transaction
	result := Db.Last(&transaction)
	if result.Error != nil {
		log.Println("GetTransactionCount 查询出错", result.Error)
		return 0
	}
	return transaction.ID
}

// GetBlockCount 7.区块总数
func GetBlockCount() string {
	var block Block
	result := Db.Last(&block)
	if result.Error != nil {
		log.Println("GetBlockCount 查询出错", result.Error)
		return "0"
	}
	return strconv.Itoa(int(block.ID))
}

// GetReceiptByHash 8.查询收据
func GetReceiptByHash(hash string) Receipt {
	var receipt Receipt
	Db.Where("transaction_hash=?", hash).First(&receipt)
	return receipt
}

// GetReceipts 9.批量查询查询收据
func GetReceipts(start, end int) []Receipt {
	var receipts []Receipt
	Db.
		Order("id ASC").
		Offset(start).
		Limit(end).
		Find(&receipts)
	return receipts
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

	var block []Block
	Db.Where("timestamp >= ?", sevenDaysAgo.Unix()).Find(&block)
	for _, r := range block {
		dateStr := time.Unix(int64(r.Timestamp), 0).Format("01-02")
		if _, exists := counts[dateStr]; exists && len(r.Transactions) > 0 {
			counts[dateStr] += len(r.Transactions)
		}
	}

	return counts
}

// GetTransactionsByAddress 11.根据地址查询交易
func GetTransactionsByAddress(address string, start int, end int) []Receipt {
	var receipts []Receipt
	Db.Where("from_address = ?", address).
		Offset(start).
		Limit(end).
		Find(&receipts)
	return receipts
}

// GetReceiptsByAddress 12.根据地址查询收据
func GetReceiptsByAddress(address string) []Receipt {
	var receipts []Receipt
	Db.Where("from_address = ?", address).Find(&receipts)
	return receipts
}

// GetTransactionCountByAddress 13. 根据地址查询地址交易总数
func GetTransactionCountByAddress(address string) int64 {
	var count int64
	Db.Model(&Receipt{}).Where("from_address = ?", address).Count(&count)
	return count
}

// GetAddresses
func GetAddresses() []string {
	var addresses []string
	if err := Db.Model(&Transaction{}).Distinct("from_address").Pluck("from_address", &addresses); err != nil {
		log.Println("查询唯一FromAddress失败:", err)
	}
	return addresses
}
