package blockchain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/gmsm"
	"github.com/ethereum/go-ethereum/gmsm/sm2"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tjfoc/gmsm/sm3"
)

const (
	LegacyTxType = iota
	AccessListTxType
	DynamicFeeTxType
)

type TxData interface {
	copy() TxData
	chainID() *big.Int
	accessList() AccessList
	data() []byte
	gas() uint64
	gasPrice() *big.Int
	gasTipCap() *big.Int
	gasFeeCap() *big.Int
	value() *big.Int
	nonce() uint64
	to() *common.Address
	txType() byte
	rawSignatureValues() (r, s *big.Int)
	from() *common.Address
	setSignatureValues(chainID, r, s *big.Int)
	setPublicKey(pubKey []byte)
}
type AccessList []AccessTuple
type AccessTuple struct {
	Address     common.Address `json:"address"     gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
}
type sigCache struct {
	signer Signer
	from   common.Address
}
type Signer interface {
	// Sender returns the sender address of the transaction.
	Sender(tx *Transaction) (common.Address, error)

	// SignatureValues returns the raw R, S, V values corresponding to the
	// given signature.
	SignatureValues(tx *Transaction, sig []byte) (r, s *big.Int, err error)
	ChainID() *big.Int

	// Hash returns 'signature hash', i.e. the transaction hash that is signed by the
	// private key. This hash does not uniquely identify the transaction.
	Hash(tx *Transaction) common.Hash

	// Equal returns true if the given signer is the same as the receiver.
	Equal(Signer) bool
}
type txJSON struct {
	Type hexutil.Uint64 `json:"type"`

	// Common transaction fields:
	Nonce                *hexutil.Uint64 `json:"nonce"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	Value                *hexutil.Big    `json:"value"`
	Data                 *hexutil.Bytes  `json:"input"`
	R                    *hexutil.Big    `json:"r"`
	S                    *hexutil.Big    `json:"s"`
	To                   *common.Address `json:"to"`
	//PublicKey            *hexutil.Bytes  `json:"publicKey"`
	From *common.Address `json:"from"`

	// Access list transaction fields:
	ChainID    *hexutil.Big `json:"chainId,omitempty"`
	AccessList *AccessList  `json:"accessList,omitempty"`

	// Only used for encoding:
	Hash common.Hash `json:"hash"`
}

type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Pointer[common.Hash]
	size atomic.Uint64
	from atomic.Pointer[sigCache]
}

