package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	Interrupt os.Signal = syscall.SIGINT
	Kill      os.Signal = syscall.SIGINT
)

func main() {
	signalHandler()
}

func signalHandler() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, Interrupt, Kill)

	s := <-signals
	switch s {
	case Interrupt:
		fmt.Println("SIGINT")
	case Kill:
		fmt.Println("SIGTERM")
	}
}

func cmdExample() {
	// cmd := exec.Command("rm main.go")
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
