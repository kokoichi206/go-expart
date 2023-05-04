package main

import "testing"

// just a struct with 3 string fields
type Hero struct {
	Name        string
	Description string
	Class       string
}

func NewHero() Hero {
	return Hero{
		Name:        "Hero",
		Description: "Hero description",
		Class:       "Hero class",
	}
}

//go:noinline
func ReturnHero() Hero {
	h := NewHero()
	return h
}

//go:noinline
func ReturnHeroPtr() *Hero {
	h := NewHero()
	return &h
}

//go:noinline
func returnHero(h Hero) Hero {
	return h
}

//go:noinline
func returnHeroPtr(h *Hero) *Hero {
	return h
}

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
