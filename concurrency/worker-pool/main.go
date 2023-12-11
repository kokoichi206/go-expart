package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const (
	limit = 3
)

func do(input string) {
	time.Sleep(1 * time.Second)
	fmt.Printf("input: %v\n", input)
}

// $ go run main.go
// input: 1
// input: 3
// input: 2
// input: 33
// input: 22
// input: 11
// input: 111
// input: 333
// input: 222
// input: 2222
// input: 3333
// input: 1111
// input: 11111
// input: 33333
// input: 22222
// channelPattern done
// ^Csignal: interrupt
func channelPattern(ctx context.Context, inputs []string) {
	defer func() { fmt.Println("channelPattern done") }()

	pool := make(chan struct{}, limit)

	eg, ctx := errgroup.WithContext(ctx)

	for _, v := range inputs {
		v := v

		pool <- struct{}{}

		eg.Go(func() error {
			do(v)
			<-pool
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("err eg.Wait(): %v\n", err)
	}
}

func channelPatternWithCancel(ctx context.Context, inputs []string) {
	defer func() { fmt.Println("channelPatternWithCancel done") }()

	pool := make(chan struct{}, limit)

	eg, ctx := errgroup.WithContext(ctx)

FORLOOP:
	for _, v := range inputs {
		v := v

		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Done() v=%v: %v\n", v, ctx.Err())

			break FORLOOP
		case pool <- struct{}{}:
		}

		eg.Go(func() error {
			do(v)
			<-pool
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("err eg.Wait(): %v\n", err)
	}
}

// ❯ go run main.go
// input: 1
// input: 3
// input: 2
// input: 11
// input: 33
// input: 22
// cancel:  222
// ctx.Done(): context canceled
// input: 111
// err eg.Wait(): error: 33
// channelPatternEGCtx done
//
// see: https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext
func channelPatternEGCtx(ctx context.Context, inputs []string) {
	defer func() { fmt.Println("channelPatternEGCtx done") }()

	pool := make(chan struct{}, limit)

	eg, ctx := errgroup.WithContext(ctx)

FORLOOP:
	for _, v := range inputs {
		v := v

		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Done() v=%v: %v\n", v, ctx.Err())
			break FORLOOP
		case pool <- struct{}{}:
		}

		eg.Go(func() error {
			do(v)
			<-pool

			if v == "33" {
				return fmt.Errorf("error: %v", v)
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("err eg.Wait(): %v\n", err)
	}
}

// $ go run main.go
// input: 1
// input: 2
// input: 3
// err sem.Acquire(): context deadline exceeded
// input: 11
// input: 33
// input: 22
// weightedPattern done
// ^Csignal: interrupt
//
// see: https://pkg.go.dev/golang.org/x/sync/semaphore
func weightedPattern(ctx context.Context, inputs []string) {
	defer func() { fmt.Println("weightedPattern done") }()

	sem := semaphore.NewWeighted(limit)

	var eg errgroup.Group

	for _, v := range inputs {
		v := v

		// キャンセルが起きた時などは、ここで sem.Acquire() が失敗する。
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("err sem.Acquire(): %v\n", err)
			break
		}

		eg.Go(func() error {
			do(v)
			sem.Release(1)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("err eg.Wait(): %v\n", err)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	inputs := []string{
		"1", "2", "3",
		"11", "22", "33",
		"111", "222", "333",
		"1111", "2222", "3333",
		"11111", "22222", "33333",
	}

	// channelPattern(ctx, inputs)
	weightedPattern(ctx, inputs)
	// channelPatternWithCancel(ctx, inputs)

	// channelPatternEGCtx(ctx, inputs)

	// 終わらないための工夫。
	time.Sleep(100 * time.Second)
}
