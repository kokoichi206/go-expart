package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
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
