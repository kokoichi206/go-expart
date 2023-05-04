package main

import (
	"testing"
)

type Value struct {
	context [64]byte
}

//go:noinline
func vf(v Value) Value {
	return v
}

//go:noinline
func vg(v *Value) *Value {
	return v
}

func Benchmark_ValueValue(b *testing.B) {
	b.ReportAllocs()
	var v Value

	for i := 0; i < b.N; i++ {
		_ = vf(v)
	}
}

func Benchmark_ValuePointer(b *testing.B) {
	b.ReportAllocs()
	var v Value

	for i := 0; i < b.N; i++ {
		_ = vg(&v)
	}
}
