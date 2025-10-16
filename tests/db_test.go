package tests

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
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

func TestCreateDB(t *testing.T) {
	dsn := fmt.Sprintf("host=192.168.10.126 user=eth password=123456 dbname=eth_explorer port=5432 sslmode=disable TimeZone=UTC")
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	//检查是否创建
	if !db.Migrator().HasTable(&Block{}) {
		//创建
		err = db.Table("blocks").AutoMigrate(&Block{})
		if err != nil {
			panic("failed to migrate Block table")
		}
	}

	if !db.Migrator().HasTable(&Transaction{}) {
		//创建
		err = db.Table("transactions").AutoMigrate(&Transaction{})
		if err != nil {
			panic("failed to migrate Transaction table")
		}
	}

	if !db.Migrator().HasTable(&Receipt{}) {
		//创建
		err = db.Table("receipts").AutoMigrate(&Receipt{})
		if err != nil {
			panic("failed to migrate Receipt table")
		}
	}

	if !db.Migrator().HasTable(&Log{}) {
		//创建
		err = db.Table("logs").AutoMigrate(&Log{})
		if err != nil {
			panic("failed to migrate Log table")
		}
	}

}

func TestDropDB(t *testing.T) {
	dsn := fmt.Sprintf("host=192.168.10.126 user=eth password=123456 dbname=eth_explorer port=5432 sslmode=disable TimeZone=UTC")
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	//检查是否创建
	if db.Migrator().HasTable(&Block{}) {
		//删除
		err = db.Migrator().DropTable(&Block{})
		if err != nil {
			panic("failed to drop Block table")
		}
	}

	//检查是否创建
	if db.Migrator().HasTable(&Transaction{}) {
		//删除
		err = db.Migrator().DropTable(&Transaction{})
		if err != nil {
			panic("failed to drop Transaction table")
		}
	}

	//检查是否创建
	if db.Migrator().HasTable(&Receipt{}) {
		//删除
		err = db.Migrator().DropTable(&Receipt{})
		if err != nil {
			panic("failed to drop Receipt table")
		}
	}

	//检查是否创建
	if db.Migrator().HasTable(&Log{}) {
		//删除
		err = db.Migrator().DropTable(&Log{})
		if err != nil {
			panic("failed to drop Log table")
		}
	}

}

func TestFunc(t *testing.T) {
	dsn := fmt.Sprintf("host=192.168.10.126 user=eth password=123456 dbname=eth_explorer port=5432 sslmode=disable TimeZone=UTC")
	gormCfg := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormCfg)
	if err != nil {
		panic(err)
	}

	block := getBlockByNumber(db, 109895)
	fmt.Println("1.getBlockByNumber", block)

	tx := getTransactionByHash(db, "0x9d2344d8f260316ed4d967ade19ee43080a8aa76994ecde397f8a7fb737761cd")
	fmt.Println("2.getTransactionByHash", tx)

	txs := getTransactions(db, 2, 8)
	fmt.Println("3.getTransactions", len(txs))

	bs := getBlocks(db, 1, 5)
	fmt.Println("4.getBlocks", len(bs))

	addrCount, err := getAddressCount(db)
	fmt.Println("5.getAddressCount: ", addrCount, err)

	txCount, err := getTransactionCount(db)
	fmt.Println("6.getTransactionCount: ", txCount, err)

	blockCount, err := getBlockCount(db)
	fmt.Println("7.getBlockCount: ", blockCount, err)

	rp := getReceiptByHash(db, "0x9d2344d8f260316ed4d967ade19ee43080a8aa76994ecde397f8a7fb737761cd")
	fmt.Println("8.getReceiptByHash", rp)

	rps := getReceipts(db, 2, 8)
	fmt.Println("9.getReceipts", len(rps))

	rbt := getReceiptsByLast(db, 2)
	fmt.Println("10.getReceiptsByLast: ", len(rbt))

	tba := getTransactionsByAddress(db, "0x30E938B0630c02f394d17925fDb5fb046F70D452")
	fmt.Println("11.getTransactionsByAddress", len(tba))

	rba := getReceiptsByAddress(db, "0x30E938B0630c02f394d17925fDb5fb046F70D452")
	fmt.Println("12.getReceiptsByAddress", len(rba))

}

// 区块详情
func getBlockByNumber(db *gorm.DB, number int) Block {
	var block Block
	db.Where("number = ?", strconv.Itoa(number)).First(&block)
	return block
}

// 交易详情
func getTransactionByHash(db *gorm.DB, hash string) Transaction {
	var transaction Transaction
	db.Where("hash=?", hash).First(&transaction)
	return transaction
}

// 批量查询
func getTransactions(db *gorm.DB, start, end int) []Transaction {
	var transactions []Transaction
	db.
		Order("id ASC").
		Offset(start).
		Limit(end).
		Find(&transactions)
	return transactions
}

// 批量查询
func getBlocks(db *gorm.DB, start, end int) []Block {
	var blocks []Block
	db.
		Order("id ASC").
		Offset(start).
		Limit(end).
		Find(&blocks)
	return blocks
}

// 地址数量
func getAddressCount(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Raw("SELECT COUNT(DISTINCT from_address) FROM transactions").Scan(&count).Error
	return count, err
}

// 交易总数
func getTransactionCount(db *gorm.DB) (string, error) {
	var transaction Transaction
	result := db.Last(&transaction)
	if result.Error != nil {
		return "", result.Error
	}
	return strconv.Itoa(int(transaction.ID)), nil
}

// 区块总数
func getBlockCount(db *gorm.DB) (string, error) {
	var block Block
	result := db.Last(&block)
	if result.Error != nil {
		return "", result.Error
	}
	return strconv.Itoa(int(block.ID)), nil
}

// 查询收据
func getReceiptByHash(db *gorm.DB, hash string) Receipt {
	var receipt Receipt
	db.Where("transaction_hash=?", hash).First(&receipt)
	return receipt
}

// 批量查询查询收据
func getReceipts(db *gorm.DB, start, end int) []Receipt {
	var receipts []Receipt
	db.
		Order("id ASC").
		Offset(start).
		Limit(end).
		Find(&receipts)
	return receipts
}

// 查询最近7天收据 x
func getReceiptsByLast(db *gorm.DB, timestamp uint64) []Receipt {
	var receipts []Receipt
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	db.Where("created_at >= ?", sevenDaysAgo).Find(&receipts)
	return receipts
}

func getTransactionsByAddress(db *gorm.DB, address string) []Receipt {
	var receipts []Receipt
	db.Where("from_address = ?", address).Find(&receipts)
	return receipts
}

func getReceiptsByAddress(db *gorm.DB, address string) []Receipt {
	var receipts []Receipt
	db.Where("from_address = ?", address).Find(&receipts)
	return receipts
}
