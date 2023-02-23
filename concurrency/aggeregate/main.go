package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	userName := fetchUser()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// unbuffered chan は何かくるまで待つ
	// buffered はそうじゃない, これ結構大事だ
	respch := make(chan any, 2)

	go fetchUserLinks(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	wg.Wait() // block until 2 wg.Done()
	close(respch)

	for resp := range respch {
		// fatal error: all goroutines are asleep - deadlock!
		// likes, ok := resp.(string)
		fmt.Println("resp: ", resp)
	}
	// fmt.Println("likes: ", links)
	// fmt.Println("matched: ", matched)

	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "John"
}

func fetchUserLinks(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 140)

	respch <- 11
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 240)

	respch <- "Doe"
	wg.Done()
}
