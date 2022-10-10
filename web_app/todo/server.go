package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

// ルーティングは引数で渡すことで、Server struct の責務から除外する。
func NewServer(mux http.Handler, l net.Listener) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// 今処理しているシグナルを返し終わってから終了する！
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンで HTTP サーバーの起動
	// (sync.WaitGroup では、goroutine が終わるのを待つけど、エラーがあったかどうかまでわからない。)
	eg.Go(func() error {
		// ErrServerClosed は Shutdown が正常に終了したことを示す（≠異常）。
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// チャネルからの終了通知を待機！
	// これ誰から来るんだっけ？ <- Context からくる！
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("Faailed to shutdown: %+v", err)
	}
	return eg.Wait()
}
