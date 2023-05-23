package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Request struct {
	ID string `json:"id"`
}

type Response struct {
	Message string `json:"message"`
}

const (
	hostname     = "10.129.93.227:50051" // ホスト名とポートを指定
	serviceName  = "SimpleApp"    // サービス名を指定
	methodName   = "getInfo"       // メソッド名を指定
	tokenHeader = "token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoiYWFhYSIsImV4cCI6MTY4NDg2MTExOX0.8dyUoaB562MTJk-wznf39-HiaVjmlwhGVtY7YRzNb2k"
	httpPort     = ":8051"           // HTTPのポート番号を指定
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// GETメソッドのクエリパラメータを取得
		queryID := r.URL.Query().Get("id")

		// JSON形式のリクエストボディを作成
		reqBody, err := json.Marshal(Request{ID: queryID})
		if err != nil {
			log.Printf("Error marshaling request body: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// HTTPリクエストをgRPCリクエストに変換して転送する
		cmd := exec.Command("grpcurl", "--plaintext", "--rpc-header", tokenHeader, "-d", string(reqBody), hostname, serviceName+"."+methodName)
		output, err := cmd.Output()
		if err != nil {
			log.Printf("Error executing grpcurl: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := string(output)
		if strings.HasPrefix(resp, "ERROR:") {
			// エラーレスポンスの場合はそのまま返す
			http.Error(w, resp, http.StatusInternalServerError)
			return
		}

		// レスポンスをパースしてJSONの中身を取得
		var response Response
		err = json.Unmarshal([]byte(resp), &response)
		if err != nil {
			log.Printf("Error unmarshaling response body: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// JSONの中身をレスポンスとして返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// HTTPサーバーを起動
	log.Println("HTTP to gRPC proxy server started on", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}