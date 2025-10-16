package bshttp

import (
	bsdb "blockchain_services/postgres"
	"encoding/json"
	"fmt"
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
	http.HandleFunc("/BalanceOfErc20", BalanceOfErc20Handler)
	http.HandleFunc("/BalanceOfErc404", BalanceOfErc404Handler)

	// 启动 HTTP 服务，默认监听 9966 端口
	fmt.Println("Starting server at http://localhost:9966")
	if err := http.ListenAndServe(":9966", nil); err != nil {
		panic(err)
	}
}

func GetBlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
func GetAddressCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count, err := bsdb.GetAddressCount()
	if err != nil {
		log.Println("GetAddressCount查询失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(count)
	if err != nil {
		log.Println("GetAddressCount编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetTransactionCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count, err := bsdb.GetTransactionCount()
	if err != nil {
		log.Println("GetTransactionCount查询失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(count)
	if err != nil {
		log.Println("GetTransactionCount编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetBlockCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count, err := bsdb.GetBlockCount()
	if err != nil {
		log.Println("GetBlockCount查询失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(count)
	if err != nil {
		log.Println("GetBlockCount编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetReceiptByHashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	receipts := bsdb.GetReceiptsByLast()
	err := json.NewEncoder(w).Encode(receipts)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetTransactionsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	address := r.URL.Query().Get("address")
	transactions := bsdb.GetTransactionsByAddress(address)
	err := json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func GetReceiptsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	address := r.URL.Query().Get("address")
	receipts := bsdb.GetReceiptsByAddress(address)
	err := json.NewEncoder(w).Encode(receipts)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func BalanceOfErc20Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//address := r.URL.Query().Get("address")
	balance := 0
	err := json.NewEncoder(w).Encode(balance)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
func BalanceOfErc404Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//address := r.URL.Query().Get("address")
	balance := 0
	err := json.NewEncoder(w).Encode(balance)
	if err != nil {
		log.Println("GetReceiptsByLast编码失败", err)
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
}
