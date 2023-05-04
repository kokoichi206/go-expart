package main

import (
	"fmt"
	"io"
	"testing"
)

//go:noinline
func fLoop(v []pien) {
	for _, p := range v {
		fmt.Fprintf(io.Discard, "p: %v", p)
	}
}

//go:noinline
func gLoop(v []*pien) {
	for _, p := range v {
		fmt.Fprintf(io.Discard, "p: %v", p)
	}
}

func Benchmark_SliceValueLoop(b *testing.B) {
	b.ReportAllocs()
	size := 1000
	v := make([]pien, size)
	for i := 0; i < size; i++ {
		v[i] = pien{
			paon:   "hogehogehogehoge",
			hoge:   "fugafugafugafuga",
			taihen: "taihen taihenda",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fLoop(v)
	}
}

func Benchmark_SlicePointerLoop(b *testing.B) {
	b.ReportAllocs()
	size := 1000
	v := make([]*pien, size)
	for i := 0; i < size; i++ {
		v[i] = &pien{
			paon:   "hogehogehogehoge",
			hoge:   "fugafugafugafuga",
			taihen: "taihen taihenda",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gLoop(v)
	}
}
