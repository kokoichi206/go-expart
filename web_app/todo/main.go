package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {

	if len(os.Args) != 2 {
		log.Printf("need port number to start\n")
		os.Exit(1)
	}

	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		// port 番号以外（数値以外など）がきても、ここで弾ける
		log.Fatalf("Failed to listen port %s: %v", p, err)
	}

	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンで HTTP サーバーの起動
	// (sync.WaitGroup では、goroutine が終わるのを待つけど、エラーがあったかどうかまでわからない。)
	eg.Go(func() error {
		// ErrServerClosed は Shutdown が正常に終了したことを示す（≠異常）。
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// チャネルからの終了通知を待機！
	// これ誰から来るんだっけ？ <- Context からくる！
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("Faailed to shutdown: %+v", err)
	}
	return eg.Wait()
}
