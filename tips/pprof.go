package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

/*
## 色々と操作する前
## gc=1 は GC した後にプロファイリング取得
curl -s http://localhost:20829/debug/pprof/heap?gc=1 > heap-before.pprof

## 色々と操作した後
curl -s http://localhost:20829/debug/pprof/heap?gc=1 > heap-after.pprof
curl -s http://localhost:20829/debug/pprof/heap?gc=1 > heap-after-2.pprof

## 比較表示
go tool pprof -http=:4646 -diff_base heap-before.pprof heap-after-2.pprof
*/
func servePProf() {
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for range ticker.C {
			fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})
	http.ListenAndServe(":20829", nil)
}
