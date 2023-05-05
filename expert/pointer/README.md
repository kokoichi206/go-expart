## 結果

``` sh
$ go test -bench .
goos: darwin
goarch: amd64
pkg: pointer
cpu: VirtualApple @ 2.50GHz
# few_field_test.go
Benchmark_FewFieldStructValue-8         558563112                2.166 ns/op           0 B/op          0 allocs/op
Benchmark_FewFieldStructPointer-8       570834327                2.170 ns/op           0 B/op          0 allocs/op
Benchmark_FewFieldSliceValue-8          544026063                2.235 ns/op           0 B/op          0 allocs/op
Benchmark_FewFieldSlicePointer-8        539479299                2.542 ns/op           0 B/op          0 allocs/op
# hero_test.go
BenchmarkReturnStruct-8                 100000000               10.57 ns/op            0 B/op          0 allocs/op
BenchmarkReturnPointer-8                40040058                30.44 ns/op           48 B/op          1 allocs/op
BenchmarkReturnStrucet-8                128752497                9.598 ns/op           0 B/op          0 allocs/op
BenchmarkReturnPoineter-8               574554602                2.107 ns/op           0 B/op          0 allocs/op
# sample_test.go
Benchmark_ValueValue-8                  388464440                3.202 ns/op           0 B/op          0 allocs/op
Benchmark_ValuePointer-8                553840508                2.211 ns/op           0 B/op          0 allocs/op
# slice_loop_test.go
Benchmark_SliceValueLoop-8                  4312            249520 ns/op           48015 B/op       1000 allocs/op
Benchmark_SlicePointerLoop-8                3913            306516 ns/op           48013 B/op       1000 allocs/op
# slice_test.go
Benchmark_SliceValue-8                  551618252                2.148 ns/op           0 B/op          0 allocs/op
Benchmark_SlicePointer-8                583519108                2.093 ns/op           0 B/op          0 allocs/op
# struct_test.go
Benchmark_SingleValue-8                 100000000               10.29 ns/op            0 B/op          0 allocs/op
Benchmark_SinglePointer-8               564624066                2.149 ns/op           0 B/op          0 allocs/op
PASS
ok      pointer 24.263s
```

## 考察

- slice は、その容量の大部分を占める arary の部分がポインタ
  - 値でもポインターでもコピーのコスト変わらなさそう
    - `Benchmark_SliceValue-8` と `Benchmark_SlicePointer-8`
- field の数によって、なんか処理が変わってる？
  - field の数が 3 つ以上とかだと、顕著に struct, pointer で差が出た
  - field の数が 2 つの時は、struct も pointer も差がわからなかった
 

## 疑問

[hero_test.go](./hero_test.go) は[こちらの記事](https://levelup.gitconnected.com/a-guide-to-benchmarking-in-go-22adeb40bea6)を参考に実施したが、結果は以下のようになった

``` sh
BenchmarkReturnStruct-8                 100000000               10.57 ns/op            0 B/op          0 allocs/op
BenchmarkReturnPointer-8                40040058                30.44 ns/op           48 B/op          1 allocs/op
BenchmarkReturnStrucet-8                128752497                9.598 ns/op           0 B/op          0 allocs/op
BenchmarkReturnPoineter-8               574554602                2.107 ns/op           0 B/op          0 allocs/op
```

BenchmarkReturnStruct-8 と BenchmarkReturnPointer-8 の結果が逆（struct の方が実行時間が大きくなる）かと思ったが、何でこうなるんだろう。。。
（なんでここだけ MemAllocsPerOp, AllocatePerOp が 0 じゃない？）

該当部分は ↓

``` go
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

func BenchmarkReturnStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReturnHero()
	}
}

func BenchmarkReturnPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReturnHeroPtr()
	}
}
```

### コンパイル時の最適化結果の確認

``` sh
go build -gcflags=-m *.go
# command-line-arguments
./hero.go:12:6: can inline NewHero
./hero.go:20:6: can inline NewHeroPtr
./hero.go:30:14: inlining call to NewHero
./hero.go:36:17: inlining call to NewHeroPtr
./hero.go:51:16: inlining call to testing.(*B).ReportAllocs
./hero.go:58:16: inlining call to testing.(*B).ReportAllocs
./hero.go:65:16: inlining call to testing.(*B).ReportAllocs
./hero.go:79:16: inlining call to testing.(*B).ReportAllocs
./main.go:3:6: can inline main
./hero.go:21:9: &Hero{...} escapes to heap
./hero.go:36:17: &Hero{...} escapes to heap
./hero.go:41:17: leaking param: h to result ~r0 level=0
./hero.go:46:20: leaking param: h to result ~r0 level=0
./hero.go:50:28: b does not escape
./hero.go:57:29: b does not escape
./hero.go:64:29: b does not escape
./hero.go:78:30: b does not escape
```

### メモ（正しいかは微妙？）

実行結果が逆になる理由は、Goのコンパイラの挙動による。
このケースでは ReturnHeroPtr 関数がポインタを返すので、メモリアロケーションが行われている。Go コンパイラは, ReturnHeroPtrWithNewHero() 関数内で構造体のポインタが返されることを検出し、**構造体(値型)をスタックではなくヒープに割り当てる**ことを決定している(値型も場合によっては heap に詰まれる？)。  
**スタック上のデータは関数の実行が終了すると破棄される**ため、関数からポインタを返すと、無効なメモリ領域を指すことになるため。ヒープ上に割り当てることで、関数のスコープを超えてデータが生き続けることができる。このような場合、Goコンパイラは自動的にヒープへの割り当てを行います (したがって、**値型の変数であっても、関数のスコープを超えて使用される場合は、ヒープに移動させられることがある?**)。

一方, NewHero 関数は構造体を返すため、スタック上に割り当てられ, ReturnHero 関数もスタック上で構造体を返すため、メモリアロケーションが行われません。

また、ベンチマーク結果が逆になる理由の一部は、Goランタイムがガーベジコレクションを行っているためです。ガーベジコレクションは、実行中にメモリを解放するプロセスで、プログラムの実行速度に影響を与えます。NewHeroPtr 関数は、**ヒープ上に新しいオブジェクトを割り当てるため、ガーベジコレクションの対象となり**ます。これに対して、NewHero 関数はスタック上に割り当てられるため、ガーベジコレクションの影響を受けない。

## Links

- https://levelup.gitconnected.com/a-guide-to-benchmarking-in-go-22adeb40bea6
- https://www.wakuwakubank.com/posts/811-go-benchmark/
- https://qiita.com/ryskiwt/items/574a07c6235735afa5d7

## memo

``` sh
go build -gcflags '-m -l' main.go struct_slice.go
```
