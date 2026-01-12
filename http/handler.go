package bshttp

import (
	db "blockchain_services/postgres"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)


//////////////////////////////////////////////////////////////////////////////Block/////////////////////////////////////////////////////////////////////////

// GetBlockByNumberHandler 查询区块
func GetBlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	if number == "" {
		http.Error(w, "缺少参数: number", http.StatusBadRequest)
		return
	}

	block, err := db.GetBlockByNumber(number)
	if err != nil {
		log.Printf("GetBlockByNumber查询失败: %v", err)
		http.Error(w, "查询区块失败", http.StatusNotFound)
		return
	}

	jsonResponse(w, block, nil)
}

// GetBlocksHandler 查询区块列表
func GetBlocksHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	blocks, err := db.GetBlocks(start, end)
	if err != nil {
		log.Printf("GetBlocks查询失败: %v", err)
		http.Error(w, "查询区块列表失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, blocks, nil)
}

// GetBlockCountHandler 查询区块总数
func GetBlockCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetBlockCount()
	if err != nil {
		log.Printf("GetBlockCount查询失败: %v", err)
		http.Error(w, "查询区块总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

//////////////////////////////////////////////////////////////////////////////Transaction/////////////////////////////////////////////////////////////////////////

// GetTransactionByHashHandler 查询交易
func GetTransactionByHashHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "缺少参数: hash", http.StatusBadRequest)
		return
	}

	transaction, err := db.GetTransactionByHash(hash)
	if err != nil {
		log.Printf("GetTransactionByHash查询失败: %v", err)
		http.Error(w, "查询交易失败", http.StatusNotFound)
		return
	}

	jsonResponse(w, transaction, nil)
}

// GetTransactionsHandler 查询交易列表
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	transactions, err := db.GetTransactions(start, end)
	if err != nil {
		log.Printf("GetTransactions查询失败: %v", err)
		http.Error(w, "查询交易列表失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, transactions, nil)
}

// GetTransactionCountHandler 查询交易总数
func GetTransactionCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetTransactionCount()
	if err != nil {
		log.Printf("GetTransactionCount查询失败: %v", err)
		http.Error(w, "查询交易总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

// GetTransactionsByAddressHandler 查询该地址的交易列表
func GetTransactionsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "invalid 'address' format: must be a valid 0x-prefixed hex address (42 chars)", http.StatusBadRequest)
		return
	}

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, "invalid 'start': must be a non-negative integer", http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, "invalid 'end': must be a non-negative integer", http.StatusBadRequest)
		return
	}

	transactions, err := db.GetTransactionsByAddress(address, start, end)
	if err != nil {
		log.Printf("GetTransactionsByAddress查询失败: %v", err)
		http.Error(w, "查询交易失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, transactions, nil)
}

// GetTransactionCountByAddressHandler 查询该地址的交易总数
func GetTransactionCountByAddressHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}

	count, err := db.GetTransactionCountByAddress(address)
	if err != nil {
		log.Printf("GetTransactionCountByAddress查询失败: %v", err)
		http.Error(w, "查询地址交易总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

//////////////////////////////////////////////////////////////////////////////Receipt/////////////////////////////////////////////////////////////////////////

// GetReceiptByHashHandler 查询收据
func GetReceiptByHashHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "缺少参数: hash", http.StatusBadRequest)
		return
	}

	receipt, err := db.GetReceiptByHash(hash)
	if err != nil {
		log.Printf("GetReceiptByHash查询失败: %v", err)
		http.Error(w, "查询收据失败", http.StatusNotFound)
		return
	}

	jsonResponse(w, receipt, nil)
}

