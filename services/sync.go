package services

import (
	bs_eth "blockchain_services/ethclient"
	bs_db "blockchain_services/postgres"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/gmsm"
	"gorm.io/gorm"
)

func Sync(url string) {
	var start, end uint64

	// 初始化时查询一次区块数量
	count, err := bs_db.GetBlockCount()
	if err != nil {
		log.Fatalf("获取区块数量失败: %v", err)
	}
	start, err = strconv.ParseUint(count, 10, 64)
	if err != nil {
		log.Fatalf("解析区块数量失败: %v", err)
	}
	end = start

	etClient, err := bs_eth.Dial(url)
	if err != nil {
		log.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer etClient.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 批量处理配置
	const batchSize = 10 // 每批处理的区块数
	var batchBlocks []*Block
	var batchTransactions []*Transaction
	var batchReceipts []*Receipt
	var batchLogs []*Log

	for t := range ticker.C {
		blockNumber, err := etClient.BlockNumber(context.Background())
		if err != nil {
			log.Printf("查询最高区块失败: %v", err)
			continue
		}

		if blockNumber > end+100 {
			end += 100
		} else {
			end = blockNumber
		}

		// 批量下载区块
		for ; start <= end-1; start++ {
			block, txs, receipts, logs, err := DownloadBlockData(etClient, big.NewInt(int64(start)))
			if err != nil {
				log.Printf("区块号 %d 下载失败: %v", start, err)
				// 如果连续失败多次，可能需要等待
				break
			}

			// 添加到批量列表
			if block != nil {
				batchBlocks = append(batchBlocks, block)
			}
			batchTransactions = append(batchTransactions, txs...)
			batchReceipts = append(batchReceipts, receipts...)
			batchLogs = append(batchLogs, logs...)

			// 达到批量大小时执行批量插入
			if len(batchBlocks) >= batchSize {
				if err := BatchInsert(batchBlocks, batchTransactions, batchReceipts, batchLogs); err != nil {
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
		}

		// 处理剩余的批量数据
		if len(batchBlocks) > 0 {
			if err := BatchInsert(batchBlocks, batchTransactions, batchReceipts, batchLogs); err != nil {
				log.Printf("批量插入剩余数据失败: %v", err)
			} else {
				// 清空批量列表
				batchBlocks = batchBlocks[:0]
				batchTransactions = batchTransactions[:0]
				batchReceipts = batchReceipts[:0]
				batchLogs = batchLogs[:0]
			}
		}

		log.Printf("同步进度: %d, 最新区块: %d, 当前时间: %s", start, blockNumber, t.Format("2006-01-02 15:04:05"))
		end = start
	}
}

// DownloadBlockData 下载区块数据并转换为本地格式，不进行数据库插入
func DownloadBlockData(etClient *bs_eth.Client, number *big.Int) (*Block, []*Transaction, []*Receipt, []*Log, error) {
	block, txs, err := etClient.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("下载区块数据失败: %w", err)
	}

	if block.Hash().String() == "" {
		return nil, nil, nil, nil, fmt.Errorf("区块处于pending状态: %s", number.String())
	}

	// 1. 转换区块数据
	b := toLocalBlock(block, txs)

	// 2. 转换交易数据
	var transactions []*Transaction
	var receipts []*Receipt
	var logs []*Log

	if len(txs) > 0 {
		// 先转换所有交易数据
		transactions = make([]*Transaction, len(txs))
		for i, tx := range txs {
			transactions[i] = toLocalTransaction(block, tx, uint64(i))
		}

		// 使用goroutine并发获取收据
		type receiptResult struct {
			index    int
			receipt  *types.Receipt
			err      error
			fromAddr string
			toAddr   string
			txHash   common.Hash
		}

		receiptChan := make(chan receiptResult, len(txs))
		var wg sync.WaitGroup

		// 限制并发数量，避免过多goroutine导致资源耗尽
		const maxConcurrency = 10
		semaphore := make(chan struct{}, maxConcurrency)

		// 启动goroutine并发获取收据
		for i, tx := range txs {
			wg.Add(1)
			semaphore <- struct{}{} // 获取信号量

			// 在goroutine外部获取地址，避免竞态条件
			transaction := transactions[i]
			fromAddr := transaction.FromAddress
			toAddr := transaction.ToAddress
			txHash := tx.Hash()

			go func(idx int, hash common.Hash, from, to string) {
				defer wg.Done()
				defer func() { <-semaphore }() // 释放信号量

				receipt, err := etClient.TransactionReceipt(context.Background(), hash)
				receiptChan <- receiptResult{
					index:    idx,
					receipt:  receipt,
					err:      err,
					fromAddr: from,
					toAddr:   to,
					txHash:   hash,
				}
			}(i, txHash, fromAddr, toAddr)
		}

		// 等待所有goroutine完成
		go func() {
			wg.Wait()
			close(receiptChan)
		}()

		// 收集结果，使用map保持索引映射以便后续排序
		receiptMap := make(map[int]receiptResult)
		for result := range receiptChan {
			if result.err != nil {
				log.Printf("获取交易收据失败 [索引:%d, 哈希:%s]: %v", result.index, result.txHash.String(), result.err)
				continue
			}
			if result.receipt == nil {
				continue
			}
			receiptMap[result.index] = result
		}

		// 按索引顺序处理结果，保持交易顺序
		for i := 0; i < len(txs); i++ {
			result, ok := receiptMap[i]
			if !ok {
				continue
			}

			r := toLocalReceipt(
				result.receipt,
				bs_eth.HexToAddress(result.fromAddr),
				bs_eth.HexToAddress(result.toAddr),
			)
			receipts = append(receipts, r)

			// 转换日志
			for _, logEntry := range result.receipt.Logs {
				logData := toLocalLog(logEntry)
				logs = append(logs, logData)
			}
		}
	}

	return b, transactions, receipts, logs, nil
}

// BatchGetReceipts 批量获取交易收据（如果RPC支持）
func BatchGetReceipts(etClient *bs_eth.Client, txs []*bs_eth.Transaction) (map[common.Hash]*types.Receipt, error) {
	receiptMap := make(map[common.Hash]*types.Receipt)

	// 尝试使用BlockReceipts（如果可用）
	// 否则逐个获取
	for _, tx := range txs {
		receipt, err := etClient.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return nil, fmt.Errorf("获取收据失败 %s: %w", tx.Hash().String(), err)
		}
		receiptMap[tx.Hash()] = receipt
	}

	return receiptMap, nil
}

// BatchInsert 批量插入数据，使用事务确保一致性
func BatchInsert(blocks []*Block, transactions []*Transaction, receipts []*Receipt, logs []*Log) error {
	// 使用事务确保数据一致性
	return bs_db.Db.Transaction(func(tx *gorm.DB) error {
		// 批量插入区块
		if len(blocks) > 0 {
			// 转换为数据库模型
			dbBlocks := make([]*bs_db.Block, len(blocks))
			for i, b := range blocks {
				dbBlocks[i] = &bs_db.Block{
					ParentHash:       b.ParentHash,
					Miner:            b.Miner,
					StateRoot:        b.StateRoot,
					TransactionsRoot: b.TransactionsRoot,
					ReceiptsRoot:     b.ReceiptsRoot,
					LogsBloom:        b.LogsBloom,
					Difficulty:       b.Difficulty,
					Number:           b.Number,
					GasLimit:         b.GasLimit,
					GasUsed:          b.GasUsed,
					Timestamp:        b.Timestamp,
					ExtraData:        b.ExtraData,
					BaseFeePerGas:    b.BaseFeePerGas,
					Transactions:     *b.Transactions,
					Hash:             b.Hash,
					Size:             b.Size,
				}
			}
			if err := tx.CreateInBatches(dbBlocks, 100).Error; err != nil {
				return fmt.Errorf("批量插入区块失败: %w", err)
			}
		}

		// 批量插入交易
		if len(transactions) > 0 {
			dbTransactions := make([]*bs_db.Transaction, len(transactions))
			for i, t := range transactions {
				dbTransactions[i] = &bs_db.Transaction{
					Hash:             t.Hash,
					BlockNumber:      t.BlockNumber,
					BlockHash:        t.BlockHash,
					FromAddress:      t.FromAddress,
					ToAddress:        t.ToAddress,
					Value:            t.Value,
					GasPrice:         t.GasPrice,
					GasLimit:         t.GasLimit,
					InputData:        t.InputData,
					Nonce:            t.Nonce,
					TransactionIndex: t.TransactionIndex,
					Type:             t.Type,
					ChainId:          t.ChainId,
					R:                t.R,
					S:                t.S,
				}
			}
			if err := tx.CreateInBatches(dbTransactions, 100).Error; err != nil {
				return fmt.Errorf("批量插入交易失败: %w", err)
			}
		}

		// 批量插入收据
		if len(receipts) > 0 {
			dbReceipts := make([]*bs_db.Receipt, len(receipts))
			for i, r := range receipts {
				dbReceipts[i] = &bs_db.Receipt{
					Type:              r.Type,
					Root:              r.Root,
					Status:            r.Status,
					CumulativeGasUsed: r.CumulativeGasUsed,
					LogsBloom:         r.LogsBloom,
					Logs:              r.Logs,
					TransactionHash:   r.TransactionHash,
					ContractAddress:   r.ContractAddress,
					GasUsed:           r.GasUsed,
					BlockHash:         r.BlockHash,
					BlockNumber:       r.BlockNumber,
					TransactionIndex:  r.TransactionIndex,
					FromAddress:       r.FromAddress,
					ToAddress:         r.ToAddress,
				}
			}
			if err := tx.CreateInBatches(dbReceipts, 100).Error; err != nil {
				return fmt.Errorf("批量插入收据失败: %w", err)
			}
		}

		// 批量插入日志
		if len(logs) > 0 {
			dbLogs := make([]*bs_db.Log, len(logs))
			for i, l := range logs {
				dbLogs[i] = &bs_db.Log{
					Address:          l.Address,
					Topics:           *l.Topics,
					Data:             l.Data,
					BlockNumber:      l.BlockNumber,
					TransactionHash:  l.TransactionHash,
					TransactionIndex: l.TransactionIndex,
					BlockHash:        l.BlockHash,
					LogIndex:         l.LogIndex,
					Removed:          l.Removed,
				}
			}
			if err := tx.CreateInBatches(dbLogs, 100).Error; err != nil {
				return fmt.Errorf("批量插入日志失败: %w", err)
			}
		}

		return nil
	})
}

type Block struct {
	ParentHash       string
	Miner            string
	StateRoot        string
	TransactionsRoot string
	ReceiptsRoot     string
	LogsBloom        string
	Difficulty       string
	Number           string
	GasLimit         uint64
	GasUsed          uint64
	Timestamp        uint64
	ExtraData        string
	BaseFeePerGas    string
	Transactions     *[]string `gorm:"type:jsonb"`
	Hash             string
	Size             uint64
}

func toLocalBlock(block *types.Block, txs []*bs_eth.Transaction) *Block {
	var transactions []string
	for _, tx := range txs {
		transactions = append(transactions, tx.Hash().String())
	}

	extra := block.Extra()
	signerPub := extra[len(extra)-gmsm.PublicKeyLength:]
	pub := gmsm.DecompressPubkey(signerPub)
	addr := gmsm.PubkeyToAddress(*pub)

	return &Block{
		ParentHash:       block.ParentHash().String(),
		Miner:            addr.String(), //block.Coinbase().String(),
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
		Transactions:     &transactions,
		Hash:             block.Hash().String(),
		Size:             uint64(block.Size()),
	}
}

type Transaction struct {
	Hash             string
	BlockNumber      string
	BlockHash        string
	FromAddress      string
	ToAddress        string
	Value            string
	GasPrice         string
	GasLimit         uint64
	InputData        string
	Nonce            uint64
	TransactionIndex uint64
	Type             uint8
	ChainId          uint64
	R                string
	S                string
}

func toLocalTransaction(block *types.Block, tx *bs_eth.Transaction, index uint64) *Transaction {
	R, S := tx.RawSignatureValues()
	var From, To string
	if tx.From() != nil {
		From = tx.From().String()
	}
	if tx.To() != nil {
		To = tx.To().String()
	}

	return &Transaction{
		Hash:             tx.Hash().String(),
		BlockNumber:      block.Number().String(),
		BlockHash:        block.Hash().String(),
		FromAddress:      From,
		ToAddress:        To,
		Value:            tx.Value().String(),
		GasPrice:         tx.GasPrice().String(),
		GasLimit:         tx.Gas(),
		InputData:        hex.EncodeToString(tx.Data()),
		Nonce:            tx.Nonce(),
		TransactionIndex: index,
		Type:             tx.Type(),
		ChainId:          tx.ChainId().Uint64(),
		R:                R.String(),
		S:                S.String(),
	}
}

type Receipt struct {
	Type              uint8  `json:"type"`
	Root              string `json:"root"`
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulative_gas_used"`
	LogsBloom         string `json:"logs_bloom"`
	Logs              int    `json:"logs"`
	TransactionHash   string `json:"transaction_hash"`
	ContractAddress   string `json:"contract_address"`
	GasUsed           uint64 `json:"gas_used"           `
	BlockHash         string `json:"block_hash"`
	BlockNumber       string `json:"block_number"`
	TransactionIndex  uint   `json:"transaction_index"`
	FromAddress       string `json:"from_address"`
	ToAddress         string `json:"to_address"`
}

func toLocalReceipt(receipt *types.Receipt, From, To *common.Address) *Receipt {
	var Root string
	var Status uint64
	var ToAddress string
	if len(receipt.PostState) > 0 {
		Root = hex.EncodeToString(receipt.PostState)
	} else {
		Status = receipt.Status
	}
	if To != nil {
		ToAddress = To.String()
	}

	return &Receipt{
		Type:              receipt.Type,
		Root:              Root,
		Status:            Status,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		LogsBloom:         hex.EncodeToString(receipt.Bloom.Bytes()),
		Logs:              len(receipt.Logs),
		TransactionHash:   receipt.TxHash.String(),
		ContractAddress:   receipt.ContractAddress.String(),
		GasUsed:           receipt.GasUsed,
		BlockHash:         receipt.BlockHash.String(),
		BlockNumber:       receipt.BlockNumber.String(),
		TransactionIndex:  receipt.TransactionIndex,
		FromAddress:       From.String(),
		ToAddress:         ToAddress,
	}
}

type Log struct {
	Address          string    `json:"address"`
	Topics           *[]string `json:"topics"  gorm:"type:jsonb"`
	Data             string    `json:"data"`
	BlockNumber      uint64    `json:"block_number"`
	TransactionHash  string    `json:"transaction_hash"`
	TransactionIndex uint      `json:"transaction_index"`
	BlockHash        string    `json:"block_hash"`
	LogIndex         uint      `json:"log_index"`
	Removed          bool      `json:"removed"`
}

func toLocalLog(log *types.Log) *Log {
	var Topics []string
	for _, topic := range log.Topics {
		Topics = append(Topics, hex.EncodeToString(topic.Bytes()))
	}

	return &Log{
		Address:          log.Address.String(),
		Topics:           &Topics,
		Data:             hex.EncodeToString(log.Data),
		BlockNumber:      log.BlockNumber,
		TransactionHash:  log.TxHash.String(),
		TransactionIndex: log.TxIndex,
		BlockHash:        log.BlockHash.String(),
		LogIndex:         log.Index,
		Removed:          log.Removed,
	}
}
