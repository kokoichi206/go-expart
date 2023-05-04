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

## Links

- https://levelup.gitconnected.com/a-guide-to-benchmarking-in-go-22adeb40bea6
- https://www.wakuwakubank.com/posts/811-go-benchmark/
