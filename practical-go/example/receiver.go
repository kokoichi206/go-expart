package example

import "fmt"

type ValueStruct struct {
	Counter int
}

type PointerStruct struct {
	Counter int
}

func (v ValueStruct) countUp() {
	v.Counter += 1
}

func (p *PointerStruct) countUp() {
	p.Counter += 1
}

// ポインターをレシーバーとするメソッドと、
// 値をレシーバーとするメソッドのどちらも定義することが可能！！！
func (p PointerStruct) countUpForValueStruct() {
	p.Counter += 1
}

func ReceiverTest() {
	v := ValueStruct{
		Counter: 0,
	}
	p := PointerStruct{
		Counter: 0,
	}
	v.countUp()
	fmt.Println("value struct: ", v.Counter) // 0
	p.countUp()
	fmt.Println("pointer struct: ", p.Counter) // 1

	p.countUpForValueStruct()
	fmt.Println("pointer struct: ", p.Counter) // 1 のまま

	x := &PointerStruct{
		Counter: 0,
	}
	x.countUpForValueStruct()
	fmt.Println("pointer struct: ", x.Counter) // 0
	x.countUp()
	fmt.Println("pointer struct: ", x.Counter) // 1
}
