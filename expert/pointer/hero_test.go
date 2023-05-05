package main

import "testing"

func BenchmarkReturnStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ReturnHero()
	}
}

func BenchmarkReturnPointer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ReturnHeroPtr()
	}
}

func BenchmarkReturnStrucet(b *testing.B) {
	b.ReportAllocs()
	h := Hero{
		Name:        "Hero",
		Description: "Hero description",
		Class:       "Hero class",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = returnHero(h)
	}
}

func BenchmarkReturnPoineter(b *testing.B) {
	b.ReportAllocs()
	h := Hero{
		Name:        "Hero",
		Description: "Hero description",
		Class:       "Hero class",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = returnHeroPtr(&h)
	}
}
