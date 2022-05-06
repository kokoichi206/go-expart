package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Practice about Concurrent & Parallel
//
// Especially, using these functions
//   goroutines
func main() {
	calcLoan()
}

func calc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	fmt.Printf("year=%d total=%d interest=%d id=%d\n", year, price+interest, interest, id)
}
func worker(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
	// タスクがなくなってタスクのチャネルが close されるまで無限ループ
	for yaer := range years {
		calc(id, price, interestRate, yaer)
		wg.Done()
	}
}

func calcLoan() {
	// 借入額
	price := 4000_0000
	// 利子1.1%
	interestRate := 0.011
	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		years <- i
	}
	var wg sync.WaitGroup
	wg.Add(35)
	// CPU コア数分の goroutine 起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}
	close(years)
	wg.Wait()
}

func goroutineCost() {
	tasks := []string{
		"go build -o main main.go",
		"mv main share",
		"./publish",
	}
	for _, task := range tasks {
		go func() {
			// CAUTION!
			// goroutine is very fast, but not cost zero
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
