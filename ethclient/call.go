package bs_eth

import (
	"blockchain_services/config"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

var (
	rpcUrl       = config.TestUrl
	chainID      = config.ChainID
	contractAddr = config.ProofContractAddr
)

// Call 用于测试合于simpleStorage
func Call(key string, typeName string, ipfsHash string) (string, error) {
	//client, err := Dial(rpcUrl)
	//if err != nil {
	//	log.Println(err)
	//	return "", err
	//}
	//defer client.Close()
	//
	//privateKey, _ := gmsm.HexToSM2(key)
	//
	//proofContract, err := contracts.NewEthClientTransactor(contractAddr, client)
	//if err != nil {
	//	fmt.Println("NewToken error : ", err)
	//	return "", err
	//}
	//
	//opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	//tx, err := proofContract.Submit(opts, typeName, ipfsHash)
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println("交易哈希 : ", tx.Hash().Hex())
	//
	//return tx.Hash().Hex(), nil
	return "", nil
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

func TokenBalance(url string, contract common.Address, account string) string {
	client, err := Dial(url)
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
