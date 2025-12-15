package config

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ethParams struct {
	TestUrl   string
	CanaryUrl string
	ChainID   *big.Int
	AdminKey  string
	AdminAddr string
	//ContractAddr common.Address
}

var EthParams = &ethParams{
	TestUrl:   "http://192.168.120.32:8545",
	CanaryUrl: "http://192.168.10.127:8545",
	ChainID:   big.NewInt(1),
	AdminKey:  "39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d",
	AdminAddr: "0xE25583099BA105D9ec0A67f5Ae86D90e50036425",
	//ContractAddr: common.HexToAddress("0x4B6Ea59a4CE0406E98FA6E29440af027dA33B970")
}

var (
	IpfsUrl = "http://192.168.90.141:30001/api/v0/add?stream-channels=true&pin=false&wrap-with-directory=false&progress=false"

	GrpcPort = "9965"

	TestUrl            = "http://192.168.120.33:8545"
	CanaryUrl          = "http://192.168.10.127:8545"
	RpcUrl             = "http://192.168.120.33:8545"
	ChainID            = big.NewInt(1)
	ProofContractAddr  = common.HexToAddress("0x23C5f582BAEdE953e6e2F5b8Dd680d5B97B39E78")
	Erc20ContractAddr  = common.HexToAddress("0x0e4bb0551b5a288addfc45e971ff5ac8d66889f5")
	Erc404ContractAddr = common.HexToAddress("0x8fde581e3e32bc98b6c0a30e663dba63643e987f")

	RedisAddr     = "192.168.90.179:6379"
	RedisPassword = "dev@123456"
	RedisDB       = 0

	DbHost     = "127.0.0.1" //"192.168.10.126"
	DbUsername = "eth"
	DbPassword = "123456"
	DbName     = "eth_explorer"
	DbPort     = "5432"
)

// InitConfig 设置是否读取环境变量，不使用内置值
func InitConfig(b bool) {
	DbHost = "192.168.10.126"
	DbUsername = "eth"
	DbPassword = "123456"
	DbName = "eth_explorer"
	DbPort = "5432"
}
