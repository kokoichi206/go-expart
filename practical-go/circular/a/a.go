package a

import (
	"fmt"
	"practical-go/circular/b"
)

func HelloA() {
	fmt.Println("Hello from A")
	b.HelloB()
}
