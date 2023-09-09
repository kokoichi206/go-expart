package main

import (
	"context"
	"time"
)

func cancelable() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx.Done()
}

type key string

const myKey key = "myKey"

func kv() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, myKey, "value")
	ctx.Value(myKey)
}

func someFunc(ctx context.Context) {
	_ = detach{ctx}
}

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (d detach) Done() <-chan struct{} {
	return nil
}

func (d detach) Err() error {
	return nil
}

func (d detach) Value(key any) any {
	// 親コンテキストの Value を返す。
	return d.ctx.Value(key)
}
