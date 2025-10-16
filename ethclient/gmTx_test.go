package bs_eth

import (
	"context"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"hash"
	"log"
	"math/big"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	p1  = "39725efee3fb28614de3bacaffe4cc4bd8c436257e2c8bb887c4b5c4be45e76d"
	p2  = "ffa912c02092899c1b359b144c6f38611d4d36e1a548bc07feaad9f6d42892bd"
	p3  = "9fd16a2008d879795506f46bb5e7c400ebf3108dc8c235318577b98894e15609"
	p4  = "f6299cb9ba7c7168a00286d341f381a312d46964ad961d67ebe5426767357c51"
	p5  = "9f843f6b690fcb7431c3b38b4ae0cc4e834757d5a122e44ea8240daef6fea720"
	a1  = "0x30E938B0630C02F394d17925FDB5fB046f70d452"
	a2  = "0x0CEe4290dD5aF09E089390B423bA876213a6a94F"
	a3  = "0x89C2D721eeBF8d27D1A89eCd336064a81BFaefcf"
	a4  = "0x45d1C3D293AB329921f9d2762B463644d43dB6FE"
	a5  = "0x1955DbE40cA8053D4A3d43fd71eCc8633cF73735"
	Url = "http://127.0.0.1:8545"
)

/*
608060405234801561001057600080fd5b5060017fc670f864e1cbd31b04bbb7f207ac27ffee1b5925910425f98ee1448a774a05e960405160405180910390a27f1e69321bbdc0510e8b5f62e2a1bbbed6143ae12e782a1783a55e8fd5019f3b0f60405161006c90610175565b60405180910390a17fc0d505515bf144d2d18fa4d0308b7069fed35d610375ae95bca9cd33f25bde0a6040516100a1906101e1565b60405180910390a17f7c4352354ddbb23a20b7890913a0983763d5052c208f998378f159f5e6bd0e306040516100d69061024d565b60405180910390a17fec5dc4146de24512ee572c21b5a93986c75fdd0eea5c66655dca6e308a38550760405161010b906102b9565b60405180910390a16102d9565b600082825260208201905092915050565b7f68656c6c6f20776f726c64210000000000000000000000000000000000000000600082015250565b600061015f600c83610118565b915061016a82610129565b602082019050919050565b6000602082019050818103600083015261018e81610152565b9050919050565b7f7468697320697320612074657374210000000000000000000000000000000000600082015250565b60006101cb600f83610118565b91506101d682610195565b602082019050919050565b600060208201905081810360008301526101fa816101be565b9050919050565b7f676f6f64206c75636b20746f20796f7521000000000000000000000000000000600082015250565b6000610237601183610118565b915061024282610201565b602082019050919050565b600060208201905081810360008301526102668161022a565b9050919050565b7f6e69636520746f2073656520796f752100000000000000000000000000000000600082015250565b60006102a3601083610118565b91506102ae8261026d565b602082019050919050565b600060208201905081810360008301526102d281610296565b9050919050565b603f806102e76000396000f3fe6080604052600080fdfea26469706673582212209ba36049e149d523abc837bb02974b059268bbe9ab1e862fc137db2b4ee359c264736f6c634300081e0033
*/

func TestPriKey(t *testing.T) {
	testKey, _ := HexToSM2(p1)
	testAddr := PubkeyToAddress(testKey.PublicKey)
	pub := sm2.Compress(&testKey.PublicKey)

	fmt.Println(len(testAddr), testAddr)
	fmt.Println(len(pub), hex.EncodeToString(pub))

	byt := make([]byte, 100)
	copy(byt[:], pub)
	copy(byt[33:], testAddr.Bytes())

	fmt.Println(hex.EncodeToString(byt[:33]), hex.EncodeToString(byt[33:53]))
	fmt.Println(hex.EncodeToString(byt[:]))

	data := []byte("hello world")
	hash := sm3.Sm3Sum(data)
	sig, _ := testKey.Sign(rand.Reader, data, nil)
	fmt.Println("len(data),len(hash),len(sig),len(pub)", len(data), len(hash), len(sig), len(pub))

	toAddr := common.HexToAddress(a2)
	to := common.Address{}
	to.SetBytes(toAddr[:])
	fmt.Println("to", toAddr, to, &to)
}

