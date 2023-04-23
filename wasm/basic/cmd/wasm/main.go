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

// test of arguments
func add() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// <undefined>
		fmt.Println(this)
		// [<number: 1> <number: 2>], [<number: 1> pien]
		fmt.Println(args)
		// str に対して Int() を呼ぶと panic する！
		// wasm_exec.js:22 panic: syscall/js: call of Value.Int on string
		a := args[0].Int()
		b := args[1].Int()
		return a + b
	})
}

func main() {
	ch := make(chan struct{}, 0)
	fmt.Printf("hello web assembly from Go\n")

	js.Global().Set("getHtml", GetHtml())
	js.Global().Set("add", add())
	<-ch
}
