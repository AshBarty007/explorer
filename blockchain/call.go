package blockchain

import (
	"blockchain_services/config"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var rpcUrl = config.TestUrl

func Call(address string, data []byte) ([]byte, error) {
	etClient, err := Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	defer etClient.Close()

	ctx := context.Background()
	contractAddress := common.HexToAddress(address)


	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := etClient.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}	

func BalanceOf(address string) (string, error) {
	client, err := Dial(rpcUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)

	return balance.String(), err
}

func TokenBalance(contract common.Address, account string) string {
	client, err := Dial(rpcUrl)
	if err != nil {
		fmt.Println("Failed to connect to Ethereum client:", err)
		return "0"
	}
	defer client.Close()

	addr := common.HexToAddress(account)
	m := crypto.Keccak256([]byte("balanceOf(address)"))[:4]
	a := make([]byte, 32)
	copy(a[12:], addr[:])
	data := append(m, a...)

	arg := map[string]interface{}{
		"to":   contract.Hex(),
		"data": hexutil.Encode(data),
	}
	var out hexutil.Bytes
	_ = client.Client().CallContext(context.Background(), &out, "eth_call", arg, "latest")
	if len(out) == 0 {
		fmt.Println("fail to get TokenBalance, the result is nil.")
		return "0"
	}
	return new(big.Int).SetBytes(out).String()
}
