package bshttp

import (
	"blockchain_services/config"
	bs_eth "blockchain_services/ethclient"
	bsdb "blockchain_services/postgres"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"net/http"
	"strconv"
)

func StartHttp() {
	// 注册路由处理函数
	http.HandleFunc("/GetBlockByNumber", GetBlockByNumberHandler)
	http.HandleFunc("/GetTransactionByHash", GetTransactionByHashHandler)
	http.HandleFunc("/GetTransactions", GetTransactionsHandler)
	http.HandleFunc("/GetBlocks", GetBlocksHandler)
	http.HandleFunc("/GetAddressCount", GetAddressCountHandler)
	http.HandleFunc("/GetTransactionCount", GetTransactionCountHandler)
	http.HandleFunc("/GetBlockCount", GetBlockCountHandler)
	http.HandleFunc("/GetReceiptByHash", GetReceiptByHashHandler)
	http.HandleFunc("/GetReceipts", GetReceiptsHandler)
	http.HandleFunc("/GetReceiptsByLast", GetReceiptsByLastHandler)

	http.HandleFunc("/GetReceiptsByAddress", GetReceiptsByAddressHandler)
	http.HandleFunc("/GetTransactionsByAddress", GetTransactionsByAddressHandler)
	http.HandleFunc("/BalanceOfToken", GetBalanceOfTokenHandler)
	http.HandleFunc("/GetAddresses", GetAddressesHandler)
	http.HandleFunc("/GetAddressDetail", GetAddressDetailHandler)

	// 启动 HTTP 服务，默认监听 9966 端口
	fmt.Println("Starting server at http://localhost:9966")
	if err := http.ListenAndServe(":9966", nil); err != nil {
		panic(err)
	}
}

func GetBlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	number := r.URL.Query().Get("number")
	block := bsdb.GetBlockByNumber(number)
	err := json.NewEncoder(w).Encode(block)
	if err != nil {
		log.Println("GetBlockByNumber编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetTransactionByHashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	hash := r.URL.Query().Get("hash")
	transaction := bsdb.GetTransactionByHash(hash)
	err := json.NewEncoder(w).Encode(transaction)
	if err != nil {
		log.Println("GetTransactionByHash编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	s, err := strconv.Atoi(start)
	if err != nil {
		log.Println("转化参数start失败", start, err)
		http.Error(w, "Invalid  parameter", http.StatusBadRequest)
		return
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		log.Println("转化参数start失败", start, err)
		http.Error(w, "Invalid  parameter", http.StatusBadRequest)
		return
	}
	transactions := bsdb.GetTransactions(s, e)
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Println("GetTransactions编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}

}
func GetBlocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	s, err := strconv.Atoi(start)
	if err != nil {
		log.Println("转化参数start失败", start, err)
		http.Error(w, "Invalid  parameter", http.StatusBadRequest)
		return
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		log.Println("转化参数start失败", start, err)
		http.Error(w, "Invalid  parameter", http.StatusBadRequest)
		return
	}
	blocks := bsdb.GetBlocks(s, e)
	err = json.NewEncoder(w).Encode(blocks)
	if err != nil {
		log.Println("GetBlocks编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetTransactionCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	count := bsdb.GetTransactionCount()
	err := json.NewEncoder(w).Encode(count)
	if err != nil {
		log.Println("GetTransactionCount编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetBlockCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	count := bsdb.GetBlockCount()
	err := json.NewEncoder(w).Encode(count)
	if err != nil {
		log.Println("GetBlockCount编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetReceiptByHashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	hash := r.URL.Query().Get("hash")
	receipt := bsdb.GetReceiptByHash(hash)
	err := json.NewEncoder(w).Encode(receipt)
	if err != nil {
		log.Println("GetReceiptByHash编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	s, err := strconv.Atoi(start)
	if err != nil {
		log.Println("转化参数start失败", start, err)
		http.Error(w, "Invalid  parameter", http.StatusBadRequest)
		return
	}
	e, err := strconv.Atoi(end)
	if err != nil {
		log.Println("转化参数start失败", start, err)
		http.Error(w, "Invalid  parameter", http.StatusBadRequest)
		return
	}
	receipts := bsdb.GetReceipts(s, e)

	err = json.NewEncoder(w).Encode(receipts)
	if err != nil {
		log.Println("GetReceipts编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetReceiptsByLastHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	receipts := bsdb.GetReceiptsByLast()
	err := json.NewEncoder(w).Encode(receipts)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetReceiptsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	address := r.URL.Query().Get("address")
	receipts := bsdb.GetReceiptsByAddress(address)
	err := json.NewEncoder(w).Encode(receipts)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetBalanceOfTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	account := r.URL.Query().Get("address")
	token := r.URL.Query().Get("address")

	client, err := bs_eth.Dial(config.TestUrl)
	if err != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		log.Println("Connect ETH Error: ", err)
		return
	}
	defer client.Close()

	balance := bs_eth.TokenBalance(config.TestUrl, common.HexToAddress(account), token)

	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}

func GetAddressCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	count := bsdb.GetAddressCount()
	err := json.NewEncoder(w).Encode(count)
	if err != nil {
		log.Println("GetAddressCount编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetTransactionsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	address := r.URL.Query().Get("address")
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	isVal := common.IsHexAddress(address)
	if !isVal {
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

	transactions := bsdb.GetTransactionsByAddress(address, start, end)
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetAddressesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	list := GetAccountsList(config.TestUrl)
	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Println("GetTransactionCountByAddress 编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetAddressDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	address := r.URL.Query().Get("address")
	detail := GetAccountDetail(config.TestUrl, address)

	err := json.NewEncoder(w).Encode(detail)
	if err != nil {
		log.Println("GetTransactionCountByAddress 编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}

//func GetTokenCount
//func GetTokenInfoHandler(w http.ResponseWriter, r *http.Request){}
//func GetTokenList
//func GetTransactionsByAddress
