package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var a, b int
var wg sync.WaitGroup

func f() {
	defer wg.Done()
	a = 1
	b = 2
}

func g() {
	defer wg.Done()
	r1 := b
	r2 := a

	if r1 == 0 && r2 == 0 {
		panic("pien")
	}
}

func main() {
	defer wg.Wait()

	// wg.Add(2)
	// go f()
	// go g()

	wg.Add(2)
	go func() {
		defer wg.Done()
		x = 1

		y.Add(1)
	}()
	go func() {
		defer wg.Done()
		x = 2
		fmt.Printf("x: %v\n", x)

		fmt.Printf("y.Load(): %v\n", y.Load())
	}()
}

var x int

var y atomic.Int64

var z sync.Map
