package main

import (
	"fmt"
)

// 型がインタフェースを満たしていることの保証を IDE 等でしたい！
var _ I = (*foo)(nil)

type I interface {
	doSomething()
}

type foo struct{}

func (f *foo) doSomething() {}

func main() {
	stoppableGoroutine()
	return
	sendMail()
	firstServer()
	// stringer コマンドで生成された文字列、Orange が表示される
	fmt.Println(Orange)
	ls()
	service()

	// recover による panic からの復帰！
	// recover は defer の中から呼び出す。
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	var a [2]int
	n := 2
	println(a[n])
}
