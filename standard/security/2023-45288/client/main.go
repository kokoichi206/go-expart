package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// basic()
	// return

	// TLS 接続の設定
	tlsConfig := &tls.Config{
		// 実運用では危険。
		InsecureSkipVerify: true,
		// HTTP/2 を明示的に指定
		NextProtos: []string{"h2"},
	}

	// TLS 接続を確立
	conn, err := tls.Dial("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	customHttp2Call(conn)
}

// frame とかを自分で扱わない場合の普通の使い方。
func basic() {
	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 自己署名証明書を使用する場合
			},
		},
	}

	req, err := http.NewRequest("GET", "https://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	// ヘッダーのカスタマイズ
	req.Header.Add("Custom-Header", "MyValue")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Protocol: %s\n", resp.Proto) // 確認のためプロトコルを表示
}
