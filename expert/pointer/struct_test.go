package main

import (
	"testing"
)

//go:noinline
func sf(v pien) pien {
	return v
}

//go:noinline
func sg(v *pien) *pien {
	return v
}

func Benchmark_SingleValue(b *testing.B) {
	b.ReportAllocs()
	v := pien{
		paon:   "hogehogehogehoge",
		hoge:   "fugafugafugafuga",
		taihen: "taihen taihenda",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sf(v)
	}
}

func Benchmark_SinglePointer(b *testing.B) {
	b.ReportAllocs()
	v := pien{
		paon:   "hogehogehogehoge",
		hoge:   "fugafugafugafuga",
		taihen: "taihen taihenda",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sg(&v)
	}
}
