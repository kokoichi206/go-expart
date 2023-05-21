``` sh
cat /dev/urandom | LC_ALL=C tr -dc '0-9a-zA-Z' | head -c 100000000 > large_text

$ ls -lh large_text
-rw-r--r--  1 kokoichi  staff    95M May 21 15:00 large_text
```

test

``` sh
$ go test -bench .
goos: darwin
goarch: amd64
pkg: speed
cpu: VirtualApple @ 2.50GHz

Benchmark_CopyToFile-8             68529             15009 ns/op
Benchmark_CopyToFile2-8           279252              4283 ns/op
Benchmark_CopyToBuf-8            1309831               814.2 ns/op
Benchmark_ReadAll-8              1106350               958.1 ns/op
PASS
ok      speed   12.062s
```
