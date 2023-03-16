package main

// 型がインタフェースを満たしていることの保証を IDE 等でしたい！
var _ I = (*foo)(nil)

type I interface {
	doSomething()
}

type foo struct{}

func (f *foo) doSomething(){}

func main() {
}
