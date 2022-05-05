package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timer(1)
}

func timer(sec int) {
	timeout := time.After(time.Duration(sec) * time.Second)

	// このforループを1秒間ずっと実行し続ける
	for {
		select {
		case <-timeout:
			fmt.Println("time out")
			return
		default:
			// fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func signalNotify() {
	signals := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C) を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルがくるまで待つ
	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}

func printPrimeNumbers() {
	pn := primeNumbers()
	for n := range pn {
		fmt.Println(n)
	}
}

func primeNumbers() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 1000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
			time.Sleep(500 * time.Millisecond)
		}
		close(result)
	}()
	return result
}

func chanel() {
	fmt.Println("start sub()")
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("sub is finished")
		done <- struct{}{}
	}()
	<-done
	fmt.Println("all tasks are finished")
}

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func goroutine() {
	fmt.Println("start sub()")
	go sub()

	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