func (tx *Transaction) setDecoded(inner TxData, size int) {
	tx.inner = inner
	tx.time = time.Now()
	if size > 0 {
		tx.size.Store(uint64(common.StorageSize(size)))
	}
}
func (tx *Transaction) UnmarshalJSON(input []byte) error {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}

	// Decode / verify fields according to transaction type.
	var inner TxData
	switch dec.Type {
	case LegacyTxType:
		var itx LegacyTx
		inner = &itx
		if dec.From != nil {
			itx.From = dec.From
		}
		if dec.To != nil {
			itx.To = dec.To
		}
		if dec.Nonce == nil {
			return errors.New("missing required field 'nonce' in transaction")
		}
		itx.Nonce = uint64(*dec.Nonce)
		if dec.GasPrice == nil {
			return errors.New("missing required field 'gasPrice' in transaction")
		}
		itx.GasPrice = (*big.Int)(dec.GasPrice)
		if dec.Gas == nil {
			return errors.New("missing required field 'gas' in transaction")
		}
		itx.Gas = uint64(*dec.Gas)
		if dec.Value == nil {
			return errors.New("missing required field 'value' in transaction")
		}
		itx.Value = (*big.Int)(dec.Value)
		if dec.Data == nil {
			return errors.New("missing required field 'input' in transaction")
		}
		itx.Data = *dec.Data
		if dec.R == nil {
			return errors.New("missing required field 'r' in transaction")
		}
		itx.R = (*big.Int)(dec.R)
		if dec.S == nil {
			return errors.New("missing required field 's' in transaction")
		}
		itx.S = (*big.Int)(dec.S)

	case AccessListTxType:
		var itx AccessListTx
		inner = &itx
		// Access list is optional for now.
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.From != nil {
			itx.From = dec.From
		}
		if dec.To != nil {
			itx.To = dec.To
		}
		if dec.Nonce == nil {
			return errors.New("missing required field 'nonce' in transaction")
		}
		itx.Nonce = uint64(*dec.Nonce)
		if dec.GasPrice == nil {
			return errors.New("missing required field 'gasPrice' in transaction")
		}
		itx.GasPrice = (*big.Int)(dec.GasPrice)
		if dec.Gas == nil {
			return errors.New("missing required field 'gas' in transaction")
		}
		itx.Gas = uint64(*dec.Gas)
		if dec.Value == nil {
			return errors.New("missing required field 'value' in transaction")
		}
		itx.Value = (*big.Int)(dec.Value)
		if dec.Data == nil {
			return errors.New("missing required field 'input' in transaction")
		}
		itx.Data = *dec.Data
		if dec.R == nil {
			return errors.New("missing required field 'r' in transaction")
		}
		itx.R = (*big.Int)(dec.R)
		if dec.S == nil {
			return errors.New("missing required field 's' in transaction")
		}
		itx.S = (*big.Int)(dec.S)

	case DynamicFeeTxType:
		var itx DynamicFeeTx
		inner = &itx
		// Access list is optional for now.
		if dec.AccessList != nil {
			itx.AccessList = *dec.AccessList
		}
		if dec.ChainID == nil {
			return errors.New("missing required field 'chainId' in transaction")
		}
		itx.ChainID = (*big.Int)(dec.ChainID)
		if dec.From != nil {
			itx.From = dec.From
		}
		if dec.To != nil {
			itx.To = dec.To
		}
		if dec.Nonce == nil {
			return errors.New("missing required field 'nonce' in transaction")
		}
		itx.Nonce = uint64(*dec.Nonce)
		if dec.MaxPriorityFeePerGas == nil {
			return errors.New("missing required field 'maxPriorityFeePerGas' for txdata")
		}
		itx.GasTipCap = (*big.Int)(dec.MaxPriorityFeePerGas)
		if dec.MaxFeePerGas == nil {
			return errors.New("missing required field 'maxFeePerGas' for txdata")
		}
		itx.GasFeeCap = (*big.Int)(dec.MaxFeePerGas)
		if dec.Gas == nil {
			return errors.New("missing required field 'gas' for txdata")
		}
		itx.Gas = uint64(*dec.Gas)
		if dec.Value == nil {
			return errors.New("missing required field 'value' in transaction")
		}
		itx.Value = (*big.Int)(dec.Value)
		if dec.Data == nil {
			return errors.New("missing required field 'input' in transaction")
		}
		itx.Data = *dec.Data
		if dec.R == nil {
			return errors.New("missing required field 'r' in transaction")
		}
		itx.R = (*big.Int)(dec.R)
		if dec.S == nil {
			return errors.New("missing required field 's' in transaction")
		}
		itx.S = (*big.Int)(dec.S)

	default:
		return errors.New("transaction type not supported")
	}

	// Now set the inner transaction.
	tx.setDecoded(inner, 0)
	tx.hash.Store(&dec.Hash)

	// TODO: check hash here?
	return nil
}
func (tx *Transaction) MarshalBinary() ([]byte, error) {
	if tx.Type() == LegacyTxType {
		return rlp.EncodeToBytes(tx.inner)
	}
	var buf bytes.Buffer
	err := tx.encodeTyped(&buf)
	return buf.Bytes(), err
}
func (tx *Transaction) Type() uint8 {
	return tx.inner.txType()
}
func (tx *Transaction) encodeTyped(w *bytes.Buffer) error {
	w.WriteByte(tx.Type())
	return rlp.Encode(w, tx.inner)
}
func (tx *Transaction) RawSignatureValues() (r, s *big.Int) {
	return tx.inner.rawSignatureValues()
}
func (tx *Transaction) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return *hash
	}
	panic("交易哈希不能生产")

	//var h common.Hash
	//if tx.Type() == LegacyTxType {
	//	h = rlpHash(tx.inner)
	//} else {
	//	h = prefixedRlpHash(tx.Type(), tx.inner)
	//}
	//
	//log.Println("交易哈希不能保证正确",h)
	//return h
}
func (tx *Transaction) From() *common.Address {
	return copyAddressPtr(tx.inner.from())
}
func (tx *Transaction) ChainId() *big.Int {
	return tx.inner.chainID()
}
func (tx *Transaction) Data() []byte           { return tx.inner.data() }
func (tx *Transaction) AccessList() AccessList { return tx.inner.accessList() }
func (tx *Transaction) Gas() uint64            { return tx.inner.gas() }
func (tx *Transaction) GasPrice() *big.Int     { return new(big.Int).Set(tx.inner.gasPrice()) }
func (tx *Transaction) GasTipCap() *big.Int    { return new(big.Int).Set(tx.inner.gasTipCap()) }
func (tx *Transaction) GasFeeCap() *big.Int    { return new(big.Int).Set(tx.inner.gasFeeCap()) }
func (tx *Transaction) Value() *big.Int        { return new(big.Int).Set(tx.inner.value()) }
func (tx *Transaction) Nonce() uint64          { return tx.inner.nonce() }
func (tx *Transaction) To() *common.Address {
	return copyAddressPtr(tx.inner.to())
}
func (tx *Transaction) WithSignature(chainID *big.Int, sig []byte) (*Transaction, error) {
	r, s := decodeSignature(sig)
	cpy := tx.inner.copy()
	cpy.setSignatureValues(chainID, r, s)
	return &Transaction{inner: cpy, time: tx.time}, nil
}

