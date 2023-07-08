# slog の確認

## Benchamrk

``` sh
$ go test -bench .
goos: linux
goarch: arm64
pkg: slog-bench
BenchmarkMarshal-4       3041424               394.4 ns/op
BenchmarkSlog-4          2085474               557.6 ns/op
BenchmarkZap-4           4228914               288.5 ns/op
PASS
ok      slog-bench      4.879s
```

slog と zap の比較については、ほぼ [zap の README](https://github.com/uber-go/zap#performance) に記載の通りになった。

## output format

``` sh
{"time":"2023-07-08T15:29:21Z","level":"DEBUG","message":"ぱおん"}
{"time":"2023-07-08T15:29:21.42899763Z","level":"DEBUG","msg":"pien; tukareta"}
{"level":"info","ts":1688830161.429072,"msg":"zap test"}
```

## その他思ったこと

- JSONHandler にすると、slog の time Format が変更できない？
