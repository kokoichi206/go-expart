package main

import (
	"errors"
	"fmt"
	"io"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
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

type Foo struct{}

func (f *Foo) Bar() string {
	return "bar"
}

func Paon(f *Foo) string {
	return "paon"
}

func (f *Foo) Error() string {
	return "error"
}

func validate() error {
	var f *Foo
	return f
}

func deferError() (err error) {
	defer func() {
		closeErr := fmt.Errorf("close error")
		fmt.Printf("err: %v\n", err)
		err = errors.Join(err, closeErr)
	}()

	return fmt.Errorf("return error")
}

// ワーカープールパターン。
// 送信
//
//	for {
//	b := make([]byte, 1024)
//		// r から b へ読み込む！
//		// 読み込むごとにチャネルに新たなタスクを発行する。
//		ch <- b
//	}
func read(r io.Reader) (int, error) {
	task := func(b []byte) int {
		return len(b)
	}

	var count int64
	wg := sync.WaitGroup{}
	// ゴルーチンプール。
	n := 10

	// プールと同じ容量のチャネルを作る。
	ch := make(chan []byte, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			// 共有チャネルからタスクを受信する！
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	close(ch)
	wg.Wait()
	return int(count), nil
}

func main() {
	maskTest()

	jsonTest()

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

	var foo *Foo
	// メソッドは第一引数をレシーバとする関数のシンタックスシュガーにすぎない？
	fmt.Println(foo.Bar())
	// メソッド式？
	// fmt.Println(Foo.Paon)
	if err := validate(); err != nil {
		// nil ポインタもエラーになる。
		fmt.Println(err)
	}

	fmt.Printf("deferError(): %v\n", deferError())
	fmt.Printf("hooooop ----- \n%s\n", deferError())

	// GOMAXPROCS の値を更新する。
	// 0 だと現在の値を返す。
	numGOMAXPROCS := runtime.GOMAXPROCS(0)
	fmt.Printf("numGOMAXPROCS: %v\n", numGOMAXPROCS)
}