func rlpHash(x interface{}) (h common.Hash) {
	// hasherPool holds LegacyKeccak256 hashers for rlpHash.
	var hasherPool = sync.Pool{
		New: func() interface{} { return sm3.New() },
	}

	sha := hasherPool.Get().(gmsm.Sm3State)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	//sha.Read(h[:])
	result := sha.Sum(nil)
	h = common.BytesToHash(result)
	return h
}
func prefixedRlpHash(prefix byte, x interface{}) (h common.Hash) {
	// hasherPool holds LegacyKeccak256 hashers for rlpHash.
	var hasherPool = sync.Pool{
		New: func() interface{} { return sm3.New() },
	}

	sha := hasherPool.Get().(gmsm.Sm3State)
	defer hasherPool.Put(sha)
	sha.Reset()
	sha.Write([]byte{prefix})
	rlp.Encode(sha, x)
	//sha.Read(h[:])
	result := sha.Sum(nil)
	h = common.BytesToHash(result)
	return h
}
func copyAddressPtr(a *common.Address) *common.Address {
	if a == nil {
		return nil
	}
	cpy := *a
	return &cpy
}
func decodeSignature(sig []byte) (r, s *big.Int) {
	if len(sig) != gmsm.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), gmsm.SignatureLength))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	return r, s
}
func SignNewTx(prv *sm2.PrivateKey, chainID *big.Int, txdata TxData) (*Transaction, error) {
	txdata.setPublicKey(gmsm.CompressPubkey(&prv.PublicKey))
	tx := new(Transaction)
	tx.setDecoded(txdata.copy(), 0)
	h := HashTx(tx, chainID)
	sig, err := gmsm.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(chainID, sig)
}
func HashTx(tx *Transaction, chainId *big.Int) common.Hash {
	switch tx.Type() {
	case LegacyTxType:
		return rlpHash([]interface{}{
			tx.Nonce(),
			tx.GasPrice(),
			tx.Gas(),
			tx.To(),
			tx.Value(),
			tx.Data(),
			chainId, uint(0), uint(0),
		})
	case AccessListTxType:
		return prefixedRlpHash(
			tx.Type(),
			[]interface{}{
				chainId,
				tx.Nonce(),
				tx.GasPrice(),
				tx.Gas(),
				tx.To(),
				tx.Value(),
				tx.Data(),
				tx.AccessList(),
			})
	case DynamicFeeTxType:
		return prefixedRlpHash(
			tx.Type(),
			[]interface{}{
				chainId,
				tx.Nonce(),
				tx.GasTipCap(),
				tx.GasFeeCap(),
				tx.Gas(),
				tx.To(),
				tx.Value(),
				tx.Data(),
				tx.AccessList(),
			})
	default:
		// This _should_ not happen, but in case someone sends in a bad
		// json struct via RPC, it's probably more prudent to return an
		// empty hash instead of killing the node with a panic
		//panic("Unsupported transaction type: %d", tx.typ)
		return common.Hash{}
	}
}

