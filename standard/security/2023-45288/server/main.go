package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP/2 world!")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// HTTP/2 を有効にする
	http2.ConfigureServer(&server, &http2.Server{})

	fmt.Println("Listening on https://localhost:8080")
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		panic(err)
	}
}
