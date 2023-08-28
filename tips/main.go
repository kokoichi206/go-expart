package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getKeys[C comparable, V any](m map[C]V) []C {
	var keys []C
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// ts の union 的なものも定義できるのか。
type customConstraint interface {
	~int | ~string
}

// ~int は int を基底型とする型の集合を表す。
type customIntBase interface {
	~int
	String() string
}

type customInt int

func (i customInt) String() string {
	return strconv.Itoa(int(i))
}

// データ構造にもジェネリクスを使える。
type Node[T any] struct {
	Val  T
	next *Node[T]
}

// 型レシーバがインスタンス化されている？
// ジェネリクスは、使う前に普通の関数・型にする必要があり、
// インスタンス化とは、それぞれの型パラメータに具体的な型引数(type argument)を代入すること。
func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})

	// sort.Interface を実装している。
	aa := sort.StringSlice{"c", "a", "b"}
	fmt.Printf("aa.Len(): %v\n", aa.Len())

	fmt.Printf("sort.IsSorted(aa): %v\n", sort.IsSorted(aa))
	sort.Slice(aa, func(i, j int) bool {
		return aa[i] < aa[j]
	})
	fmt.Printf("sort.IsSorted(aa): %v\n", sort.IsSorted(aa))

	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}
	// これがインスタンス化。明示しなくても暗黙のうちに使われている。
	// keys := getKeys[string](m) と同じ
	keys := getKeys(m)
	_ = keys
}
