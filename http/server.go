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
)

// corsMiddleware 设置CORS头
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		
		// 处理OPTIONS预检请求
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

// jsonResponse 统一JSON响应处理
func jsonResponse(w http.ResponseWriter, data interface{}, err error) {
	if err != nil {
		log.Printf("处理请求失败: %v", err)
		http.Error(w, fmt.Sprintf("处理请求失败: %v", err), http.StatusInternalServerError)
		return
	}
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON编码失败: %v", err)
		http.Error(w, "JSON编码失败", http.StatusInternalServerError)
		return
	}
}

func StartHttp() {
	// 注册路由处理函数
	http.HandleFunc("/GetBlockByNumber", corsMiddleware(GetBlockByNumberHandler))
	http.HandleFunc("/GetTransactionByHash", corsMiddleware(GetTransactionByHashHandler))
	http.HandleFunc("/GetTransactions", corsMiddleware(GetTransactionsHandler))
	http.HandleFunc("/GetBlocks", corsMiddleware(GetBlocksHandler))
	http.HandleFunc("/GetAddressCount", corsMiddleware(GetAddressCountHandler))
	http.HandleFunc("/GetTransactionCount", corsMiddleware(GetTransactionCountHandler))
	http.HandleFunc("/GetBlockCount", corsMiddleware(GetBlockCountHandler))
	http.HandleFunc("/GetReceiptByHash", corsMiddleware(GetReceiptByHashHandler))
	http.HandleFunc("/GetReceipts", corsMiddleware(GetReceiptsHandler))
	http.HandleFunc("/GetReceiptsByLast", corsMiddleware(GetReceiptsByLastHandler))

	http.HandleFunc("/GetReceiptsByAddress", corsMiddleware(GetReceiptsByAddressHandler))
	http.HandleFunc("/GetTransactionsByAddress", corsMiddleware(GetTransactionsByAddressHandler))
	http.HandleFunc("/BalanceOfToken", corsMiddleware(GetBalanceOfTokenHandler))
	http.HandleFunc("/GetAddresses", corsMiddleware(GetAddressesHandler))
	http.HandleFunc("/GetAddressDetail", corsMiddleware(GetAddressDetailHandler))

	// 启动 HTTP 服务，默认监听 9966 端口
	fmt.Println("Starting server at http://localhost:9966")
	if err := http.ListenAndServe(":9966", nil); err != nil {
		log.Fatalf("HTTP服务器启动失败: %v", err)
	}
}

func GetBlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	if number == "" {
		http.Error(w, "缺少参数: number", http.StatusBadRequest)
		return
	}
	
	block, err := bsdb.GetBlockByNumber(number)
	if err != nil {
		log.Printf("GetBlockByNumber查询失败: %v", err)
		http.Error(w, "查询区块失败", http.StatusNotFound)
		return
	}
	
	jsonResponse(w, block, nil)
}
func GetTransactionByHashHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "缺少参数: hash", http.StatusBadRequest)
		return
	}
	
	transaction, err := bsdb.GetTransactionByHash(hash)
	if err != nil {
		log.Printf("GetTransactionByHash查询失败: %v", err)
		http.Error(w, "查询交易失败", http.StatusNotFound)
		return
	}
	
	jsonResponse(w, transaction, nil)
}
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
	
	transactions, err := bsdb.GetTransactions(start, end)
	if err != nil {
		log.Printf("GetTransactions查询失败: %v", err)
		http.Error(w, "查询交易列表失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, transactions, nil)
}
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
	
	blocks, err := bsdb.GetBlocks(start, end)
	if err != nil {
		log.Printf("GetBlocks查询失败: %v", err)
		http.Error(w, "查询区块列表失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, blocks, nil)
}
func GetTransactionCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := bsdb.GetTransactionCount()
	if err != nil {
		log.Printf("GetTransactionCount查询失败: %v", err)
		http.Error(w, "查询交易总数失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, count, nil)
}
func GetBlockCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := bsdb.GetBlockCount()
	if err != nil {
		log.Printf("GetBlockCount查询失败: %v", err)
		http.Error(w, "查询区块总数失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, count, nil)
}
func GetReceiptByHashHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "缺少参数: hash", http.StatusBadRequest)
		return
	}
	
	receipt, err := bsdb.GetReceiptByHash(hash)
	if err != nil {
		log.Printf("GetReceiptByHash查询失败: %v", err)
		http.Error(w, "查询收据失败", http.StatusNotFound)
		return
	}
	
	jsonResponse(w, receipt, nil)
}
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
	
	receipts, err := bsdb.GetReceipts(start, end)
	if err != nil {
		log.Printf("GetReceipts查询失败: %v", err)
		http.Error(w, "查询收据列表失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, receipts, nil)
}
func GetReceiptsByLastHandler(w http.ResponseWriter, r *http.Request) {
	receipts := bsdb.GetReceiptsByLast()
	jsonResponse(w, receipts, nil)
}
func GetReceiptsByAddressHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}
	
	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}
	
	receipts, err := bsdb.GetReceiptsByAddress(address)
	if err != nil {
		log.Printf("GetReceiptsByAddress查询失败: %v", err)
		http.Error(w, "查询收据失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, receipts, nil)
}
func GetBalanceOfTokenHandler(w http.ResponseWriter, r *http.Request) {
	account := r.URL.Query().Get("address")
	token := r.URL.Query().Get("token")
	
	if account == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}
	
	if token == "" {
		http.Error(w, "缺少参数: token", http.StatusBadRequest)
		return
	}
	
	if !common.IsHexAddress(account) {
		http.Error(w, "无效的账户地址格式", http.StatusBadRequest)
		return
	}
	
	if !common.IsHexAddress(token) {
		http.Error(w, "无效的代币地址格式", http.StatusBadRequest)
		return
	}

	client, err := bs_eth.Dial(config.TestUrl)
	if err != nil {
		log.Printf("连接ETH客户端失败: %v", err)
		http.Error(w, "连接区块链节点失败", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	balance := bs_eth.TokenBalance(config.TestUrl, common.HexToAddress(account), token)
	jsonResponse(w, balance, nil)
}

func GetAddressCountHandler(w http.ResponseWriter, r *http.Request) {
	count, err := bsdb.GetAddressCount()
	if err != nil {
		log.Printf("GetAddressCount查询失败: %v", err)
		http.Error(w, "查询地址数量失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, count, nil)
}
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

	transactions, err := bsdb.GetTransactionsByAddress(address, start, end)
	if err != nil {
		log.Printf("GetTransactionsByAddress查询失败: %v", err)
		http.Error(w, "查询交易失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, transactions, nil)
}
func GetAddressesHandler(w http.ResponseWriter, r *http.Request) {
	list, err := GetAccountsList(config.TestUrl)
	if err != nil {
		log.Printf("GetAddresses查询失败: %v", err)
		http.Error(w, "查询地址列表失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, list, nil)
}
func GetAddressDetailHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "缺少参数: address", http.StatusBadRequest)
		return
	}
	
	if !common.IsHexAddress(address) {
		http.Error(w, "无效的地址格式", http.StatusBadRequest)
		return
	}
	
	detail, err := GetAccountDetail(config.TestUrl, address)
	if err != nil {
		log.Printf("GetAddressDetail查询失败: %v", err)
		http.Error(w, "查询地址详情失败", http.StatusInternalServerError)
		return
	}
	
	jsonResponse(w, detail, nil)
}

