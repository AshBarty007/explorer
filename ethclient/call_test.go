package bs_eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var chainId = big.NewInt(18197)
var url = "http://127.0.0.1:8545" //"http://192.168.10.128:8545"
var contractAddress = common.HexToAddress("0xc7b7a73cb27f069acd201768b3cdae513cdd76f0")

// 5. 合约字节码（替换为你的 Solidity 编译输出的 .bin 内容）
// 0.5.6
// var contractBytecode = "6080604052606460005534801561001557600080fd5b5060c6806100246000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80632e64cec11460375780636057361d146053575b600080fd5b603d607e565b6040518082815260200191505060405180910390f35b607c60048036036020811015606757600080fd5b81019080803590602001909291905050506087565b005b60008054905090565b806000819055505056fea265627a7a7231582058462df485346139397b6409edbf8ea93d3b0fcb98284f173eca0f6a49580e6264736f6c63430005110032" // 你的合约bin
// 0.8.2
var contractBytecode = "608060405234801561001057600080fd5b5060017fc670f864e1cbd31b04bbb7f207ac27ffee1b5925910425f98ee1448a774a05e960405160405180910390a27f1e69321bbdc0510e8b5f62e2a1bbbed6143ae12e782a1783a55e8fd5019f3b0f60405161006c90610175565b60405180910390a17fc0d505515bf144d2d18fa4d0308b7069fed35d610375ae95bca9cd33f25bde0a6040516100a1906101e1565b60405180910390a17f7c4352354ddbb23a20b7890913a0983763d5052c208f998378f159f5e6bd0e306040516100d69061024d565b60405180910390a17fec5dc4146de24512ee572c21b5a93986c75fdd0eea5c66655dca6e308a38550760405161010b906102b9565b60405180910390a16102d9565b600082825260208201905092915050565b7f68656c6c6f20776f726c64210000000000000000000000000000000000000000600082015250565b600061015f600c83610118565b915061016a82610129565b602082019050919050565b6000602082019050818103600083015261018e81610152565b9050919050565b7f7468697320697320612074657374210000000000000000000000000000000000600082015250565b60006101cb600f83610118565b91506101d682610195565b602082019050919050565b600060208201905081810360008301526101fa816101be565b9050919050565b7f676f6f64206c75636b20746f20796f7521000000000000000000000000000000600082015250565b6000610237601183610118565b915061024282610201565b602082019050919050565b600060208201905081810360008301526102668161022a565b9050919050565b7f6e69636520746f2073656520796f752100000000000000000000000000000000600082015250565b60006102a3601083610118565b91506102ae8261026d565b602082019050919050565b600060208201905081810360008301526102d281610296565b9050919050565b603f806102e76000396000f3fe6080604052600080fdfea26469706673582212209ba36049e149d523abc837bb02974b059268bbe9ab1e862fc137db2b4ee359c264736f6c634300081e0033" // 你的合约bin

func TestDepoly(t *testing.T) {
	// 1. 连接到以太坊节点（可以是 Infura、Alchemy 或本地节点）
	client, err := Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// 2. 加载私钥（⚠️ 仅用于测试！不要在生产中硬编码私钥）
	privateKey, _ := HexToSM2("39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d")

	// 3. 获取发送地址和 nonce
	fromAddress := PubkeyToAddress(privateKey.PublicKey)

	bal, _ := client.BalanceAt(context.Background(), fromAddress, nil)
	fmt.Println("fromAddress, bal:", fromAddress, bal)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// 4. 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// 6. 构造交易
	gasLimit := uint64(500000) // 根据合约复杂度调整

	// 创建一个 to 为空的交易，data 为合约字节码
	inner := GmTx{
		ChainID: big.NewInt(1),
		Nonce:   nonce,
		//To:      &contractAddress,
		//Value:    big.NewInt(1).Mul(big.NewInt(1000000000), big.NewInt(1000000000)),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     common.FromHex(contractBytecode),
		//R, S      *big.Int,
		PublicKey: sm2.Compress(&privateKey.PublicKey),
	}

	// 7. 签名交易
	h := Hash(inner)
	r, s, err := Sign(h[:], privateKey)
	if err != nil {
		t.Fatal(err)
	}
	inner.R = r
	inner.S = s
	signedTx := GmTransaction{
		inner: inner,
		time:  time.Now(),
	}

	// 8. 序列化为 RLP 编码
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v", err)
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)

	// 9. 打印原始交易（可用于调试）
	// fmt.Printf("Raw Transaction (hex): 0x%s\n", rawTxHex)

	// 10. 调用 eth_sendRawTransaction
	var result interface{}
	err = client.Client().CallContext(context.Background(), &result, "eth_sendRawTransaction", "0x"+rawTxHex)
	if err != nil {
		log.Fatalf("Failed to send raw transaction: %v", err)
	}

	// 11. 打印交易哈希
	fmt.Printf("Contract deployment transaction sent!\nTransaction Hash: %s\n", result)
	fmt.Println("Wait for confirmation on the blockchain...")

}

