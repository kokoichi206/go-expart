package main

import (
	"fmt"
	"testing"
)

func BenchmarkListData(b *testing.B) {
	res, _ := ListData()
	fmt.Printf("res: %v\n", res)
}
