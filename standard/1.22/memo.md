## 変更

- ループ変数のスコープ
  - エスケープされるような場合はいてレーションごとに変数を作成
- range over integer
  - 61405
  - range over func
    - iterator
    - 1.23~
- Vet
- Runtime
  - 型ベース GC のメタデータヒープオブジェクト近くに配置
- Compiler
  - PGO
    - 2-14% の改善が見られる
- cmp.Or(os.Getenv)
  - [61372](https://github.com/golang/go/issues/61372)

## Serve Mux

- Method の指定が可能に
- Path から値を取得できるようになった
- chi との違い
  - chi
    - ルーティングの完全一致
  - http
    - 前方一致
    - `{$}` とかをつけると、完全一致にできる
    - `/a/{$}`
- ルーティングのパターンマッチ
- `/a/` に登録して `/a` にリクエストをとばす → 301
- `/a` に登録し, middleware でリクエストの末尾 / を削除する
- 競合
  - 登録時に panic が発生する

## ループ変数

- 従来の for ループ変数は、イテレーションごとに使いまわされていた
  - コンパイラが適切なタイミングで、イテレーションごとに独立したループ変数を確保するように！
- 使い回しのメリット
  - メモリ割り当ての削減
    - ヒープメモリ
  - パフォーマンス向上
  - シンプルな実装
- 事例
  - goroutine での並行処理
  - テストの並行処理
  - ループ変数のポインタを利用した処理
- let's encrypt も踏んでしまってた
  - https://jovi0608.hatenablog.com/entry/2020/03/09/094737

## インライン展開

- Caller, Callee
- Heuristic improvements
- 1.21 以前
  - callee だけに着目
  - all or nothing
- Call site-aware heuristics
  - for 文内のループでクロージャを引数にとる高階関数の呼び出し
    - sort.Search のような

## vet

- ci とかでも入れてるらしい
  - golangci の中に入れてればよかったりするか
- defer の引数として渡したものは、即時に評価されてしまう
  - time.Since とかが絶対０秒になってしまう問題
- errors.As にポインタやエラーでない値を渡した時

## runtime

- [63340](https://github.com/golang/go/issues/63340)
- runtime/proc.go
  - Mark or Term 以外では Other として Metric が加算される
- runtime/trace
  - exp/trace
  - P に紐付けたバッファにイベントを書き込む
- https://github.com/golang/go/issues/60773

## archive

- zip, tar に AddFS が実装された
  - https://future-architect.github.io/articles/20240131a/
- `zw := zip.NewWriter(w)`, ResponseWriter
  - `zw.AddFS(contents)`

## Links

- [runtime/HACKING.md](https://github.com/golang/go/blob/master/src/runtime/HACKING.md)
- [Reducing Go Execution Tracer Overhead With Frame Pointer Unwinding](https://blog.felixge.de/reducing-gos-execution-tracer-overhead-with-frame-pointer-unwinding/)
