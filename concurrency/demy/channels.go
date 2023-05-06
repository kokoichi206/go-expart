package main

import "fmt"

func getnMsg(ch chan<- string) {
	ch <- "hello"
}

func replayMsg(ch1 <-chan string, ch2 chan<- string) {
	m := <-ch1

	ch2 <- m
}

func chanDirection() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go getnMsg(ch1)

	go replayMsg(ch1, ch2)

	v := <-ch2

	println(v)
}

func chanOwner() {
	owner := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)

			for i := 0; i < 5; i++ {
				ch <- i
			}
		}()

		return ch
	}

	consumer := func(ch <-chan int) {
		for v := range ch {
			println(v)
		}

		fmt.Println("done")
	}

	ch := owner()
	consumer(ch)
}
