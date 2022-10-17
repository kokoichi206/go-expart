package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	config "github.com/kokoichi206/go-expert/web/todo/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
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

	mux, cleanup, err := NewMux(ctx, config)
	if err != nil {
		return err
	}
	defer cleanup()

	s := NewServer(mux, l)
	return s.Run(ctx)
}