type LegacyTx struct {
	From      *common.Address `rlp:"nil"`
	ChainID   *big.Int
	Nonce     uint64          // nonce of sender account
	GasPrice  *big.Int        // wei per gas
	Gas       uint64          // gas limit
	To        *common.Address `rlp:"nil"` // nil means contract creation
	Value     *big.Int        // wei amount
	Data      []byte          // contract invocation input data
	R, S      *big.Int        // signature values
	PublicKey []byte
}

func (tx *LegacyTx) copy() TxData {
	cpy := &LegacyTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		//V:        new(big.Int),
		R:         new(big.Int),
		S:         new(big.Int),
		ChainID:   tx.ChainID,
		PublicKey: common.CopyBytes(tx.PublicKey),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	//if tx.V != nil {
	//	cpy.V.Set(tx.V)
	//}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}
func (tx *LegacyTx) txType() byte { return LegacyTxType }
func (tx *LegacyTx) chainID() *big.Int {
	if tx.ChainID == nil {
		return new(big.Int) //params.MainnetChainConfig.ChainID
	}
	return tx.ChainID
}
func (tx *LegacyTx) accessList() AccessList { return nil }
func (tx *LegacyTx) data() []byte           { return tx.Data }
func (tx *LegacyTx) gas() uint64            { return tx.Gas }
func (tx *LegacyTx) gasPrice() *big.Int     { return tx.GasPrice }
func (tx *LegacyTx) gasTipCap() *big.Int    { return tx.GasPrice }
func (tx *LegacyTx) gasFeeCap() *big.Int    { return tx.GasPrice }
func (tx *LegacyTx) value() *big.Int        { return tx.Value }
func (tx *LegacyTx) nonce() uint64          { return tx.Nonce }
func (tx *LegacyTx) to() *common.Address    { return tx.To }
func (tx *LegacyTx) rawSignatureValues() (r, s *big.Int) {
	return tx.R, tx.S
}
func (tx *LegacyTx) from() *common.Address { return tx.From }
func (tx *LegacyTx) setPublicKey(pubKey []byte) {
	tx.PublicKey = pubKey
}
func (tx *LegacyTx) setSignatureValues(chainID, r, s *big.Int) {
	tx.ChainID, tx.R, tx.S = chainID, r, s
}

type AccessListTx struct {
	From       *common.Address `rlp:"nil"`
	ChainID    *big.Int        // destination chain ID
	Nonce      uint64          // nonce of sender account
	GasPrice   *big.Int        // wei per gas
	Gas        uint64          // gas limit
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int        // wei amount
	Data       []byte          // contract invocation input data
	AccessList AccessList      // EIP-2930 access list
	R, S       *big.Int        // signature values
	PublicKey  []byte
}

func (tx *AccessListTx) copy() TxData {
	cpy := &AccessListTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		AccessList: make(AccessList, len(tx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasPrice:   new(big.Int),
		//V:          new(big.Int),
		R:         new(big.Int),
		S:         new(big.Int),
		PublicKey: common.CopyBytes(tx.PublicKey),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	//if tx.V != nil {
	//	cpy.V.Set(tx.V)
	//}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}
func (tx *AccessListTx) txType() byte           { return AccessListTxType }
func (tx *AccessListTx) chainID() *big.Int      { return tx.ChainID }
func (tx *AccessListTx) accessList() AccessList { return tx.AccessList }
func (tx *AccessListTx) data() []byte           { return tx.Data }
func (tx *AccessListTx) gas() uint64            { return tx.Gas }
func (tx *AccessListTx) gasPrice() *big.Int     { return tx.GasPrice }
func (tx *AccessListTx) gasTipCap() *big.Int    { return tx.GasPrice }
func (tx *AccessListTx) gasFeeCap() *big.Int    { return tx.GasPrice }
func (tx *AccessListTx) value() *big.Int        { return tx.Value }
func (tx *AccessListTx) nonce() uint64          { return tx.Nonce }
func (tx *AccessListTx) to() *common.Address    { return tx.To }
func (tx *AccessListTx) rawSignatureValues() (r, s *big.Int) {
	return tx.R, tx.S
}
func (tx *AccessListTx) setSignatureValues(chainID, r, s *big.Int) {
	tx.ChainID, tx.R, tx.S = chainID, r, s
}
func (tx *AccessListTx) from() *common.Address { return tx.From }
func (tx *AccessListTx) setPublicKey(pubKey []byte) {
	tx.PublicKey = pubKey
}

type DynamicFeeTx struct {
	From       *common.Address `rlp:"nil"`
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       []byte
	AccessList AccessList
	R          *big.Int `json:"r" gencodec:"required"`
	S          *big.Int `json:"s" gencodec:"required"`
	PublicKey  []byte
}

func (tx *DynamicFeeTx) copy() TxData {
	cpy := &DynamicFeeTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		AccessList: make(AccessList, len(tx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasTipCap:  new(big.Int),
		GasFeeCap:  new(big.Int),
		//V:          new(big.Int),
		R:         new(big.Int),
		S:         new(big.Int),
		PublicKey: common.CopyBytes(tx.PublicKey),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.GasTipCap)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	//if tx.V != nil {
	//	cpy.V.Set(tx.V)
	//}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}
func (tx *DynamicFeeTx) txType() byte           { return DynamicFeeTxType }
func (tx *DynamicFeeTx) chainID() *big.Int      { return tx.ChainID }
func (tx *DynamicFeeTx) accessList() AccessList { return tx.AccessList }
func (tx *DynamicFeeTx) data() []byte           { return tx.Data }
func (tx *DynamicFeeTx) gas() uint64            { return tx.Gas }
func (tx *DynamicFeeTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *DynamicFeeTx) gasTipCap() *big.Int    { return tx.GasTipCap }
func (tx *DynamicFeeTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *DynamicFeeTx) value() *big.Int        { return tx.Value }
func (tx *DynamicFeeTx) nonce() uint64          { return tx.Nonce }
func (tx *DynamicFeeTx) to() *common.Address    { return tx.To }
func (tx *DynamicFeeTx) rawSignatureValues() (r, s *big.Int) {
	return tx.R, tx.S
}
func (tx *DynamicFeeTx) from() *common.Address { return tx.From }
func (tx *DynamicFeeTx) setPublicKey(pubKey []byte) {
	tx.PublicKey = pubKey
}
func (tx *DynamicFeeTx) setSignatureValues(chainID, r, s *big.Int) {
	tx.ChainID, tx.R, tx.S = chainID, r, s
}
func (tx *DynamicFeeTx) getHash() *big.Int { return tx.ChainID }
