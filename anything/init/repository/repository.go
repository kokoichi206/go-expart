package repository

import "fmt"

// global variable
var b = func() string {
	fmt.Println("repository global variable")
	return "repository"
}()

// init function
func init() {
	fmt.Println("repository init")
}
