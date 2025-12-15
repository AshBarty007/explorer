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

// Account 账户表 - 存储区块链地址账户信息
type Account struct {
	gorm.Model
	// Address 账户地址，唯一索引
	Address string `json:"address" gorm:"uniqueIndex:idx_account_address;type:varchar(42);not null;comment:账户地址"`

	// Balance 账户余额（Wei），使用NUMERIC类型存储大数
	Balance string `json:"balance" gorm:"type:numeric(78,0);default:0;comment:账户余额(Wei)"`

	// TotalTransaction 总交易数
	TotalTransaction int64 `json:"total_transaction" gorm:"default:0;index:idx_account_tx_count;comment:总交易数"`

	// FirstTransactionAt 首次交易时间
	FirstTransactionAt *time.Time `json:"first_transaction_at" gorm:"comment:首次交易时间"`

	// LastTransactionAt 最后交易时间
	LastTransactionAt *time.Time `json:"last_transaction_at" gorm:"index:idx_account_last_tx;comment:最后交易时间"`

	// BalanceUpdatedAt 余额更新时间，用于缓存控制
	BalanceUpdatedAt *time.Time `json:"balance_updated_at" gorm:"comment:余额更新时间"`

	// ERC20Tokens ERC20代币关联
	ERC20Tokens []Token `json:"erc20_tokens" gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`

	// NFTs NFT资产关联
	NFTs []NFT `json:"nfts" gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE"`
}

// TableName 指定Account表名
func (Account) TableName() string {
	return "accounts"
}

// Token ERC20代币表 - 存储账户持有的ERC20代币信息
type Token struct {
	gorm.Model
	// AccountID 账户ID，外键关联Account
	AccountID uint `json:"account_id" gorm:"index:idx_token_account;not null;comment:账户ID"`

	// TokenAddress 代币合约地址
	TokenAddress string `json:"token_address" gorm:"type:varchar(42);not null;index:idx_token_address;comment:代币合约地址"`

	// Symbol 代币符号（如ETH, USDT）
	Symbol string `json:"symbol" gorm:"type:varchar(20);comment:代币符号"`

	// Name 代币名称
	Name string `json:"name" gorm:"type:varchar(100);comment:代币名称"`

	// Decimals 代币精度
	Decimals uint8 `json:"decimals" gorm:"default:18;comment:代币精度"`

	// Balance 代币余额，使用NUMERIC类型存储大数
	Balance string `json:"balance" gorm:"type:numeric(78,0);default:0;comment:代币余额"`

	// BalanceUpdatedAt 余额更新时间
	BalanceUpdatedAt *time.Time `json:"balance_updated_at" gorm:"comment:余额更新时间"`

	// Account 账户关联
	Account Account `json:"account" gorm:"foreignKey:AccountID"`
}

// TableName 指定Token表名
func (Token) TableName() string {
	return "tokens"
}

// NFT NFT资产表 - 存储账户持有的NFT资产信息
type NFT struct {
	gorm.Model
	// AccountID 账户ID，外键关联Account
	AccountID uint `json:"account_id" gorm:"index:idx_nft_account;not null;comment:账户ID"`

	// ContractAddress NFT合约地址
	ContractAddress string `json:"contract_address" gorm:"type:varchar(42);not null;index:idx_nft_contract;comment:NFT合约地址"`

	// TokenID Token ID
	TokenID string `json:"token_id" gorm:"type:varchar(100);not null;comment:Token ID"`

	// Standard NFT标准（ERC721, ERC1155, ERC404等）
	Standard string `json:"standard" gorm:"type:varchar(20);default:'ERC721';index:idx_nft_standard;comment:NFT标准"`

	// Name NFT名称
	Name string `json:"name" gorm:"type:varchar(255);comment:NFT名称"`

	// Description NFT描述
	Description string `json:"description" gorm:"type:text;comment:NFT描述"`

	// Image NFT图片URL
	Image string `json:"image" gorm:"type:varchar(500);comment:NFT图片URL"`

	// Metadata NFT元数据，使用JSONB类型存储
	Metadata string `json:"metadata" gorm:"type:jsonb;comment:NFT元数据(JSON格式)"`

	// MetadataUpdatedAt 元数据更新时间
	MetadataUpdatedAt *time.Time `json:"metadata_updated_at" gorm:"comment:元数据更新时间"`

	// Account 账户关联
	Account Account `json:"account" gorm:"foreignKey:AccountID"`
}

// TableName 指定NFT表名
func (NFT) TableName() string {
	return "nfts"
}
