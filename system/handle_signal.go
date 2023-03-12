package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handle() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	s := <- signals

	switch s {
	case syscall.SIGINT:
		fmt.Printf("SIGINT")
	case syscall.SIGTERM:
		fmt.Printf("SIGTERM")
	}

	// シグナルによりサイドハンドラが呼ばれないようにする。
	signal.Stop(signals)
}
