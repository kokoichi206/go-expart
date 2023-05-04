package main

import (
	"testing"
)

type pien struct {
	paon   string
	hoge   string
	taihen string
}

//go:noinline
func f(v []pien) []pien {
	return v
}

//go:noinline
func g(v []*pien) []*pien {
	return v
}

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
		g(v)
	}
}
