package main

import (
	"fmt"
	"sync"
)

// Pass the **pointer** of the waitGroup
func printSomething(s string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}

	wg.Add(len(words))

	for _, s := range words {
		go printSomething(s, &wg)
	}

	wg.Wait()
}
