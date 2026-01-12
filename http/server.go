package bshttp

import (
	"encoding/json"
	"fmt"
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

// StartHttp 启动HTTP服务
func StartHttp() {
	// 注册路由处理函数
	// Block相关路由
	http.HandleFunc("/GetBlockByNumber", corsMiddleware(GetBlockByNumberHandler))
	http.HandleFunc("/GetBlocks", corsMiddleware(GetBlocksHandler))
	http.HandleFunc("/GetBlockCount", corsMiddleware(GetBlockCountHandler))

	// Transaction相关路由
	http.HandleFunc("/GetTransactionByHash", corsMiddleware(GetTransactionByHashHandler))
	http.HandleFunc("/GetTransactions", corsMiddleware(GetTransactionsHandler))
	http.HandleFunc("/GetTransactionCount", corsMiddleware(GetTransactionCountHandler))
	http.HandleFunc("/GetTransactionsByAddress", corsMiddleware(GetTransactionsByAddressHandler))
	http.HandleFunc("/GetTransactionCountByAddress", corsMiddleware(GetTransactionCountByAddressHandler))

	// Receipt相关路由
	http.HandleFunc("/GetReceiptByHash", corsMiddleware(GetReceiptByHashHandler))
	http.HandleFunc("/GetReceipts", corsMiddleware(GetReceiptsHandler))
	http.HandleFunc("/GetReceiptCount", corsMiddleware(GetReceiptCountHandler))
	http.HandleFunc("/GetReceiptsByLast", corsMiddleware(GetReceiptsByLastHandler))
	http.HandleFunc("/GetReceiptsByAddress", corsMiddleware(GetReceiptsByAddressHandler))
	http.HandleFunc("/GetReceiptCountByAddress", corsMiddleware(GetReceiptCountByAddressHandler))

	// Account相关路由
	http.HandleFunc("/GetAccount", corsMiddleware(GetAccountHandler))
	http.HandleFunc("/GetAccounts", corsMiddleware(GetAccountsHandler))
	http.HandleFunc("/GetAccountCount", corsMiddleware(GetAccountCountHandler))

	// Token相关路由
	http.HandleFunc("/GetToken", corsMiddleware(GetTokenHandler))
	http.HandleFunc("/GetTokens", corsMiddleware(GetTokensHandler))
	http.HandleFunc("/GetTokenCount", corsMiddleware(GetTokenCountHandler))

	// NFT相关路由
	http.HandleFunc("/GetNFT", corsMiddleware(GetNFTHandler))
	http.HandleFunc("/GetNFTs", corsMiddleware(GetNFTsHandler))
	http.HandleFunc("/GetNFTCount", corsMiddleware(GetNFTCountHandler))

	// 启动 HTTP 服务，默认监听 9966 端口
	fmt.Println("Starting server at http://localhost:9966")
	if err := http.ListenAndServe(":9966", nil); err != nil {
		log.Fatalf("HTTP服务器启动失败: %v", err)
	}
}
