package main

import (
	"fmt"
	"unique"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

func intern() {
	ps := []*Person{
		{"Alice", 20},
		{"Bob", 21},
		{"Alice", 20},
	}
	for i, p := range ps {
		m := unique.Make(p)
		ps[i] = m.Value()
	}
	for _, p := range ps {
		fmt.Println(unsafe.Pointer(p))
	}
}

func main() {
	intern()
}
