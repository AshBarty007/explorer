package bsdb

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

type AccountItem struct {
	gorm.Model
	Address          string       `json:"address" gorm:"uniqueIndex"` // Unique for accounts
	Balance          string       `json:"balance"`                    // ETH balance as string (for big ints)
	TotalTransaction int64        `json:"total_transaction"`
	ERC20Tokens      []ERC20Token `gorm:"foreignKey:AccountID"` // One-to-many: multiple ERC-20s
	NFTs             []NFT        `gorm:"foreignKey:AccountID"` // One-to-many: multiple NFTs (or ERC-404 if custom)
}

type ERC20Token struct {
	gorm.Model
	AccountID    uint   `json:"account_id" gorm:"index"` // Foreign key to AccountItem
	TokenAddress string `json:"token_address"`           // e.g., contract address like "0x..."
	Symbol       string `json:"symbol"`                  // Optional: e.g., "USDT"
	Balance      string `json:"balance"`                 // Token balance as string (for decimals/big ints)
	// Add more fields if needed, e.g., Decimals int, LastUpdated time.Time
}

type NFT struct {
	gorm.Model
	AccountID       uint   `json:"account_id" gorm:"index"` // Foreign key to AccountItem
	ContractAddress string `json:"contract_address"`        // e.g., "0x..." for the NFT collection
	TokenID         string `json:"token_id"`                // NFT's unique ID (as string for big ints)
	Standard        string `json:"standard"`                // e.g., "ERC721", "ERC1155", or "ERC404"
	Metadata        string `json:"metadata"`                // Optional: JSON string for traits/URI
	// Add more fields if needed, e.g., Quantity int (for ERC-1155), AcquiredAt time.Time
}
