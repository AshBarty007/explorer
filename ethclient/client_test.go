package bs_eth

import (
	"blockchain_services/ethclient/contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/gmsm"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"testing"
)

func TestSendTransaction(t *testing.T) {
	client, err := Dial("http://192.168.10.128:8545")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	chainID := big.NewInt(1) //3151908
	testKey, _ := gmsm.HexToSM2("39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d")
	testAddr := gmsm.PubkeyToAddress(testKey.PublicKey)
	toAddress := HexToAddress("0x29266a99b4b968776c1ef06f28b98635110cf0ff")
	nonce, err := client.NonceAt(context.Background(), testAddr, nil)
	if err != nil {
		t.Fatal(err)
	}

	//signer := types.LatestSignerForChainID(chainID)
	tx, err := SignNewTx(testKey, chainID, &LegacyTx{
		Nonce:    nonce,
		To:       toAddress, //&common.Address{2},
		Value:    big.NewInt(1).Mul(big.NewInt(1000000000), big.NewInt(1000000000)),
		Gas:      22000,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})

	fmt.Println("toAddress", toAddress)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("tx hash: ", tx.Hash().Hex())
}

func TestContracts(t *testing.T) {
	client, err := Dial("http://192.168.10.126:8545")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	erc20Contract, err := contracts.NewContractsCaller(common.HexToAddress("0xcaF566B6936A73C2913cA4CA77Dd4E95CcD0eD03"), client)
	if err != nil {
		fmt.Println("NewToken error : ", err)
	}

	opts := &bind.CallOpts{}
	balance, err := erc20Contract.BalanceOf(opts, common.HexToAddress("0xE25583099BA105D9ec0A67f5Ae86D90e50036425"))
	if err != nil {
		t.Fatal(err)
	}

	log.Print("balanceOf : ", balance)
}

func TestTxHash(t *testing.T) {
	client, err := Dial("http://192.168.10.126:8545")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash("0x59f12016427ffc397b738497fd4b0ebac4e52971612165fe2126f90334356564"))
	if err != nil {
		t.Fatal(err)
	}
	if isPending {
		fmt.Println("isPending", isPending)
		return
	}
	fmt.Println("tx:", tx.To(), tx.Value(), tx.Hash())
}
