package main

import "fmt"

// global variable
var z = func() string {
	fmt.Println("main global variable z.go")
	return "z"
}()

func init() {
	println("main init z.go")
}
