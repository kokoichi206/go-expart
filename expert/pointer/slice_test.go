package main

import (
	"testing"
)

func Benchmark_SliceValue(b *testing.B) {
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
		f(v)
	}
}

func Benchmark_SlicePointer(b *testing.B) {
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
		_ = g(v)
	}
}
