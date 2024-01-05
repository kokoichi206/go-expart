package z

import "fmt"

// global variable
var b = func() string {
	fmt.Println("z global variable")
	return "z"
}()

// init function
func init() {
	fmt.Println("z init")
}
