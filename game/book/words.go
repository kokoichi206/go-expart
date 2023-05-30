package main

import (
	"fmt"
	"index/suffixarray"
)

func words() {
	ngIdx := suffixarray.New([]byte("foobarrpien"))
	res := ngIdx.Lookup([]byte("pie"), -1)
	fmt.Printf("res: %v\n", res)
}
