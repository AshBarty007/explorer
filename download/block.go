package download

import (
	"blockchain_services/blockchain"
	"blockchain_services/config"
	db "blockchain_services/postgres"
	"blockchain_services/redis"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"gorm.io/gorm"
)

// Sync 实时同步区块信息
func Sync() {
	var start, end uint64

	// 尝试从Redis获取同步状态
	syncStateKey := "sync:current_block"
	cachedBlock, err := redis.CacheGetString(syncStateKey)
	if err == nil && cachedBlock != "" {
		// 从缓存恢复同步状态
		start, err = strconv.ParseUint(cachedBlock, 10, 64)
		if err == nil {
			log.Printf("从Redis恢复同步状态，起始区块: %d", start)
		}
	}

	// 如果Redis中没有状态，从数据库查询
	if start == 0 {
		count, err := db.GetBlockCount()
		if err != nil {
			log.Fatalf("获取区块数量失败: %v", err)
		}
		start = uint64(count)
	}
	end = start

	// 连接以太坊客户端
	etClient, err := blockchain.Dial(config.LocalUrl)
	if err != nil {
		log.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer etClient.Close()

	// 每5秒检查一次新区块
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 批量处理配置
	const (
		maxBlocksPerBatch   = 100 // 每次最多同步100个区块
		batchSize           = 50  // 批量插入数据库的批次大小
		stateUpdateInterval = 10  // 每10个区块更新一次Redis状态
	)

	// 批量数据收集
	var batchBlocks []*db.Block
	var batchTransactions []*db.Transaction
	var batchReceipts []*db.Receipt
	var batchLogs []*db.Log

	lastStateUpdate := start

	for range ticker.C {
		// 获取最新区块号
		blockNumber, err := etClient.BlockNumber(context.Background())
		if err != nil {
			log.Printf("查询最高区块失败: %v", err)
			continue
		}

		// 如果最新区块号大于当前同步位置，开始同步
		if blockNumber > end {
			end = blockNumber
		}

		// 限制每次同步的区块数量，避免一次性同步过多
		syncEnd := start + maxBlocksPerBatch
		if syncEnd > end {
			syncEnd = end
		}

		// 同步从start到syncEnd的区块
		for ; start < syncEnd; start++ {
			// 收集区块数据
			blocks, txs, receipts, logs, err := collectBlockData(etClient, start)
			if err != nil {
				log.Printf("收集区块 %d 数据失败: %v", start, err)
				// 如果连续失败，可能需要等待
				break
			}

			// 添加到批量列表
			if blocks != nil {
				batchBlocks = append(batchBlocks, blocks)
			}
			batchTransactions = append(batchTransactions, txs...)
			batchReceipts = append(batchReceipts, receipts...)
			batchLogs = append(batchLogs, logs...)

			// 达到批量大小时执行批量插入
			if len(batchBlocks) >= batchSize {
				if err := batchInsertData(batchBlocks, batchTransactions, batchReceipts, batchLogs); err != nil {
					log.Printf("批量插入失败: %v", err)
					// 清空批量列表，避免重复插入
					batchBlocks = batchBlocks[:0]
					batchTransactions = batchTransactions[:0]
					batchReceipts = batchReceipts[:0]
					batchLogs = batchLogs[:0]
					break
				}
				// 清空批量列表
				batchBlocks = batchBlocks[:0]
				batchTransactions = batchTransactions[:0]
				batchReceipts = batchReceipts[:0]
				batchLogs = batchLogs[:0]
			}

			// 每N个区块更新一次Redis状态，减少写入频率
			if start-lastStateUpdate >= stateUpdateInterval {
				updateSyncState(start)
				lastStateUpdate = start
			}

			// 添加小延迟，避免过快请求RPC节点
			time.Sleep(10 * time.Millisecond)
		}

		// 处理剩余的批量数据
		if len(batchBlocks) > 0 {
			if err := batchInsertData(batchBlocks, batchTransactions, batchReceipts, batchLogs); err != nil {
				log.Printf("批量插入剩余数据失败: %v", err)
			} else {
				// 清空批量列表
				batchBlocks = batchBlocks[:0]
				batchTransactions = batchTransactions[:0]
				batchReceipts = batchReceipts[:0]
				batchLogs = batchLogs[:0]
			}
		}

		// 更新同步状态（兜底：处理了少于10个区块时也要更新状态）
		if start > lastStateUpdate {
			updateSyncState(start - 1)
			lastStateUpdate = start - 1
		}

		log.Printf("同步进度: %d/%d, 最新区块: %d", start-1, end, blockNumber)
	}
}

// collectBlockData 收集单个区块的所有数据（不写入数据库）
func collectBlockData(etClient *blockchain.Client, blockNumber uint64) (*db.Block, []*db.Transaction, []*db.Receipt, []*db.Log, error) {
	ctx := context.Background()

	// 1. 获取区块和交易
	block, txs, err := etClient.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("获取区块失败: %w", err)
	}

	// 检查区块是否处于pending状态
	if block.Hash().String() == "" {
		return nil, nil, nil, nil, fmt.Errorf("区块处于pending状态")
	}

	// 2. 转换区块数据
	dbBlock := convertBlock(block, txs)

	// 3. 批量获取整个区块的所有receipts（一次性获取，避免大量RPC调用）
	var receipts []*types.Receipt
	if len(txs) > 0 {
		// 使用BlockReceipts一次性获取整个区块的所有receipts
		blockHash := block.Hash()
		receipts, err = etClient.BlockReceipts(ctx, rpc.BlockNumberOrHashWithHash(blockHash, false))
		if err != nil {
			// 如果BlockReceipts失败，回退到逐笔获取（兼容不支持eth_getBlockReceipts的节点）
			log.Printf("批量获取receipts失败，回退到逐笔获取: %v", err)
			receipts = nil
		}
	}

	// 4. 处理交易
	var dbTransactions []*db.Transaction
	var dbReceipts []*db.Receipt
	var dbLogs []*db.Log

	if len(txs) > 0 {
		for i, tx := range txs {
			// 转换交易
			dbTx := convertTransaction(block, tx, uint64(i))
			dbTransactions = append(dbTransactions, dbTx)

			// 获取receipt（优先使用批量获取的结果）
			var receipt *types.Receipt
			if receipts != nil && i < len(receipts) {
				// 使用批量获取的receipt
				receipt = receipts[i]
			} else {
				// 回退到逐笔获取（兼容或处理边界情况）
				receipt, err = etClient.TransactionReceipt(ctx, tx.Hash())
				if err != nil {
					log.Printf("获取receipt失败 %s: %v", tx.Hash().String(), err)
					continue
				}
			}

			// 转换receipt
			dbReceipt := convertReceipt(receipt, block, tx, uint64(i))
			dbReceipts = append(dbReceipts, dbReceipt)

			// 转换logs
			if len(receipt.Logs) > 0 {
				for logIndex, logEntry := range receipt.Logs {
					dbLog := convertLog(logEntry, block, tx, uint64(i), uint(logIndex))
					dbLogs = append(dbLogs, dbLog)
				}
			}

			// 处理合约创建和账户（这些需要立即处理，不能批量）
			if err := processTransaction(etClient, tx, receipt, block); err != nil {
				log.Printf("处理交易失败 %s: %v", tx.Hash().String(), err)
			}
		}
	}

	return dbBlock, dbTransactions, dbReceipts, dbLogs, nil
}

