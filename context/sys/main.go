package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// // transport is an http.RoundTripper that keeps track of the in-flight
// // request and implements hooks to report HTTP tracing events.
// type transport struct {
// 	current *http.Request
// }

// // RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// // of the current request.
// func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	t.current = req
// 	return http.DefaultTransport.RoundTrip(req)
// }

// // GotConn prints whether the connection has been used previously
// // for the current request.
// func (t *transport) GotConn(info httptrace.GotConnInfo) {
// 	fmt.Printf("info: %v\n", info)
// 	// info.Conn.Close()
// 	fmt.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
// }

// t.dialConnFor(w) と pconn.readLoop() の分の増加が確認できる？
// pconn.writeLoop() に関してはなんで増加されない？
func numGoroutines() {
	// 1 msec とか細かくすると、call done 前に一瞬だけ 5 になるタイミングがある。
	interval := 50 * time.Millisecond

	for range time.Tick(interval) {
		fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 829*time.Millisecond)
	// ctx := context.Background()

	go numGoroutines()
	time.Sleep(150 * time.Millisecond)

	fmt.Println("make req")
	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://kokoichi0206.mydns.jp/", nil)
	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:7878", nil)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:21829", nil)

	// t := &transport{}
	// trace := &httptrace.ClientTrace{
	// 	GotConn: t.GotConn,
	// }
	// req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := http.DefaultClient
	// client := &http.Client{
	// 	Transport: t,
	// }

	// dialer := net.Dialer{}
	// conn, err := dialer.DialContext(ctx, "", "")

	// http.DefaultTransport

	fmt.Println("call start")

	resp, err := client.Do(req)
	fmt.Println("call done")

	time.Sleep(300 * time.Millisecond)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		if errors.Is(err, context.Canceled) {
			fmt.Println("canceled error")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("DeadlineExceeded error")
		}

		return
	}

	// http.DefaultTransport.RoundTrip()

	ctx.Done()
	defer resp.Body.Close()
}
