package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

type Request struct {
	ID string `json:"id"`
}

const (
	hostname     = "10.129.93.227:50051" // ホスト名とポートを指定
	serviceName  = "SimpleApp"    // サービス名を指定
	methodName   = "getInfo"       // メソッド名を指定
	tokenHeader = "token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoidGVzdHUiLCJleHAiOjE2ODQ3NzI0MTd9.DfVyzGde6IjLfxgin65F7aAaplqMxJ0jteB6CQtVaVo"
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
		cmd.Stdout = w

		err = cmd.Run()
		if err != nil {
			log.Printf("Error executing grpcurl: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Println("HTTP to gRPC proxy server started on", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}