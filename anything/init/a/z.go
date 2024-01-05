package a

import "fmt"

// global variable
var b = func() string {
	fmt.Println("a global variable")
	return "a"
}()

// init function
func init() {
	fmt.Println("a init")
}
