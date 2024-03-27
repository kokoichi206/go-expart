package main

import (
	"black/hat"
	"fmt"
)

func foo(i any) {
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	// hat.Proxy()
	// nc()

	hat.NcServer()
}
