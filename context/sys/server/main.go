package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		for {
			select {
			// サーバーとして request のキャンセルを受け取る。
			case <-r.Context().Done():
				fmt.Println("Context DONE!!")

				return
			default:
				time.Sleep(300 * time.Millisecond)

				// 5 秒以上たったらクライアントに送信。
				if time.Now().After(now.Add(5 * time.Second)) {
					w.Write([]byte("hello"))

					return
				}
			}
		}
	})

	http.ListenAndServe(":21829", nil)
}
