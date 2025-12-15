package bshttp

import (
	"blockchain_services/config"
	bseth "blockchain_services/ethclient"
	bsdb "blockchain_services/postgres"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"strconv"
)

type AccountItem struct {
	Address          string `json:"address"`
	Balance          string `json:"balance"`
	Erc20            string `json:"erc20"`
	Erc404           string `json:"erc404"`
	TotalTransaction int64  `json:"total_transaction"`
}

type TokenItem struct {
	Address          string `json:"address"`
	Name             string `json:"name"`
	Supply           string `json:"supply"`
	Holders          string `json:"holders"`
	CREATOR          string `json:"creator"`
	TotalTransaction int64  `json:"total_transaction"`
}

func GetAccountsList(url string) ([]AccountItem, error) {
	var accounts []AccountItem
	addresses, err := bsdb.GetAddresses()
	if err != nil {
		return nil, fmt.Errorf("获取地址列表失败: %w", err)
	}

	client, err := bseth.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("连接ETH客户端失败: %w", err)
	}
	defer client.Close()

	for _, addr := range addresses {
		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
		if err != nil {
			log.Printf("查询地址 %s 余额失败: %v", addr, err)
			balance = big.NewInt(0) // 使用默认值继续处理
		}
		
		transactions, err := bsdb.GetTransactionCountByAddress(addr)
		if err != nil {
			log.Printf("查询地址 %s 交易数失败: %v", addr, err)
			transactions = 0 // 使用默认值继续处理
		}
		
		accounts = append(accounts, AccountItem{
			Address:          addr,
			Balance:          balance.String(),
			TotalTransaction: transactions,
		})
	}

	return accounts, nil
}

func GetAccountDetail(url string, addr string) (AccountItem, error) {
	var account AccountItem

	client, err := bseth.Dial(url)
	if err != nil {
		return account, fmt.Errorf("连接ETH客户端失败: %w", err)
	}
	defer client.Close()

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	if err != nil {
		return account, fmt.Errorf("查询余额失败: %w", err)
	}
	
	transactions, err := bsdb.GetTransactionCountByAddress(addr)
	if err != nil {
		return account, fmt.Errorf("查询交易数失败: %w", err)
	}
	
	erc20Balance := bseth.TokenBalance(url, config.Erc20ContractAddr, addr)
	erc404Balance := bseth.TokenBalance(url, config.Erc404ContractAddr, addr)
	
	account = AccountItem{
		Address:          addr,
		Balance:          balance.String(),
		Erc20:            erc20Balance,
		Erc404:           erc404Balance,
		TotalTransaction: transactions,
	}

	return account, nil
}

func GetTokensList(url string) {}

func GetTokenDetail(url string, addr string) {}

func parseNonNegativeInt(s string, defaultValue int) (int, error) {
	if s == "" {
		return defaultValue, nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, &strconv.NumError{Func: "parseNonNegativeInt", Num: s, Err: strconv.ErrRange}
	}
	return n, nil
}
