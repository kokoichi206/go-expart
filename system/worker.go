package main

import (
	"runtime"
	"sync"
)

func calccc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
}

func workerrr(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
	for year := range years {
		calc(id, price, interestRate, year)
		wg.Done()
	}
}

func calcMain() {
	price := 400_0000
	// 利子 1.1%
	interestRate := 0.011
	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		years <- i
	}
	var wg sync.WaitGroup
	wg.Add(35)
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}

	// 全てのワーカーの終了
	close(years)
	wg.Wait()
}
