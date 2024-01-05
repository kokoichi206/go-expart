package main

// import "fmt"
import (
	"fmt"
	_ "initialize-order/a"
	_ "initialize-order/z"
	// _ "initialize-order/repository"
	// _ "initialize-order/usecase"
)

// global variable
var m = func() string {
	fmt.Println("main global variable")
	return "m"
}()

// global variable
var m2 = func() string {
	fmt.Println("main global variable2")
	return "m2"
}()

// init function
func init() {
	fmt.Println("main init")
}

func init() {
	fmt.Println("main init2")
}

func main() {
	fmt.Println("main")
}
