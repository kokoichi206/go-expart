package main

import (
	"fmt"
	"iter"
	"slices"
	"unique"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

func intern() {
	ps := []*Person{
		{"Alice", 20},
		{"Bob", 21},
		{"Alice", 20},
	}
	for i, p := range ps {
		m := unique.Make(p)
		ps[i] = m.Value()
	}
	for _, p := range ps {
		fmt.Println(unsafe.Pointer(p))
	}
}

func itr() {
	for x := range 4 {
		fmt.Println(x)
	}

	items := []string{"a", "b", "c"}
	for _, v := range slices.All(items) {
		fmt.Println(v)
	}

	for v := range double([]int{1, 2, 3, 4}) {
		fmt.Println(v)
	}

	for k, v := range double2([]int{1, 2, 3, 4}) {
		fmt.Println(k, v)
	}

	base := slices.Values([]int{1, 2, 3, 4})

	isEven := func(v int) bool {
		return v%2 == 0
	}
	for v := range Filter(base, isEven) {
		fmt.Println(v)
	}
	d := func(v int) int {
		return v * 2
	}
	for v := range Map(base, d) {
		fmt.Println(v)
	}
}

func double(ns []int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for _, v := range ns {
			res := v * 2
			if !yield(res) {
				break
			}
		}
	}
}

// iter.Seq2
func double2(ns []int) func(yield func(string, int) bool) {
	return func(yield func(string, int) bool) {
		for i, v := range ns {
			if !yield(fmt.Sprintf("%d", i), v*2) {
				break
			}
		}
	}
}

func Filter[T any](delegate func(func(T) bool), condition func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range delegate {
			if condition(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func Map[T, V any](delegate func(func(T) bool), f func(T) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range delegate {
			if !yield(f(v)) {
				break
			}
		}
	}
}

func main() {
	intern()

	itr()
}
