package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	// t.Skip でテストを飛ばせる！！
	t.Skip("Now refactoring...")

	// テストでは空いてるポートから自由に選ぶ
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to listen port %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// go test -v とするとテストが成功した時でも Logf の出力結果を確認できる
	t.Logf("try request to %q", url)

	rsp, err := http.Get(url)
	if err != nil {
		// %+v のフラグ付きだと、フィールドの名前もついてくる！
		// https://pkg.go.dev/fmt
		// 	%v	the value in a default format
		// when printing structs, the plus flag (%+v) adds field names
		t.Errorf("Failed to get: %+v", err)
	}
	// この Read とかより先に Close するスタイル使っていきたい。
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("Failed to read body: %+v", err)
	}
	want := fmt.Sprintf("Hello, %s", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run 関数に終了通知を送信！
	// コンテキストから（と思ってよき？）
	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, _ := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	// Server の責務からルーティングを切り離しているので、テスト用のルーティングで良い
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
	})

	eg.Go(func() error {
		// return が複数あるのは Go メソッドの説明にそう書いてある。
		return run(ctx)
		s := NewServer(mux, l)
		return s.Run(ctx)
	})
}
