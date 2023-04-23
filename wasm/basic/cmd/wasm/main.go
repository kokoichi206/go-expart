package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = `
<h3>Hello, I'm an HTML from Go!!!</h3>
`

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	ch := make(chan struct{}, 0)
	fmt.Printf("hello web assembly from Go\n")

	js.Global().Set("getHtml", GetHtml())
	<-ch
}
