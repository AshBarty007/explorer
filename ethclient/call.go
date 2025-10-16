package bs_eth

import (
	"blockchain_services/config"
	"blockchain_services/ethclient/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/gmsm"
	"log"
)

var (
	rpcUrl       = config.RpcUrl
	chainID      = config.ChainID
	contractAddr = config.ProofContractAddr
)

func Call(key string, typeName string, ipfsHash string) (string, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer client.Close()

	privateKey, _ := gmsm.HexToSM2(key)

	proofContract, err := contracts.NewEthClientTransactor(contractAddr, client)
	if err != nil {
		fmt.Println("NewToken error : ", err)
		return "", err
	}

	opts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	tx, err := proofContract.Submit(opts, typeName, ipfsHash)
	if err != nil {
		return "", err
	}
	fmt.Println("交易哈希 : ", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func BalanceOf(address string) (string, error) {
	client, err := Dial(config.RpcUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	erc20Contract, err := contracts.NewContractsCaller(config.Erc20ContractAddr, client)
	if err != nil {
		fmt.Println("NewToken error : ", err)
	}

	opts := &bind.CallOpts{}
	balance, err := erc20Contract.BalanceOf(opts, common.HexToAddress(address))
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}