func TestSendSm2Tx(t *testing.T) {
	client, err := Dial(Url)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	testKey, _ := HexToSM2(p1)
	testAddr := PubkeyToAddress(testKey.PublicKey)
	fmt.Println("账户地址:", testAddr)

	nonce, err := client.NonceAt(context.Background(), testAddr, nil)
	if err != nil {
		t.Fatal(err)
	}

	//toAddr, _ := hex.DecodeString(a2)
	toAddr := common.HexToAddress(a4)
	to := common.Address{}
	to.SetBytes(toAddr[:])
	//fmt.Printf("发送交易 from: %x to: %x, value: %x \n", testAddr, to, big.NewInt(1).Mul(big.NewInt(1000000000), big.NewInt(1000000000)))

	gmTx := GmTx{
		ChainID:  big.NewInt(1),
		Nonce:    nonce,
		To:       &to,
		Value:    big.NewInt(1).Mul(big.NewInt(1000000000), big.NewInt(1000000000)),
		Gas:      2200000,
		GasPrice: big.NewInt(params.InitialBaseFee),
		//Data:     bytecode,
		//R, S      *big.Int,
		PublicKey: FromSM2Pub(&testKey.PublicKey),
	}
	h := Hash(gmTx)
	r, s, err := Sign(h[:], testKey)
	if err != nil {
		t.Fatal(err)
	}
	gmTx.R = r
	gmTx.S = s
	tx := GmTransaction{
		inner: gmTx,
		time:  time.Now(),
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	encodeTx := hexutil.Encode(data)
	var txHash interface{}
	err = client.c.CallContext(context.Background(), &txHash, "eth_sendRawTransaction", encodeTx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("nonce", nonce)
	fmt.Println("签名哈希", h)
	fmt.Println("裸交易", encodeTx)
	fmt.Println("国密签名(r,s): ", r, s)
	fmt.Println("交易哈希: ", txHash)

	var result interface{}
	err = client.c.CallContext(context.Background(), &result, "eth_getTransactionByHash", txHash)
	fmt.Println("已上链，上链数据: ", result)
	fmt.Printf("交易后账户余额：")
	balanceOf(toAddr)
}

func balanceOf(addr common.Address) {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	bal, _ := client.BalanceAt(context.Background(), addr, nil)
	fmt.Printf("address: %v balance: %v \n", addr, bal)
}

func HexToSM2(hexkey string) (*sm2.PrivateKey, error) {
	b, err := hex.DecodeString(hexkey)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in private key", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for private key")
	}
	return ToSM2(b)
}

func ToSM2(d []byte) (*sm2.PrivateKey, error) {
	return toSM2(d, true)
}

