package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for x := range in {
			out <- x * x
		}

		close(out)
	}()

	return out
}

func pipelineSample() {
	fmt.Println("========== pipelineSample ==========")

	inCh := generator(1, 2, 3, 4, 5)

	outCh := square(inCh)

	for x := range outCh {
		println(x)
	}

	inCh2 := generator(1, 2, 3, 4, 5)

	composedCh := square(square(inCh2))

	for x := range composedCh {
		println(x)
	}
}

func fanOutFanIn() {
	fmt.Println("========== fanOutFanIn ==========")

	inCh := generator(1, 2, 3, 4, 5)

	ch1 := square(inCh)
	ch2 := square(inCh)

	for x := range merge(ch1, ch2) {
		println(x)
	}
}

// Fan In
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}

		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generatorWithDoneCh(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func squareWithDoneCh(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for x := range in {
			select {
			case out <- x * x:
			case <-done:
				return
			}
		}
	}()

	return out
}

// Fan In
func mergeWithDoneCh(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()

		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func cancelGoroutine() {
	fmt.Println("========== cancelGoroutine ==========")

	done := make(chan struct{})
	inCh := generatorWithDoneCh(done, 1, 2, 3, 4, 5)

	ch1 := squareWithDoneCh(done, inCh)
	ch2 := squareWithDoneCh(done, inCh)

	out := mergeWithDoneCh(done, ch1, ch2)

	// cancel goroutines after receiving one value.
	fmt.Println(<-out)
	close(done)

	time.Sleep(1 * time.Second)
	fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
}
