package main

import "fmt"

func main() {
	defer fmt.Println("finished...")
	fmt.Println("hello world!")
}
