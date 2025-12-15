package bshttp

import (
	"blockchain_services/config"
	bseth "blockchain_services/ethclient"
	bsdb "blockchain_services/postgres"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"
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

func GetAccountsList(url string) []AccountItem {
	var accounts []AccountItem
	addresses := bsdb.GetAddresses()

	client, err := bseth.Dial(url)
	if err != nil {
		log.Fatal("Connect ETH Error: ", err)
		return accounts
	}
	defer client.Close()

	for _, addr := range addresses {
		balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
		transactions := bsdb.GetTransactionCountByAddress(addr)
		accounts = append(accounts, AccountItem{
			Address:          addr,
			Balance:          balance.String(),
			TotalTransaction: transactions,
		})
	}

	return accounts
}

func GetAccountDetail(url string, addr string) AccountItem {
	var account AccountItem

	client, err := bseth.Dial(url)
	if err != nil {
		log.Fatal("Connect ETH Error: ", err)
		return account
	}
	defer client.Close()

	balance, _ := client.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	transactions := bsdb.GetTransactionCountByAddress(addr)
	erc20Balance := bseth.TokenBalance(url, config.Erc20ContractAddr, addr)
	erc404Balance := bseth.TokenBalance(url, config.Erc404ContractAddr, addr)
	account = AccountItem{
		Address:          addr,
		Balance:          balance.String(),
		Erc20:            erc20Balance,
		Erc404:           erc404Balance,
		TotalTransaction: transactions,
	}

	return account
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
