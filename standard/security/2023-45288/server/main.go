package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/net/http2"
)

func checkMem() {
	for range time.Tick(1 * time.Second) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
		fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
		fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
		fmt.Printf("\tNumGC = %v\n", m.NumGC)
	}
}

func main() {
	printMemory := flag.Bool("m", false, "memory check")
	flag.Parse()

	if printMemory != nil && *printMemory {
		go checkMem()
	}

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