// batchInsertData 批量插入数据到数据库
func batchInsertData(blocks []*db.Block, transactions []*db.Transaction, receipts []*db.Receipt, logs []*db.Log) error {
	// 使用事务确保数据一致性
	return db.Db.Transaction(func(tx *gorm.DB) error {
		// 批量插入区块
		if len(blocks) > 0 {
			// 先检查哪些区块已存在，只插入不存在的
			var newBlocks []*db.Block
			for _, block := range blocks {
				var existing db.Block
				result := tx.Where("hash = ? OR number = ?", block.Hash, block.Number).First(&existing)
				if result.Error != nil {
					if errors.Is(result.Error, gorm.ErrRecordNotFound) {
						newBlocks = append(newBlocks, block)
					}
				}
			}
			if len(newBlocks) > 0 {
				if err := tx.CreateInBatches(newBlocks, 100).Error; err != nil {
					return fmt.Errorf("批量插入区块失败: %w", err)
				}
			}
		}

		// 批量插入交易
		if len(transactions) > 0 {
			var newTxs []*db.Transaction
			for _, txData := range transactions {
				var existing db.Transaction
				result := tx.Where("hash = ?", txData.Hash).First(&existing)
				if result.Error != nil {
					if errors.Is(result.Error, gorm.ErrRecordNotFound) {
						newTxs = append(newTxs, txData)
					}
				}
			}
			if len(newTxs) > 0 {
				if err := tx.CreateInBatches(newTxs, 100).Error; err != nil {
					return fmt.Errorf("批量插入交易失败: %w", err)
				}
			}
		}

		// 批量插入收据
		if len(receipts) > 0 {
			var newReceipts []*db.Receipt
			for _, receipt := range receipts {
				var existing db.Receipt
				result := tx.Where("transaction_hash = ?", receipt.TransactionHash).First(&existing)
				if result.Error != nil {
					if errors.Is(result.Error, gorm.ErrRecordNotFound) {
						newReceipts = append(newReceipts, receipt)
					}
				}
			}
			if len(newReceipts) > 0 {
				if err := tx.CreateInBatches(newReceipts, 100).Error; err != nil {
					return fmt.Errorf("批量插入收据失败: %w", err)
				}
			}
		}

		// 批量插入日志
		if len(logs) > 0 {
			var newLogs []*db.Log
			for _, logData := range logs {
				var existing db.Log
				result := tx.Where("transaction_hash = ? AND log_index = ?", logData.TransactionHash, logData.LogIndex).First(&existing)
				if result.Error != nil {
					if errors.Is(result.Error, gorm.ErrRecordNotFound) {
						newLogs = append(newLogs, logData)
					}
				}
			}
			if len(newLogs) > 0 {
				if err := tx.CreateInBatches(newLogs, 100).Error; err != nil {
					return fmt.Errorf("批量插入日志失败: %w", err)
				}
			}
		}

		return nil
	})
}

