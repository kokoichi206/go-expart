package main

import (
	"context"
	"fmt"
	"time"
)

type ctxResult struct {
	err   error
	timer *time.Timer
}

type Cmd struct {
	ctxResult <-chan ctxResult

	ctx context.Context
}

func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	if ctx == nil {
		panic("nil Context")
	}
	cmd := &Cmd{}

	cmd.ctx = ctx

	return cmd
}

func (c *Cmd) Run() error {
	if err := c.Start(); err != nil {
		return err
	}
	return c.Wait()
}

func (c *Cmd) Start() error {
	fmt.Printf("c.ctx: %v\n", c.ctx)

	if c.ctx != nil {
		// バッファなしチャネルは、送信と受信が同時に行われないとデッドロックする。
		resultc := make(chan ctxResult)
		c.ctxResult = resultc

		go c.watchCtx(resultc)
	}

	return nil
}

func (c *Cmd) watchCtx(resultc chan<- ctxResult) {
	fmt.Println("watchCtx")
	select {
	case resultc <- ctxResult{}:
		fmt.Println("resultc <- ctxResult{} and return")
		return
	case <-c.ctx.Done():
	}

	// ...
	err := fmt.Errorf("timeout?: %w", c.ctx.Err())

	fmt.Printf("c.ctx.Err(): %v\n", c.ctx.Err())

	resultc <- ctxResult{err: err}
}

func (c *Cmd) Wait() error {

	// 時間がかかる実行のシミュレーション
	time.Sleep(2 * time.Second)

	fmt.Printf("c.ctxResult: %v\n", c.ctxResult)
	if c.ctxResult != nil {
		// バッファなしチャネルの受信を開始する（ここでようやく送信が可能になる！）
		watch := <-c.ctxResult

		fmt.Printf("watch: %v\n", watch)
	}

	return nil
}
