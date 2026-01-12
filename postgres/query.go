package bsdb

import (
	"blockchain_services/redis"
	"fmt"
	"log"
	"time"
)

///////////////////////////////////////////////////////////////////////////////BLOCK/////////////////////////////////////////////////////////////////////////

// GetBlockByNumber 区块详情（带缓存）
func GetBlockByNumber(number string) (Block, error) {
	var block Block
	cacheKey := fmt.Sprintf("block:number:%s", number)

	// 尝试从缓存获取
	err := redis.CacheGetOrSet(cacheKey, &block, 5*time.Minute, func() (interface{}, error) {
		var b Block
		err := Db.Where("number = ?", number).First(&b).Error
		if err != nil {
			return nil, fmt.Errorf("查询区块失败: %w", err)
		}
		return b, nil
	})

	if err != nil {
		return block, err
	}
	return block, nil
}

// GetBlocks 批量区块查询
func GetBlocks(start, end int) ([]Block, error) {
	var blocks []Block
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
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

// GetBlockCount 区块总数（带缓存）
func GetBlockCount() (int64, error) {
	var count int64
	cacheKey := "stats:block_count"

	// 尝试从缓存获取，缓存30秒
	err := redis.CacheGetOrSet(cacheKey, &count, 30*time.Second, func() (interface{}, error) {
		var c int64
		err := Db.Model(&Block{}).Count(&c).Error
		if err != nil {
			return nil, fmt.Errorf("查询区块总数失败: %w", err)
		}
		return c, nil
	})

	if err != nil {
		return 0, err
	}
	return count, nil
}

///////////////////////////////////////////////////////////////////////////////TRANSACTION/////////////////////////////////////////////////////////////////////////

// GetTransactionByHash 交易详情（带缓存）
func GetTransactionByHash(hash string) (Transaction, error) {
	var transaction Transaction
	cacheKey := fmt.Sprintf("tx:hash:%s", hash)

	// 尝试从缓存获取，交易数据相对稳定，缓存10分钟
	err := redis.CacheGetOrSet(cacheKey, &transaction, 10*time.Minute, func() (interface{}, error) {
		var tx Transaction
		err := Db.Where("hash = ?", hash).First(&tx).Error
		if err != nil {
			return nil, fmt.Errorf("查询交易失败: %w", err)
		}
		return tx, nil
	})

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

// GetTransactions 批量交易查询
func GetTransactions(start, end int) ([]Transaction, error) {
	var transactions []Transaction
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.
		Order("idx_transactions_transaction_index_block_number ASC").
		Offset(start).
		Limit(limit).
		Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("查询交易列表失败: %w", err)
	}
	return transactions, nil
}

// GetTransactionCount 交易总数（带缓存）
func GetTransactionCount() (int64, error) {
	var count int64
	cacheKey := "stats:transaction_count"

	// 尝试从缓存获取，缓存30秒
	err := redis.CacheGetOrSet(cacheKey, &count, 30*time.Second, func() (interface{}, error) {
		var c int64
		err := Db.Model(&Transaction{}).Count(&c).Error
		if err != nil {
			return nil, fmt.Errorf("查询交易总数失败: %w", err)
		}
		return c, nil
	})

	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetTransactionsByAddress 批量查询该地址的交易
func GetTransactionsByAddress(address string, start int, end int) ([]Transaction, error) {
	var transactions []Transaction
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.Where("from_address = ? OR to_address = ?", address, address).
		Order("idx_transactions_transaction_index_block_number DESC").
		Offset(start).
		Limit(limit).
		Find(&transactions).Error
	if err != nil {
		return nil, fmt.Errorf("批量查询该地址的交易失败: %w", err)
	}
	return transactions, nil
}

// GetTransactionCountByAddress 查询该地址的交易总数
func GetTransactionCountByAddress(address string) (int64, error) {
	var count int64
	err := Db.Model(&Transaction{}).
		Where("from_address = ? OR to_address = ?", address, address).
		Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询地址交易总数失败: %w", err)
	}
	return count, nil
}

///////////////////////////////////////////////////////////////////////////////RECEIPT/////////////////////////////////////////////////////////////////////////

// GetReceiptByHash 查询收据
func GetReceiptByHash(hash string) (Receipt, error) {
	var receipt Receipt
	err := Db.Where("transaction_hash = ?", hash).First(&receipt).Error
	if err != nil {
		return receipt, fmt.Errorf("查询收据失败: %w", err)
	}
	return receipt, nil
}

// GetReceipts 批量查询收据
func GetReceipts(start, end int) ([]Receipt, error) {
	var receipts []Receipt
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.
		Order("idx_receipts_transaction_index_block_number ASC").
		Offset(start).
		Limit(limit).
		Find(&receipts).Error
	if err != nil {
		return nil, fmt.Errorf("查询收据列表失败: %w", err)
	}
	return receipts, nil
}

// GetReceiptCount 交易总数
// GetReceiptCount 查询收据总数（带缓存）
func GetReceiptCount() (int64, error) {
	var count int64
	cacheKey := "stats:receipt_count"

	// 尝试从缓存获取，缓存30秒
	err := redis.CacheGetOrSet(cacheKey, &count, 30*time.Second, func() (interface{}, error) {
		var c int64
		err := Db.Model(&Receipt{}).Count(&c).Error
		if err != nil {
			return nil, fmt.Errorf("查询收据总数失败: %w", err)
		}
		return c, nil
	})

	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetReceiptsByAddress 批量查询该地址的收据
func GetReceiptsByAddress(address string, start int, end int) ([]Receipt, error) {
	var receipts []Receipt
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.Where("from_address = ? OR to_address = ?", address, address).
		Order("idx_receipts_transaction_index_block_number DESC").
		Offset(start).
		Limit(limit).
		Find(&receipts).Error
	if err != nil {
		return nil, fmt.Errorf("批量查询该地址的收据失败: %w", err)
	}
	return receipts, nil
}

// GetReceiptCountByAddress 查询该地址的收据总数
func GetReceiptCountByAddress(address string) (int64, error) {
	var count int64
	err := Db.Model(&Receipt{}).
		Where("from_address = ? OR to_address = ?", address, address).
		Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询地址交易总数失败: %w", err)
	}
	return count, nil
}

// GetReceiptsByLast 查询最近7天收据（带缓存）
func GetReceiptsByLast() map[string]int {
	// 使用当前日期作为缓存key的一部分，确保每天的数据独立缓存
	today := time.Now().Format("2006-01-02")
	cacheKey := fmt.Sprintf("stats:receipts_last7days:%s", today)

	var counts map[string]int

	// 尝试从缓存获取，缓存5分钟
	err := redis.CacheGetOrSet(cacheKey, &counts, 5*time.Minute, func() (interface{}, error) {
		now := time.Now().Truncate(24 * time.Hour)
		sevenDaysAgo := now.AddDate(0, 0, -6) // 包含今天共7天：-6 到 0

		// 初始化7天的计数为0
		result := make(map[string]int)
		for i := 0; i < 7; i++ {
			date := sevenDaysAgo.AddDate(0, 0, i)
			dateStr := date.Format("01-02")
			result[dateStr] = 0
		}

		// 从receipts表查询最近7天的收据统计
		type DateCount struct {
			Date  string
			Count int64
		}

		var results []DateCount
		err := Db.Raw(`
			SELECT 
				TO_CHAR(TO_TIMESTAMP(timestamp), 'MM-DD') AS date,
				COUNT(*) AS count
			FROM receipts
			WHERE timestamp >= ?
			GROUP BY TO_CHAR(TO_TIMESTAMP(timestamp), 'MM-DD')
			ORDER BY date
		`, sevenDaysAgo.Unix()).Scan(&results).Error

		if err != nil {
			log.Printf("GetReceiptsByLast查询失败: %v", err)
			return result, nil
		}

		// 填充查询结果
		for _, r := range results {
			if _, exists := result[r.Date]; exists {
				result[r.Date] = int(r.Count)
			}
		}

		return result, nil
	})

	if err != nil {
		log.Printf("GetReceiptsByLast缓存获取失败: %v", err)
		// 返回空结果
		return make(map[string]int)
	}

	return counts
}

//////////////////////////////////////////////////////////////////////////////Account/////////////////////////////////////////////////////////////////////////

// GetAccount 查询账户
// GetAccount 查询账户（带缓存）
func GetAccount(address string) (Account, error) {
	var account Account
	cacheKey := fmt.Sprintf("account:%s", address)

	// 尝试从缓存获取，账户信息可能变化，缓存2分钟
	err := redis.CacheGetOrSet(cacheKey, &account, 2*time.Minute, func() (interface{}, error) {
		var acc Account
		err := Db.Where("address = ?", address).First(&acc).Error
		if err != nil {
			return nil, fmt.Errorf("查询账户失败: %w", err)
		}
		return acc, nil
	})

	if err != nil {
		return account, err
	}
	return account, nil
}

// GetAccounts 批量查询账户
func GetAccounts(start, end int) ([]Account, error) {
	var addresses []Account
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.
		Order("id ASC").
		Offset(start).
		Limit(limit).
		Find(&addresses).Error
	if err != nil {
		return nil, fmt.Errorf("查询账户列表失败: %w", err)
	}
	return addresses, nil
}

// GetAccountCount 账户数量
func GetAccountCount() (int64, error) {
	var count int64
	err := Db.Model(&Account{}).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询收据总数失败: %w", err)
	}
	return count, nil
}