// processTransaction 处理交易，检测合约创建、账户更新等
func processTransaction(etClient *blockchain.Client, tx *blockchain.Transaction, receipt *types.Receipt, block *types.Block) error {
	blockTime := time.Unix(int64(block.Time()), 0)

	// 1. 检测合约创建
	if receipt.ContractAddress != (common.Address{}) {
		contractAddr := receipt.ContractAddress.String()

		// 检测合约类型（Token/NFT）
		contractType := detectContractType(etClient, contractAddr)

		switch contractType {
		case "ERC20":
			// 创建Token记录
			token := &db.Token{
				Address:            contractAddr,
				Standard:           contractType,
				Creator:            tx.From().String(),
				CreatedTime:        blockTime,
				CreatedHash:        tx.Hash().String(),
				CreatedBlockNumber: block.Number().Uint64(),
			}

			// 尝试获取Token信息
			if err := fetchTokenInfo(etClient, token); err != nil {
				log.Printf("获取Token信息失败 %s: %v", contractAddr, err)
			}

			if err := db.WriteToken(*token); err != nil {
				log.Printf("写入Token失败 %s: %v", contractAddr, err)
			}
		case "ERC721":
			// 创建NFT记录
			nft := &db.NFT{
				Address:            contractAddr,
				Standard:           contractType,
				Creator:            tx.From().String(),
				CreatedTime:        blockTime,
				CreatedHash:        tx.Hash().String(),
				CreatedBlockNumber: block.Number().Uint64(),
			}

			// 尝试获取NFT信息
			if err := fetchNFTInfo(etClient, nft); err != nil {
				log.Printf("获取NFT信息失败 %s: %v", contractAddr, err)
			}

			if err := db.WriteNFT(*nft); err != nil {
				log.Printf("写入NFT失败 %s: %v", contractAddr, err)
			}
		default:
			// 创建合约记录
			contract := &db.Contract{
				Address:            tx.To().String(),
				Creator:            tx.From().String(),
				CreatedTime:        blockTime,
				CreatedHash:        tx.Hash().String(),
				CreatedBlockNumber: block.Number().Uint64(),
			}

			if err := db.WriteContract(*contract); err != nil {
				log.Printf("写入合约失败 %s: %v", tx.To().String(), err)
			}
		}
	}

	// 2. 处理账户（from和to地址）
	var fromAddr, toAddr string
	if tx.From() != nil {
		fromAddr = tx.From().String()
	}
	if tx.To() != nil {
		toAddr = tx.To().String()
	}

	// 更新from地址账户（发送方，主动发起交易）
	// 注意：在以太坊中，每笔有效交易都应该有发送者地址
	// 但如果 From() 返回 nil，说明交易数据可能异常，跳过处理
	if fromAddr != "" {
		if err := updateAccount(etClient, fromAddr, blockTime, true); err != nil {
			log.Printf("更新账户失败 %s: %v", fromAddr, err)
		}
	}

	// 更新to地址账户（接收方，被动接收）
	if toAddr != "" {
		if err := updateAccount(etClient, toAddr, blockTime, false); err != nil {
			log.Printf("更新账户失败 %s: %v", toAddr, err)
		}
	}

	return nil
}

