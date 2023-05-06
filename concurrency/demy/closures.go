package main

import (
	"fmt"
	"sync"
)

func localAccess() {
	wg := sync.WaitGroup{}

	inside := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)

		go func() {
			defer wg.Done()

			i = i + 1
			fmt.Printf("i: %v\n", i)
		}()

		fmt.Println("returning from inside")

		return
	}

	inside(&wg)
	wg.Wait()
	fmt.Println("done")
}

func loopIndex() {
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Printf("i: %v\n", i)
		}()

		// go func(i int) {
		// 	defer wg.Done()
		// 	fmt.Printf("fixed ver i: %v\n", i)
		// }(i)
	}

	wg.Wait()
	fmt.Println("done")
}