///////////////////////////////////////////////////////////////////////////////Token//////////////////////////////////////////////////////////////////////////

// GetToken 查询Token
func GetToken(address string) (Token, error) {
	var token Token
	err := Db.Where("address = ?", address).First(&token).Error
	if err != nil {
		return token, fmt.Errorf("查询Token失败: %w", err)
	}
	return token, nil
}

// GetTokens 批量查询Token
func GetTokens(start, end int) ([]Token, error) {
	var tokens []Token
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.
		Order("id ASC").
		Offset(start).
		Limit(limit).
		Find(&tokens).Error
	if err != nil {
		return nil, fmt.Errorf("查询token列表失败: %w", err)
	}
	return tokens, nil
}

// GetTokenCount Token数量
func GetTokenCount() (int64, error) {
	var count int64
	err := Db.Model(&Token{}).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询Token总数失败: %w", err)
	}
	return count, nil
}

////////////////////////////////////////////////////////////////////////////////NFT///////////////////////////////////////////////////////////////////////////

// GetNFT 查询NFT
func GetNFT(address string) (NFT, error) {
	var nft NFT
	err := Db.Where("address = ?", address).First(&nft).Error
	if err != nil {
		return nft, fmt.Errorf("查询Token失败: %w", err)
	}
	return nft, nil
}

// GetNFTs 批量查询NFT
func GetNFTs(start, end int) ([]NFT, error) {
	var tokens []NFT
	// 参数验证
	if start < 0 || end <= 0 || end < start {
		return nil, fmt.Errorf("无效的分页参数: start=%d, end=%d", start, end)
	}
	limit := end - start
	if limit > 100 {
		limit = 100 // 限制最大查询数量
	}

	err := Db.
		Order("id ASC").
		Offset(start).
		Limit(limit).
		Find(&tokens).Error
	if err != nil {
		return nil, fmt.Errorf("查询nft列表失败: %w", err)
	}
	return tokens, nil
}

// GetNFTCount NFT数量
func GetNFTCount() (int64, error) {
	var count int64
	err := Db.Model(&NFT{}).Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("查询NFT总数失败: %w", err)
	}
	return count, nil
}
