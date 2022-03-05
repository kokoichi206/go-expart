package main

import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the data is %v.\n", data)
	}
	
	// intStream := make(chan int)
	// close(intStream)
	// // 閉じたチャネルからも読み取れる
	// integer, ok := <-intStream
	// fmt.Printf("(%v): %v\n", ok, integer)

	intStream := make(chan int)
	go func() {
		defer close(intStream)

		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	// チャネルを引数にとり、チャネルが閉じた時に自動的にループを終了する
	for integer := range intStream {
		fmt.Printf("%v \n", integer)
	}
}