// convertBlock 转换区块数据
func convertBlock(block *types.Block, txs []*blockchain.Transaction) *db.Block {
	var transactions db.StringSlice
	for _, tx := range txs {
		transactions = append(transactions, tx.Hash().String())
	}

	return &db.Block{
		ParentHash:       block.ParentHash().String(),
		Miner:            block.Coinbase().String(),
		StateRoot:        block.Root().String(),
		TransactionsRoot: block.TxHash().String(),
		ReceiptsRoot:     block.ReceiptHash().String(),
		LogsBloom:        hex.EncodeToString(block.Bloom().Bytes()),
		Difficulty:       block.Difficulty().String(),
		Number:           block.Number().String(),
		GasLimit:         block.GasLimit(),
		GasUsed:          block.GasUsed(),
		Timestamp:        block.Time(),
		ExtraData:        hex.EncodeToString(block.Extra()),
		BaseFeePerGas:    block.BaseFee().String(),
		Transactions:     transactions,
		Hash:             block.Hash().String(),
		Size:             uint64(block.Size()),
	}
}

// convertTransaction 转换交易数据
func convertTransaction(block *types.Block, tx *blockchain.Transaction, index uint64) *db.Transaction {
	R, S := tx.RawSignatureValues()
	var from, to string
	if tx.From() != nil {
		from = tx.From().String()
	}
	if tx.To() != nil {
		to = tx.To().String()
	}

	return &db.Transaction{
		Hash:             tx.Hash().String(),
		BlockNumber:      block.Number().String(),
		BlockHash:        block.Hash().String(),
		FromAddress:      from,
		ToAddress:        to,
		Value:            tx.Value().String(),
		GasPrice:         tx.GasPrice().String(),
		GasLimit:         tx.Gas(),
		InputData:        hex.EncodeToString(tx.Data()),
		Nonce:            tx.Nonce(),
		TransactionIndex: index,
		Type:             uint8(tx.Type()),
		ChainId:          tx.ChainId().Uint64(),
		R:                R.String(),
		S:                S.String(),
		Timestamp:        block.Time(),
	}
}

// convertReceipt 转换收据数据
func convertReceipt(receipt *types.Receipt, block *types.Block, tx *blockchain.Transaction, index uint64) *db.Receipt {
	var fromAddr, toAddr string
	if tx.From() != nil {
		fromAddr = tx.From().String()
	}
	if tx.To() != nil {
		toAddr = tx.To().String()
	}

	// PostState是[]byte类型，需要转换为hex字符串
	postState := ""
	if len(receipt.PostState) > 0 {
		postState = hex.EncodeToString(receipt.PostState)
	}

	return &db.Receipt{
		Type:              uint8(receipt.Type),
		Root:              postState,
		Status:            receipt.Status,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		LogsBloom:         hex.EncodeToString(receipt.Bloom.Bytes()),
		Logs:              len(receipt.Logs),
		TransactionHash:   tx.Hash().String(),
		ContractAddress:   receipt.ContractAddress.String(),
		GasUsed:           receipt.GasUsed,
		BlockHash:         block.Hash().String(),
		BlockNumber:       block.Number().String(),
		TransactionIndex:  uint(index),
		FromAddress:       fromAddr,
		ToAddress:         toAddr,
		Timestamp:         block.Time(),
	}
}

// convertLog 转换日志数据
func convertLog(logEntry *types.Log, block *types.Block, tx *blockchain.Transaction, txIndex uint64, logIndex uint) *db.Log {
	var topics db.StringSlice
	for _, topic := range logEntry.Topics {
		topics = append(topics, topic.Hex())
	}

	return &db.Log{
		Address:          logEntry.Address.String(),
		Topics:           topics,
		Data:             hex.EncodeToString(logEntry.Data),
		BlockNumber:      block.Number().Uint64(),
		TransactionHash:  tx.Hash().String(),
		TransactionIndex: uint(txIndex),
		BlockHash:        block.Hash().String(),
		LogIndex:         logIndex,
		Removed:          logEntry.Removed,
	}
}

// updateSyncState 更新同步状态到Redis
func updateSyncState(blockNumber uint64) {
	syncStateKey := "sync:current_block"
	blockStr := strconv.FormatUint(blockNumber, 10)
	err := redis.CacheSetString(syncStateKey, blockStr, 24*time.Hour)
	if err != nil {
		log.Printf("更新同步状态失败: %v", err)
	}
}
