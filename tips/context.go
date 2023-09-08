package main

import "context"

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
