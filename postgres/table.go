package bsdb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}
	if bytes, ok := value.([]byte); ok {
		return json.Unmarshal(bytes, s)
	}
	return fmt.Errorf("cannot scan %T into StringSlice", value)
}

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
	Transactions     StringSlice `gorm:"type:jsonb"`
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
	Timestamp        uint64
}

type Receipt struct {
	gorm.Model
	Type              uint8  `json:"type"`
	Root              string `json:"root"`
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulative_gas_used"`
	LogsBloom         string `json:"logs_bloom"`
	Logs              int    `json:"logs"`
	TransactionHash   string `json:"transaction_hash"`
	ContractAddress   string `json:"contract_address"`
	GasUsed           uint64 `json:"gas_used"`
	BlockHash         string `json:"block_hash"`
	BlockNumber       string `json:"block_number"`
	TransactionIndex  uint   `json:"transaction_index"`
	FromAddress       string `json:"from_address"`
	ToAddress         string `json:"to_address"`
	Timestamp         uint64 `json:"timestamp"`
}

type Log struct {
	gorm.Model
	Address          string      `json:"address"`
	Topics           StringSlice `json:"topics"  gorm:"type:jsonb"`
	Data             string      `json:"data"`
	BlockNumber      uint64      `json:"block_number"`
	TransactionHash  string      `json:"transaction_hash"`
	TransactionIndex uint        `json:"transaction_index"`
	BlockHash        string      `json:"block_hash"`
	LogIndex         uint        `json:"log_index"`
	Removed          bool        `json:"removed"`
}

type Account struct {
	gorm.Model
	Address            string     `json:"address" gorm:"uniqueIndex:idx_account_address;type:varchar(42);not null;comment:账户地址"`
	Type               string     `json:"type" gorm:"type:text;comment:账户类型"`
	Description        string     `json:"description" gorm:"type:text;comment:账户描述"`
	Balance            string     `json:"balance" gorm:"type:numeric(78,0);default:0;comment:账户余额(Wei)"`
	TotalTransaction   int64      `json:"total_transaction" gorm:"default:0;index:idx_account_tx_count;comment:总交易数"`
	FirstTransactionAt *time.Time `json:"first_transaction_at" gorm:"comment:首次交易时间"`
	LastTransactionAt  *time.Time `json:"last_transaction_at" gorm:"index:idx_account_last_tx;comment:最后交易时间"`
	Code               string     `json:"code" gorm:"type:text;comment:合约代码"`
	Token              []string   `json:"token" gorm:"type:jsonb;serializer:json;comment:代币地址列表"`
	NFT                []string   `json:"nft" gorm:"type:jsonb;serializer:json;comment:NFT地址列表"`
}

type Token struct {
	gorm.Model
	Address            string    `json:"address" gorm:"type:varchar(42);not null;index:idx_address;comment:代币地址"`
	Name               string    `json:"name" gorm:"type:varchar(100);comment:代币名称"`
	Symbol             string    `json:"symbol" gorm:"type:varchar(20);comment:代币符号"`
	Decimals           uint8     `json:"decimals" gorm:"default:18;comment:代币精度"`
	Supply             string    `json:"balance" gorm:"type:numeric(78,0);default:0;comment:代币总量"`
	Standard           string    `json:"standard" gorm:"type:varchar(20);default:'ERC20';index:idx_standard;comment:代币标准"`
	Description        string    `json:"description" gorm:"type:text;comment:代币描述"`
	Creator            string    `json:"creator" gorm:"type:varchar(42);not null;comment:创建者地址"`
	CreatedTime        time.Time `json:"created_time" gorm:"comment:创建时间"`
	CreatedHash        string    `json:"created_hash" gorm:"type:varchar(66);not null;comment:创建哈希"`
	CreatedBlockNumber uint64    `json:"created_block_number" gorm:"comment:创建区块高度"`
}

type NFT struct {
	gorm.Model
	Address            string    `json:"address" gorm:"type:varchar(42);not null;index:idx_address;comment:NFT地址"`
	Name               string    `json:"name" gorm:"type:varchar(100);comment:NFT名称"`
	Symbol             string    `json:"symbol" gorm:"type:varchar(20);comment:NFT符号"`
	Supply             string    `json:"balance" gorm:"type:numeric(78,0);default:0;comment:NFT总量"`
	Standard           string    `json:"standard" gorm:"type:varchar(20);default:'ERC721';index:idx_standard;comment:NFT标准"`
	Description        string    `json:"description" gorm:"type:text;comment:NFT描述"`
	TokenUri           string    `json:"tokenUri" gorm:"type:text;comment:元数据的链接"`
	Metadata           string    `json:"metadata" gorm:"type:jsonb;comment:NFT元数据(JSON格式)"`
	Creator            string    `json:"creator" gorm:"type:varchar(42);not null;comment:创建者地址"`
	CreatedTime        time.Time `json:"created_time" gorm:"comment:创建时间"`
	CreatedHash        string    `json:"created_hash" gorm:"type:varchar(66);not null;comment:创建哈希"`
	CreatedBlockNumber uint64    `json:"created_block_number" gorm:"comment:创建区块高度"`
}

type Contract struct {
	gorm.Model
	Address            string    `json:"address" gorm:"type:varchar(42);not null;uniqueIndex:idx_contract_address;comment:合约地址"`
	Creator            string    `json:"creator" gorm:"type:varchar(42);not null;index:idx_contract_creator;comment:创建者地址"`
	CreatedTime        time.Time `json:"created_time" gorm:"index:idx_contract_created_time;comment:创建时间"`
	CreatedHash        string    `json:"created_hash" gorm:"type:varchar(66);not null;comment:创建交易哈希"`
	CreatedBlockNumber uint64    `json:"created_block_number" gorm:"index:idx_contract_created_block;comment:创建区块高度"`
	Description        string    `json:"description" gorm:"type:text;comment:合约描述"`
}
