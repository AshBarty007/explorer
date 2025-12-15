package bs_eth

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/tjfoc/gmsm/sm2"
	"log"
	"math/big"
	"sync"
	"testing"
	"time"
)

const (
	rpcURL        = "http://192.168.120.32:8545"                                       // 替换为你的节点地址（如 Ganache）
	privateKeyHex = "39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d" // 发送方私钥（无 0x 前缀）
	toAddressStr  = "0x742d35Cc6634C0532925a3b8D4C9db96C0f7c3F8"                       // 接收地址
	gasLimit      = uint64(2100000)                                                    // 标准 ETH 转账 gas
	gasPriceGwei  = int64(20)                                                          // gas price (Gwei)
	numTx         = 5000                                                               // 总交易数
	concurrency   = 100                                                                // 并发 goroutine 数
)

func TestTps(t *testing.T) {
	client, err := Dial(rpcURL)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}
	defer client.Close()

	privateKey, err := HexToSM2(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}
	publicKey := privateKey.Public().(*sm2.PublicKey)

	fromAddress := PubkeyToAddress(*publicKey)

	// 获取 nonce（注意：并发时需预分配 nonce，否则会冲突）
	startNonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to get pending nonce:", err)
	}

	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal("Failed to get balance:", err)
	}
	fmt.Printf("From: %s, Balance: %s wei\n", fromAddress.Hex(), balance.String())

	gasPrice := big.NewInt(gasPriceGwei)
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(1e9)) // 转为 wei

	var wg sync.WaitGroup
	txChan := make(chan string, numTx)
	errChan := make(chan error, numTx)

	// 启动消费者协程收集结果
	go func() {
		wg.Add(1)
		defer wg.Done()
		//index := 0
		//for tx := range txChan {
		//	index++
		//	fmt.Println(index, tx)
		//}
	}()

	// 分配 nonce（每个 goroutine 有自己的起始 nonce 段）
	noncePerWorker := numTx / concurrency
	remaining := numTx % concurrency

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		start := startNonce + uint64(i*noncePerWorker)
		count := noncePerWorker
		if i == concurrency-1 {
			count += remaining // 最后一个 worker 处理余数
		}

		go func(workerID, txCount int, initialNonce uint64) {
			defer wg.Done()
			nonce := initialNonce
			//ix := 0
			for j := 0; j < txCount; j++ {
				//value := big.NewInt(1) // 转 1 wei（极小金额）
				//toAddress := common.HexToAddress(toAddressStr)
				//
				//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
				//
				//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
				//if err != nil {
				//	errChan <- fmt.Errorf("worker %d: sign error: %v", workerID, err)
				//	continue
				//}
				//
				//err = client.SendTransaction(context.Background(), signedTx)
				//if err != nil {
				//	errChan <- fmt.Errorf("worker %d: send error: %v", workerID, err)
				//	continue
				//}
				txHash, errMsg := Tx(client, nonce)
				if errMsg != nil {
					errChan <- fmt.Errorf("worker %d: send error: %v", workerID, errMsg)
					continue
				}
				txChan <- txHash

				//txChan <- txHash

				nonce++
			}
			//fmt.Println(workerID, ix, nonce)
		}(i, count, start)
	}

	// 等待所有发送完成
	go func() {
		wg.Wait()
		close(txChan)
		close(errChan)
	}()

	// 等待结束并统计
	totalSent := 0
	for range txChan {
		totalSent++
	}

	// 打印错误（如有）
	//for err := range errChan {
	//	log.Println("Error:", err)
	//}

	duration := time.Since(startTime)
	tps := float64(totalSent) / duration.Seconds()

	fmt.Printf("\n✅ Sent %d transactions in %v\n", totalSent, duration)
	fmt.Printf("📈 Estimated TPS: %.2f\n", tps)
}

func Tx(client *Client, nonce uint64) (string, error) {
	testKey, _ := HexToSM2(p1)
	toAddr := common.HexToAddress(toAddressStr)
	to := common.Address{}
	to.SetBytes(toAddr[:])

	gmTx := GmTx{
		ChainID: big.NewInt(1),
		Nonce:   nonce,
		//To:       &to,
		//Value:    big.NewInt(1000),
		Gas:      gasLimit,
		GasPrice: big.NewInt(params.InitialBaseFee),
		Data:     bytecode,
		//R, S      *big.Int,
		PublicKey: FromSM2Pub(&testKey.PublicKey),
	}
	h := Hash(gmTx)
	r, s, err := Sign(h[:], testKey)
	if err != nil {
		return "", errors.New("sign error")
	}
	gmTx.R = r
	gmTx.S = s
	tx := GmTransaction{
		inner: gmTx,
		time:  time.Now(),
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		return "", errors.New("Marshal error")
	}

	encodeTx := hexutil.Encode(data)
	var txHash string
	err = client.c.CallContext(context.Background(), &txHash, "eth_sendRawTransaction", encodeTx)
	if err != nil {
		return "", err
	}
	return txHash, nil

	//var result interface{}
	//err = client.c.CallContext(context.Background(), &result, "eth_getTransactionByHash", txHash)
	//if err != nil {
	//	return err
	//}
	//
	//return nil
	//fmt.Println("已上链，上链数据: ", result)
}
