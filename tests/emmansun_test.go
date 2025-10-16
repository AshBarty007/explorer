package tests

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/emmansun/gmsm/sm2"
	"log"
	"math/big"
	"os"
	"testing"
)

func TestGo(t *testing.T) {
	//default_uid := []byte("1234567812345678")
	toSign := []byte("hello")
	// real private key should be from secret storage
	privKey, _ := hex.DecodeString("9fd16a2008d879795506f46bb5e7c400ebf3108dc8c235318577b98894e15609")
	testkey, err := sm2.NewPrivateKey(privKey)
	if err != nil {
		log.Fatalf("fail to new private key %v", err)
	}

	// force SM2 sign standard and use default UID
	sig, err := testkey.Sign(rand.Reader, toSign, sm2.DefaultSM2SignerOpts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from sign: %s\n", err)
		return
	}
	// Since sign is a randomized function, signature will be
	// different each time.
	fmt.Printf("%x\n", sig)
}

func TestJs(t *testing.T) {
	default_uid := []byte("1234567812345678")
	toSign := []byte("hello")
	// real private key should be from secret storage
	privKey, _ := hex.DecodeString("9fd16a2008d879795506f46bb5e7c400ebf3108dc8c235318577b98894e15609")
	testkey, err := sm2.NewPrivateKey(privKey)
	if err != nil {
		log.Fatalf("fail to new private key %v", err)
	}

	signature, _ := hex.DecodeString("eb7f586f093e8e254c03902cb2de248d5863c4049a722f61fecaf14f7ac9a5046f3ed9c4a4d4ebb772e5ac621f80377749b53c73a3407645432a4ad0cdd6a8c4")
	ok := sm2.VerifyASN1WithSM2(&testkey.PublicKey, default_uid, toSign, signature)
	fmt.Printf("%v\n", ok) //false

	r, _ := big.NewInt(0).SetString("eb7f586f093e8e254c03902cb2de248d5863c4049a722f61fecaf14f7ac9a504", 16)
	s, _ := big.NewInt(0).SetString("6f3ed9c4a4d4ebb772e5ac621f80377749b53c73a3407645432a4ad0cdd6a8c4", 16)
	ok = sm2.VerifyWithSM2(&testkey.PublicKey, default_uid, toSign, r, s)
	fmt.Printf("%v\n", ok) //false

	//ok = sm2.Verify()
	//fmt.Printf("%v\n", ok) //false
}