func toSM2(d []byte, strict bool) (*sm2.PrivateKey, error) {
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = sm2.P256Sm2()
	sm2p256v1N, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	if strict && 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	if priv.D.Cmp(sm2p256v1N) >= 0 {
		return nil, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}

type Sm3State interface {
	hash.Hash
	//Read([]byte) (int, error)
}

func NewSm3State() Sm3State {
	return sm3.New().(Sm3State)
}

func SM3(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := NewSm3State()
	for _, b := range data {
		d.Write(b)
	}
	//d.Read(b)

	b = d.Sum(nil)
	return b
}

func FromSM2Pub(pub *sm2.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(sm2.P256Sm2(), pub.X, pub.Y)
}

func PubkeyToAddress(p sm2.PublicKey) common.Address {
	pubBytes := FromSM2Pub(&p)
	return common.BytesToAddress(SM3(pubBytes[1:])[12:])
}

type GmTx struct {
	ChainID   *big.Int
	Nonce     uint64
	GasPrice  *big.Int
	Gas       uint64
	To        *common.Address `rlp:"nil"`
	Value     *big.Int
	Data      []byte
	R, S      *big.Int
	PublicKey []byte
}

func Hash(tx GmTx) common.Hash {
	return rlpHash([]interface{}{
		tx.Nonce,
		tx.GasPrice,
		tx.Gas,
		tx.To,
		tx.Value,
		tx.Data,
		tx.ChainID,
		uint(0),
		uint(0),
	})
}

var hasherPool = sync.Pool{
	New: func() interface{} { return sm3.New() },
}

//func rlpHash(x interface{}) (h common.Hash) {
//	sha := hasherPool.Get().(Sm3State)
//	defer hasherPool.Put(sha)
//	sha.Reset()
//	rlp.Encode(sha, x)
//	//sha.Read(h[:])
//	result := sha.Sum(nil)
//	h = common.BytesToHash(result)
//
//	var buf bytes.Buffer
//	rlp.Encode(&buf, x)
//	encoded := buf.Bytes()
//	eh := common.Bytes2Hex(encoded)
//	fmt.Println("未哈希交易:", eh, len(eh))
//
//	return h
//}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

const DigestLength int = 32

func Sign(digestHash []byte, prv *sm2.PrivateKey) (r, s *big.Int, err error) {
	if len(digestHash) != DigestLength {
		return big.NewInt(0), big.NewInt(0), fmt.Errorf("hash is required to be exactly %d bytes (%d)", DigestLength, len(digestHash))
	}

	return sm2.Sm2Sign(prv, digestHash, nil, rand.Reader)
	//return prv.Sign(rand.Reader, digestHash, nil)
}

type GmTransaction struct {
	inner GmTx      // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

func (tx *GmTransaction) MarshalBinary() ([]byte, error) {
	return rlp.EncodeToBytes(tx.inner)
}

type EcdsaTx struct {
	ChainID  *big.Int
	Nonce    uint64
	GasPrice *big.Int
	Gas      uint64
	To       *common.Address `rlp:"nil"`
	Value    *big.Int
	Data     []byte
	R, S, V  *big.Int
}

func etxHash(tx EcdsaTx) common.Hash {
	return rlpHash([]interface{}{
		tx.Nonce,
		tx.GasPrice,
		tx.Gas,
		tx.To,
		tx.Value,
		tx.Data,
		tx.ChainID, uint(0), uint(0), uint(0),
	})
}

type EcdsaTransaction struct {
	inner EcdsaTx   // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

func (tx *EcdsaTransaction) MarshalBinary() ([]byte, error) {
	return rlp.EncodeToBytes(tx.inner)
}

func UnCompressBytesToPub(e []byte) *sm2.PublicKey {
	key := &sm2.PublicKey{}
	key.X = new(big.Int).SetBytes(e[1:33])
	key.Y = new(big.Int).SetBytes(e[33:])
	key.Curve = sm2.P256Sm2()
	return key
}

func PubToUnCompressBytes(pub *sm2.PublicKey) []byte {
	xBytes := bigIntTo32Bytes(pub.X)
	yBytes := bigIntTo32Bytes(pub.Y)
	xl := len(xBytes)
	yl := len(yBytes)

	raw := make([]byte, 1+KeyBytes*2)
	raw[0] = UnCompress
	if xl > KeyBytes {
		copy(raw[1:1+KeyBytes], xBytes[xl-KeyBytes:])
	} else if xl < KeyBytes {
		copy(raw[1+(KeyBytes-xl):1+KeyBytes], xBytes)
	} else {
		copy(raw[1:1+KeyBytes], xBytes)
	}

	if yl > KeyBytes {
		copy(raw[1+KeyBytes:], yBytes[yl-KeyBytes:])
	} else if yl < KeyBytes {
		copy(raw[1+KeyBytes+(KeyBytes-yl):], yBytes)
	} else {
		copy(raw[1+KeyBytes:], yBytes)
	}
	return raw
}

const (
	BitSize    = 256
	KeyBytes   = (BitSize + 7) / 8
	UnCompress = 0x04
)

func bigIntTo32Bytes(bn *big.Int) []byte {
	byteArr := bn.Bytes()
	byteArrLen := len(byteArr)
	if byteArrLen == KeyBytes {
		return byteArr
	}
	byteArr = append(make([]byte, KeyBytes-byteArrLen), byteArr...)
	return byteArr
}

func TestSm2(t *testing.T) {
	hash := common.HexToHash("2daef60e7a0b8f5e024c81cd2ab3109f2b4f155cf83adeb2ae5532f74a157fdf")
	fmt.Println("hash", hash)

	pri, _ := HexToSM2("9fd16a2008d879795506f46bb5e7c400ebf3108dc8c235318577b98894e15609")
	pub := FromSM2Pub(&pri.PublicKey)
	R, S, _ := Sign(hash[:], pri)
	pb := UnCompressBytesToPub(pub)

	//R, _ := big.NewInt(0).SetString("f5aabd6d515dc7fb54d36172cd172769f15c73685f204f00c9ac88c14cf8f036", 16)
	//S, _ := big.NewInt(0).SetString("25be6382bd22e13bc231565f44418e0925a7a9f81370270fe62077fc7802f173", 16)
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, 64)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	//sig := common.Hex2Bytes("304502206a446b4edc15089ec6eb5b06cf500c19bad961484aa6fb65575066e152fb6361022100ce7da6b250079c5bd3f2f5b640901647160220cad7c2964091c60814489251ae")
	//fmt.Println("sig", common.Bytes2Hex(sig))

	//pub := common.Hex2Bytes("043ce455fd10aff3f6ebf82a10270829b93b5e65aeebac604efcaba66e4c8a6091d59f747406d50b3bebf74cef1bc5bd1df28a0f181c9aa65ef612aa531ceec7c3")
	fmt.Println("pub", common.Bytes2Hex(pub))

	addr := PubkeyToAddress(pri.PublicKey)
	fmt.Println("addr:", addr)

	result := sm2.Sm2Verify(&pri.PublicKey, hash[:], []byte("1234567812345678"), R, S)
	fmt.Println(result)
	result = sm2.Sm2Verify(pb, hash[:], []byte("1234567812345678"), R, S)
	fmt.Println(result)
	result = pri.Verify(hash[:], sig)
	fmt.Println(result)
}
