package main

import (
	"fmt"
	"sync"
)

// $ cat main.go
//
// package main
//
// func main() {
// 	mutexSample()
// 	mutexSampleRace()
//
//
// $ go run -race *.go
//
// x[j]: 1
// x[j]: 2
// x[j]: 3
// x[j]: 1
// x[j]: 1
// x[j]: 2
// x[j]: 3
// x[j]: 2
// x[j]: 3
// x: [1 2 3]
// ==================
// WARNING: DATA RACE
// Read at 0x00c000110058 by goroutine 13:
//   main.mutexSampleRace.func1()
//       /Users/kokoichi/ghq/github.com/kokoichi206/go-expart/concurrency/demy/mutex.go:43 +0x39

// Previous write at 0x00c000110058 by goroutine 17:
//   main.mutexSampleRace.func1()
//       /Users/kokoichi/ghq/github.com/kokoichi206/go-expart/concurrency/demy/mutex.go:43 +0x4b

// Goroutine 13 (running) created at:
//   main.mutexSampleRace()
//       /Users/kokoichi/ghq/github.com/kokoichi206/go-expart/concurrency/demy/mutex.go:42 +0x8d
//   main.main()
//       /Users/kokoichi/ghq/github.com/kokoichi206/go-expart/concurrency/demy/main.go:31 +0x29

// Goroutine 17 (finished) created at:
//   main.mutexSampleRace()
//       /Users/kokoichi/ghq/github.com/kokoichi206/go-expart/concurrency/demy/mutex.go:42 +0x8d
//   main.main()
//       /Users/kokoichi/ghq/github.com/kokoichi206/go-expart/concurrency/demy/main.go:31 +0x29
// ==================
// count: 100
// Found 1 data race(s)
// exit status 66

func mutexSample() {
	x := []int{1, 2, 3}

	var wg sync.WaitGroup

	wg.Add(len(x))
	for i := 0; i < len(x); i++ {
		go manipulate(x, i, &wg)
	}

	wg.Wait()

	fmt.Printf("x: %v\n", x)
}

func manipulate(x []int, i int, wg *sync.WaitGroup) int {
	for j := 0; j < len(x); j++ {
		// x[j]++
		// -race 実行してみた結果、
		// アクセスのみは race condition にならなそう
		fmt.Printf("x[j]: %v\n", x[j])
	}

	defer wg.Done()

	return x[i]
}

func mutexSampleRace() {
	count := 0

	var wg sync.WaitGroup

	n := 100
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("count: %v\n", count)
}
