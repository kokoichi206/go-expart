package main

import "fmt"

// global variable
var a = func() string {
	fmt.Println("main global variable a.go")
	return "a"
}()

func init() {
	println("main init a.go")
}
