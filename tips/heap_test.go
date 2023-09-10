package main

import "testing"

var globalValue int

func BenchmarkSumValue(b *testing.B) {
	// heap への割り当てを測定する。
	b.ReportAllocs()
	var local int
	for i := 0; i < b.N; i++ {
		local = sumValue(i, i)
	}
	globalValue = local
}

func BenchmarkPointer(b *testing.B) {
	// heap への割り当てを測定する。
	b.ReportAllocs()
	var local *int
	for i := 0; i < b.N; i++ {
		local = sumPointer(i, i)
	}
	globalValue = *local
}
