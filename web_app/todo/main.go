package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/kokoichi206/go-expert/web/todo/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// 今処理しているシグナルを返し終わってから終了する！
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 環境変数から設定値を読み込む
	config, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("Failed to listen port %d: %v", config.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// シグナルハンドリング確認用
			// time.Sleep(5 * time.Second)
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
