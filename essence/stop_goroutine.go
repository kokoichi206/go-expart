package main

import (
	"fmt"
	"math/rand"
)

func stoppableGoroutine() {
	quit := make(chan bool)
	ch := generator("Hi!", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	// quit を使って generate を止める！
	quit <- true
}

func generator(msg string, quit chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				fmt.Println("goroutine DONE!!")
				return
			}
		}
	}()

	return ch
}
