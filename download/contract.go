package download

import (
	"blockchain_services/blockchain"
	db "blockchain_services/postgres"
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// detectContractType 检测合约类型
// 通过读取合约的公共变量 uint8 ContractType
// 0 代表 ERC20，1 代表 ERC721，否则代表其他合约
func detectContractType(etClient *blockchain.Client, contractAddr string) string {
	ctx := context.Background()
	address := common.HexToAddress(contractAddr)

	// 读取合约的 ContractType 变量
	// Solidity 公共变量会自动生成 getter 函数，函数签名为 ContractType()
	data := getFunctionSelector("ContractType()")

	msg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	result, err := etClient.CallContract(ctx, msg, nil)
	if err != nil {
		// 如果调用失败，返回默认类型
		return "OTHER"
	}

	// 解码返回的 uint8 值
	// uint8 在返回数据的最后一个字节（32字节对齐）
	if len(result) < 32 {
		return "OTHER"
	}

	contractType := result[31] // uint8 在最后字节

	// 根据值返回类型
	switch contractType {
	case 0:
		return "ERC20"
	case 1:
		return "ERC721"
	default:
		return "OTHER"
	}
}

// fetchTokenInfo 获取Token信息（名称、符号、精度、总量等）
func fetchTokenInfo(etClient *blockchain.Client, token *db.Token) error {
	ctx := context.Background()
	address := common.HexToAddress(token.Address)

	contractABI, err := abi.JSON(strings.NewReader(blockchain.ERC20_abi))
	if err != nil {
		return fmt.Errorf("解析ERC20 ABI失败: %w", err)
	}

	// 调用name()方法
	if name, err := callStringMethod(etClient, ctx, address, &contractABI, "name"); err == nil {
		token.Name = name
	}

	// 调用symbol()方法
	if symbol, err := callStringMethod(etClient, ctx, address, &contractABI, "symbol"); err == nil {
		token.Symbol = symbol
	}

	// 调用decimals()方法
	if decimals, err := callUint8Method(etClient, ctx, address, &contractABI, "decimals"); err == nil {
		token.Decimals = decimals
	}

	// 调用totalSupply()方法
	if supply, err := callUint256Method(etClient, ctx, address, &contractABI, "totalSupply"); err == nil {
		token.Supply = supply.String()
	}

	return nil
}

// fetchNFTInfo 获取NFT信息（名称、符号、总量等）
func fetchNFTInfo(etClient *blockchain.Client, nft *db.NFT) error {
	ctx := context.Background()
	address := common.HexToAddress(nft.Address)

	// 默认使用ERC721
	contractABI, err := abi.JSON(strings.NewReader(blockchain.ERC721_abi))
	if err != nil {
		return fmt.Errorf("读取NFT ABI失败: %w", err)
	}

	// 尝试调用name()方法
	if name, err := callStringMethod(etClient, ctx, address, &contractABI, "name"); err == nil {
		nft.Name = name
	}

	// 尝试调用symbol()方法
	if symbol, err := callStringMethod(etClient, ctx, address, &contractABI, "symbol"); err == nil {
		nft.Symbol = symbol
	}

	return nil
}

// callStringMethod 调用返回string的方法
func callStringMethod(etClient *blockchain.Client, ctx context.Context, address common.Address, contractABI *abi.ABI, method string) (string, error) {
	data, err := contractABI.Pack(method)
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	result, err := etClient.CallContract(ctx, msg, nil)
	if err != nil {
		return "", err
	}

	var ret string
	err = contractABI.UnpackIntoInterface(&ret, method, result)
	if err != nil {
		return "", err
	}

	return ret, nil
}

// callUint8Method 调用返回uint8的方法
func callUint8Method(etClient *blockchain.Client, ctx context.Context, address common.Address, contractABI *abi.ABI, method string) (uint8, error) {
	data, err := contractABI.Pack(method)
	if err != nil {
		return 0, err
	}

	msg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	result, err := etClient.CallContract(ctx, msg, nil)
	if err != nil {
		return 0, err
	}

	var ret uint8
	err = contractABI.UnpackIntoInterface(&ret, method, result)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

// callUint256Method 调用返回uint256的方法
func callUint256Method(etClient *blockchain.Client, ctx context.Context, address common.Address, contractABI *abi.ABI, method string) (*big.Int, error) {
	data, err := contractABI.Pack(method)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &address,
		Data: data,
	}

	result, err := etClient.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	var ret *big.Int
	err = contractABI.UnpackIntoInterface(&ret, method, result)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// updateAccount 更新账户信息
// isFromAddress: true表示是from_address（发送方），false表示是to_address（接收方）
func updateAccount(etClient *blockchain.Client, address string, txTime time.Time, isFromAddress bool) error {
	ctx := context.Background()

	// 查询账户是否已存在
	var existingAccount db.Account

	// 获取账户余额
	balance, err := etClient.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		fmt.Printf("获取账户余额失败 %s: %v", address, err)
	}

	// 统计该地址的交易数
	var txCount int64
	db.Db.Model(&db.Transaction{}).
		Where("from_address = ? OR to_address = ?", address, address).
		Count(&txCount)

	account := &db.Account{
		Address:           address,
		Balance:           balance.String(),
		TotalTransaction:  txCount,
		LastTransactionAt: &txTime,
	}

	// 只有from_address（发送方）且FirstTransactionAt为空时，才设置首次交易时间
	if isFromAddress && existingAccount.FirstTransactionAt == nil {
		account.FirstTransactionAt = &txTime
	}

	if err := db.WriteAccount(*account); err != nil {
		return fmt.Errorf("更新账户失败: %w", err)
	}

	return nil
}

func getFunctionSelector(signature string) []byte {
	hash := crypto.NewKeccakState()
	hash.Write([]byte(signature))
	return hash.Sum(nil)[:4]
}