func TestCall(t *testing.T) {
	client, err := Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	// 调用 get() 函数
	_, err = callEthGet(client, contractAddress)
	if err != nil {
		log.Fatalf("Call failed: %v", err)
	}

	//fmt.Printf("Value from contract.get(): %d\n", value)
}

// 调用 eth_call 查询合约数据
func callEthGet(client *Client, contractAddr common.Address) (*big.Int, error) {
	// 1. 构造 data 字段：函数选择器 + 参数（get() 无参数）
	functionSelector := getFunctionSelector("retrieve()") //a3a6c447
	data := functionSelector                              // common.Hex2Bytes("2e64cec1")
	fmt.Println("call data: ", hex.EncodeToString(data))

	// 2. 构造 eth_call 请求参数
	type CallArgs struct {
		From string `json:"from"`
		To   string `json:"to"`
		Data string `json:"data"`
	}

	args := CallArgs{
		From: common.HexToAddress("0x30e938b0630c02f394d17925fdb5fb046f70d452").Hex(),
		To:   contractAddr.Hex(),
		Data: "0x34f42020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000046871736b00000000000000000000000000000000000000000000000000000000", //"0x" + hex.EncodeToString(data),
	}

	var result string
	err := client.Client().CallContext(context.Background(), &result, "eth_call", args, "latest")
	if err != nil {
		return nil, fmt.Errorf("failed to call eth_call: %v", err)
	}

	if result == "0x" || result == "" {
		return nil, fmt.Errorf("empty return value")
	}

	// 3. 解析返回值（32 字节 big.Int）
	decoded, err := hex.DecodeString(result[2:]) // 去掉 0x
	if err != nil {
		return nil, fmt.Errorf("failed to decode result: %v", err)
	}

	if len(decoded) < 32 {
		return nil, fmt.Errorf("return value too short: %d bytes", len(decoded))
	}

	// 取最后 32 字节（ABI 编码的 uint256）
	valueBytes := decoded[len(decoded)-32:]
	fmt.Println("采用sm3进行哈希:", hex.EncodeToString(valueBytes))
	value := new(big.Int).SetBytes(valueBytes)

	return value, nil
}

func getFunctionSelector(signature string) []byte {
	//hash := NewSm3State()
	//hash.Write([]byte(signature))
	//return hash.Sum(nil)[:4]
	hash2 := crypto.NewKeccakState()
	hash2.Write([]byte(signature))
	return hash2.Sum(nil)[:4]
}

func TestSelector(t *testing.T) {
	data := getFunctionSelector("sender()")
	t.Log("0x" + hex.EncodeToString(data))

	hash := NewSm3State()
	hash.Write([]byte("heelo"))
	ok := hash.Sum(nil)
	kk := sm3.Sm3Sum([]byte("heelo"))

	t.Log(hex.EncodeToString(ok), hex.EncodeToString(kk))
}

func TestTransafer(t *testing.T) {
	// 1. 连接到以太坊节点（可以是 Infura、Alchemy 或本地节点）
	client, err := Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	n, err := client.BlockNumber(context.Background())
	fmt.Println("start ", n, err)

	for i := 0; i < 20; i++ {
		Transfer(client)
	}

	n, err = client.BlockNumber(context.Background())
	fmt.Println("end ", n, err)
}

func Transfer(client *Client) {
	// 2. 加载私钥（⚠️ 仅用于测试！不要在生产中硬编码私钥）
	privateKey, _ := HexToSM2("39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d")

	// 3. 获取发送地址和 nonce
	fromAddress := PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// 4. 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// 6. 构造交易
	gasLimit := uint64(500000) // 根据合约复杂度调整

	// 创建一个 to 为空的交易，data 为合约字节码
	inner := GmTx{
		ChainID:  big.NewInt(1),
		Nonce:    nonce,
		To:       HexToAddress("0x89c2d721eebf8d27d1a89ecd336064a81bfaefcf"),
		Value:    big.NewInt(1).Mul(big.NewInt(1000000000), big.NewInt(1000000000)),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		//Data:     common.FromHex(contractBytecode),
		//R, S      *big.Int,
		PublicKey: sm2.Compress(&privateKey.PublicKey),
	}

	// 7. 签名交易
	h := Hash(inner)
	r, s, err := Sign(h[:], privateKey)
	if err != nil {
		panic(err)
	}
	inner.R = r
	inner.S = s
	signedTx := GmTransaction{
		inner: inner,
		time:  time.Now(),
	}

	// 8. 序列化为 RLP 编码
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v", err)
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)

	// 10. 调用 eth_sendRawTransaction
	var result interface{}
	err = client.Client().CallContext(context.Background(), &result, "eth_sendRawTransaction", "0x"+rawTxHex)
	if err != nil {
		log.Fatalf("Failed to send raw transaction: %v", err)
	}

	// 11. 打印交易哈希
	fmt.Printf("Transaction Hash: %s\n", result)
}