// GetReceiptsHandler 查询收据列表
func GetReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	receipts, err := db.GetReceipts(start, end)
	if err != nil {
		log.Printf("GetReceipts查询失败: %v", err)
		http.Error(w, "查询收据列表失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, receipts, nil)
}

// GetReceiptCountHandler 查询收据总数
func GetReceiptCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetReceiptCount()
	if err != nil {
		log.Printf("GetReceiptCount查询失败: %v", err)
		http.Error(w, "查询收据总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

// GetReceiptsByAddressHandler 查询该地址的收据列表
func GetReceiptsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	receipts, err := db.GetReceiptsByAddress(address, start, end)
	if err != nil {
		log.Printf("GetReceiptsByAddress查询失败: %v", err)
		http.Error(w, "查询收据失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, receipts, nil)
}

// GetReceiptCountByAddressHandler 查询该地址的收据总数
func GetReceiptCountByAddressHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}

	count, err := db.GetReceiptCountByAddress(address)
	if err != nil {
		log.Printf("GetReceiptCountByAddress查询失败: %v", err)
		http.Error(w, "查询地址收据总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

// GetReceiptsByLastHandler 查询最近7天收据
func GetReceiptsByLastHandler(w http.ResponseWriter, r *http.Request) {
	receipts := db.GetReceiptsByLast()
	jsonResponse(w, receipts, nil)
}

//////////////////////////////////////////////////////////////////////////////Account/////////////////////////////////////////////////////////////////////////

// GetAccountHandler 查询账户
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}

	account, err := db.GetAccount(address)
	if err != nil {
		log.Printf("GetAccount查询失败: %v", err)
		http.Error(w, "查询账户失败", http.StatusNotFound)
		return
	}

	jsonResponse(w, account, nil)
}

// GetAccountsHandler 批量查询账户
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	accounts, err := db.GetAccounts(start, end)
	if err != nil {
		log.Printf("GetAccounts查询失败: %v", err)
		http.Error(w, "查询账户列表失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, accounts, nil)
}

// GetAccountCountHandler 查询账户总数
func GetAccountCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetAccountCount()
	if err != nil {
		log.Printf("GetAccountCount查询失败: %v", err)
		http.Error(w, "查询账户总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

//////////////////////////////////////////////////////////////////////////////Token/////////////////////////////////////////////////////////////////////////

// GetTokenHandler 查询Token
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}

	token, err := db.GetToken(address)
	if err != nil {
		log.Printf("GetToken查询失败: %v", err)
		http.Error(w, "查询Token失败", http.StatusNotFound)
		return
	}

	jsonResponse(w, token, nil)
}

// GetTokensHandlerFromDB 批量查询Token（使用数据库查询）
func GetTokensHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	tokens, err := db.GetTokens(start, end)
	if err != nil {
		log.Printf("GetTokens查询失败: %v", err)
		http.Error(w, "查询Token列表失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, tokens, nil)
}

// GetTokenCountHandler 查询Token总数
func GetTokenCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetTokenCount()
	if err != nil {
		log.Printf("GetTokenCount查询失败: %v", err)
		http.Error(w, "查询Token数量失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

//////////////////////////////////////////////////////////////////////////////NFT/////////////////////////////////////////////////////////////////////////

// GetNFTHandler 查询NFT
func GetNFTHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}

	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}

	nft, err := db.GetNFT(address)
	if err != nil {
		log.Printf("GetNFT查询失败: %v", err)
		http.Error(w, "查询NFT失败", http.StatusNotFound)
		return
	}

	jsonResponse(w, nft, nil)
}

// GetNFTsHandler 批量查询NFT
func GetNFTsHandler(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := parseNonNegativeInt(startStr, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的start参数: %v", err), http.StatusBadRequest)
		return
	}

	end, err := parseNonNegativeInt(endStr, 10)
	if err != nil {
		http.Error(w, fmt.Sprintf("无效的end参数: %v", err), http.StatusBadRequest)
		return
	}

	nfts, err := db.GetNFTs(start, end)
	if err != nil {
		log.Printf("GetNFTs查询失败: %v", err)
		http.Error(w, "查询NFT列表失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, nfts, nil)
}

// GetNFTCountHandler 查询NFT总数
func GetNFTCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db.GetNFTCount()
	if err != nil {
		log.Printf("GetNFTCount查询失败: %v", err)
		http.Error(w, "查询NFT总数失败", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, count, nil)
}

//////////////////////////////////////////////////////////////////////////////Helper/////////////////////////////////////////////////////////////////////////

// parseNonNegativeInt 解析非负整数
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
