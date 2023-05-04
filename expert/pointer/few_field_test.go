package main

import (
	"testing"
)

type fuga struct {
	hoge   string
	pien   string
}

//go:noinline
func fff(v fuga) fuga {
	return v
}

//go:noinline
func ggg(v *fuga) *fuga {
	return v
}

//go:noinline
func ffff(v []fuga) []fuga {
	return v
}

//go:noinline
func gggg(v []*fuga) []*fuga {
	return v
}

func Benchmark_FewFieldStructValue(b *testing.B) {
	b.ReportAllocs()
	v := fuga{
		hoge: "hogehogehogehoge",
		pien: "pienpienpienpien",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fff(v)
	}
}

func Benchmark_FewFieldStructPointer(b *testing.B) {
	b.ReportAllocs()
	v := fuga{
		hoge: "hogehogehogehoge",
		pien: "pienpienpienpien",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ggg(&v)
	}
}

func Benchmark_FewFieldSliceValue(b *testing.B) {
	b.ReportAllocs()
	size := 1000
	v := make([]fuga, size)
	for i := 0; i < size; i++ {
		v[i] = fuga{
			hoge: "hogehogehogehoge",
			pien: "pienpienpienpien",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ffff(v)
	}
}

func Benchmark_FewFieldSlicePointer(b *testing.B) {
	b.ReportAllocs()
	size := 1000
	v := make([]*fuga, size)
	for i := 0; i < size; i++ {
		v[i] = &fuga{
			hoge: "hogehogehogehoge",
			pien: "pienpienpienpien",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gggg(v)
	}
}
