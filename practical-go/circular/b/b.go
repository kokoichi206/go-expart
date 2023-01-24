package b

import (
	"fmt"
	"practical-go/circular/c"
)
func HelloB() {
	fmt.Println("Hello from B")
	c.HelloC()
}
