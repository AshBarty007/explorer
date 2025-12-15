package tests

import (
	"blockchain_services/ethclient"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/gmsm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestTicker(T *testing.T) {
	//ticker := time.NewTicker(5 * time.Second)
	//defer ticker.Stop()
	//
	//var i uint
	//
	//for t := range ticker.C {
	//	i++
	//
	//	log.Println("当前时间:", i, t.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	//
	//	if i%2 == 0 {
	//		time.Sleep(2 * time.Second)
	//	}
	//
	//	log.Println("当前时间:", i, t.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	//}

	start := 1
	end := 10

	for ; start <= end-1; start++ {
		if start == 5 {
			break
		}
		fmt.Println(start, start <= end-1)
	}

	fmt.Println("=================================")

	for ; start <= end-1; start++ {
		fmt.Println(start, start <= end-1)
	}

	fmt.Println(start, end, start <= end)
}

func TestSync(t *testing.T) {
	//var start, end uint64

	//block 109366 included 20 txs
	//

	etClient, err := bs_eth.Dial("http://192.168.10.126:8545")
	if err != nil {
		t.Fatal(err)
	}
	defer etClient.Close()

	dsn := fmt.Sprintf("host=192.168.10.126 user=eth password=123456 dbname=eth_explorer port=5432 sslmode=disable TimeZone=UTC")
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	for i := 109888; i <= 109899; i++ {
		err = DownloadBlock(etClient, db, big.NewInt(int64(i)))
		if err != nil {
			log.Println("区块号:", i, ", 错误信息:", err)
		}
	}

	//ticker := time.NewTicker(5 * time.Second)
	//defer ticker.Stop()
	//
	//for t := range ticker.C {
	//	fmt.Println("当前时间:", t.Format("2006-01-02 15:04:05"))
	//
	//	blockNumber, err := etClient.BlockNumber(context.Background())
	//	if err != nil {
	//		log.Println("没有查到最高区块:", err)
	//		continue
	//	}
	//
	//	if blockNumber > end+10 {
	//		end += 10
	//	} else {
	//		end = blockNumber
	//	}
	//
	//	for i := start; i <= end-1; i++ {
	//		err = DownloadBlock(etClient, db, big.NewInt(int64(i)))
	//		if err != nil {
	//			log.Println("区块号:", i, ", 错误信息:", err)
	//		}
	//	}
	//	log.Println("同步进度:", end, ", 最新区块:", blockNumber)
	//
	//	start = end
	//}

}

func DownloadBlock(etClient *bs_eth.Client, db *gorm.DB, number *big.Int) error {
	block, txs, err := etClient.BlockByNumber(context.Background(), number)
	if err != nil {
		log.Println("下载block数据失败:", number, err)
		return err
	}

	//1.存储区块数据
	b := toLocalBlock(block, txs)
	result := db.Create(&b)
	if result.Error != nil {
		log.Println("插入block数据失败:", result.Error)
		return result.Error
	} else {
		log.Printf("block数据已插入,number: %d\n", block.Number())
	}

	//2.存储交易数据
	if len(txs) != 0 {
		for i, tx := range txs {
			transaction := toLocalTransaction(block, tx, i)
			result := db.Create(&transaction)
			if result.Error != nil {
				log.Println("插入transaction数据失败:", result.Error)
				return result.Error
			} else {
				log.Printf("transaction数据已插入")
				//3.存储收据数据
				err = DownloadReceipt(etClient, db, transaction)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func DownloadReceipt(etClient *bs_eth.Client, db *gorm.DB, transaction *Transaction) error {
	//3.存储收据数据
	receipt, err := etClient.TransactionReceipt(context.Background(), common.HexToHash(transaction.Hash)) //28453
	if err != nil {
		log.Println("下载receipt数据失败:", transaction.Hash, err)
		return err
	}
	r := toLocalReceipt(receipt, bs_eth.HexToAddress(transaction.FromAddress), bs_eth.HexToAddress(transaction.ToAddress))
	fmt.Println("receipt", r.TransactionHash, r.BlockHash, r.BlockNumber)
	result := db.Create(&r)
	if result.Error != nil {
		log.Println("插入Receipt数据失败:", result.Error)
	} else {
		log.Printf("Receipt数据已插入")
	}

	//4.存储日志数据
	if r.Logs > 0 {
		for _, v := range receipt.Logs {
			data := toLocalLog(v)
			result := db.Create(&data)
			if result.Error != nil {
				log.Println("插入Log数据失败:", result.Error)
			} else {
				log.Printf("Log数据已插入")
			}
		}
	}
	return nil
}

func TestBlock(t *testing.T) {
	etClient, err := bs_eth.Dial("http://192.168.10.126:8545")
	if err != nil {
		t.Fatal(err)
	}
	defer etClient.Close()

	block, txs, err := etClient.BlockByNumber(context.Background(), big.NewInt(75648)) //28453
	if err != nil {
		t.Fatal(err)
	}

	// 连接数据库
	dsn := fmt.Sprintf("host=192.168.10.126 user=eth password=123456 dbname=eth_explorer port=5432 sslmode=disable TimeZone=UTC")
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	b := toLocalBlock(block, txs)
	result := db.Create(&b)
	if result.Error != nil {
		log.Println("插入测试数据失败:", result.Error)
	} else {
		fmt.Printf("测试数据已插入，number: %d\n", block.Number())
	}

	if len(txs) != 0 {
		tx := toLocalTransaction(block, txs[0], 1)
		fmt.Println(tx.Hash)
		result := db.Create(&tx)
		if result.Error != nil {
			log.Println("插入测试数据失败:", result.Error)
		} else {
			fmt.Printf("测试数据已插入")
		}
	}

}

func TestReceipt(t *testing.T) {
	etClient, err := bs_eth.Dial("http://192.168.10.126:8545")
	if err != nil {
		t.Fatal(err)
	}
	defer etClient.Close()

	tx, isPending, err := etClient.TransactionByHash(context.Background(), common.HexToHash("0x59f12016427ffc397b738497fd4b0ebac4e52971612165fe2126f90334356564"))
	if err != nil {
		t.Fatal(err)
	}
	if isPending {
		fmt.Println("isPending", isPending)
		return
	}
	fmt.Println("tx:", tx.From(), tx.To(), tx.Value(), tx.Hash())

	receipt, err := etClient.TransactionReceipt(context.Background(), common.HexToHash("0x59f12016427ffc397b738497fd4b0ebac4e52971612165fe2126f90334356564")) //28453
	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println("txHash:", receipt.TxHash, receipt.BlockHash, receipt.BlockNumber)

	// 连接数据库
	dsn := fmt.Sprintf("host=192.168.10.126 user=eth password=123456 dbname=eth_explorer port=5432 sslmode=disable TimeZone=UTC")
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	r := toLocalReceipt(receipt, tx.From(), tx.To())
	fmt.Println("receipt", r.TransactionHash, r.BlockHash, r.BlockNumber)
	result := db.Create(&r)
	if result.Error != nil {
		log.Println("插入Receipt测试数据失败:", result.Error)
	} else {
		fmt.Printf("Receipt测试数据已插入")
	}

	if r.Logs > 0 {
		for _, v := range receipt.Logs {
			data := toLocalLog(v)
			result := db.Create(&data)
			if result.Error != nil {
				log.Println("插入Log测试数据失败:", result.Error)
			} else {
				fmt.Printf("Log测试数据已插入")
			}
		}
	}
}

func toLocalBlock(block *types.Block, txs []*bs_eth.Transaction) *Block {
	var transactions StringSlice
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
		Transactions:     transactions,
		Hash:             block.Hash().String(),
		Size:             uint64(block.Size()),
	}
}

func toLocalTransaction(block *types.Block, tx *bs_eth.Transaction, index int) *Transaction {
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
		TransactionIndex: uint64(index),
		Type:             tx.Type(),
		ChainId:          tx.ChainId().Uint64(),
		R:                R.String(),
		S:                S.String(),
	}
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

func toLocalLog(log *types.Log) *Log {
	var Topics StringSlice
	for _, topic := range log.Topics {
		Topics = append(Topics, hex.EncodeToString(topic.Bytes()))
	}

	return &Log{
		Address:          log.Address.String(),
		Topics:           Topics,
		Data:             hex.EncodeToString(log.Data),
		BlockNumber:      log.BlockNumber,
		TransactionHash:  log.TxHash.String(),
		TransactionIndex: log.TxIndex,
		BlockHash:        log.BlockHash.String(),
		LogIndex:         log.Index,
		Removed:          log.Removed,
	}
}

/*
type Block struct {
	gorm.Model
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
type Transaction struct {
	gorm.Model
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
type Receipt struct {
	gorm.Model
	Type              uint8  `json:"type"` //type: "0x0"
	Root              string `json:"root"`
	Status            uint64 `json:"status"`              //status: "0x1",
	CumulativeGasUsed uint64 `json:"cumulative_gas_used"` //cumulativeGasUsed: 319971,
	LogsBloom         string `json:"logs_bloom"`          //logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	Logs              int    `json:"logs"`                //logs: [],
	TransactionHash   string `json:"transaction_hash"`    //transactionHash: "0x11f07f52853cd5a4eb6a618f6f75472f6b3233101d24a8560c23b8ff2bbb3ff4",
	ContractAddress   string `json:"contract_address"`    //contractAddress: "0x97077686617a8f4478863c23b73b7387ba35a802",
	GasUsed           uint64 `json:"gas_used"           ` //gasUsed: 319971,
	BlockHash         string `json:"block_hash"`          //blockHash: "0x0b71e27419dd0dd04f551c573c65bd6cacfe760b0e3d1fa12566115ed0e320fb",
	BlockNumber       string `json:"block_number"`        //blockNumber: 28453,
	TransactionIndex  uint   `json:"transaction_index"`   //transactionIndex: 0,
	FromAddress       string `json:"from_address"`
	ToAddress         string `json:"to_address"`
}
type Log struct {
	gorm.Model
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
*/

func TestTime(t *testing.T) {
	now := time.Now().Truncate(24 * time.Hour)
	sevenDaysAgo := now.AddDate(0, 0, -6) // 包含今天共7天：-6 到 0
	fmt.Println(time.Now().Unix())
	fmt.Println(sevenDaysAgo.Unix())
}
