Go カンファレンスでのメモです。

## Memory Model

- The go Programming Language Specification
  - https://go.dev/ref/spec
- The Go Memory Model
  - テスト実行に -race をつける
  - https://go.dev/ref/mem

### 並行処理の難しさ

- 起こりうる全ての順序を列挙すれば完全に分析する
  - 間違い！
  - 逐次一貫モデル
    - 演算を一列に並べて、その通りに実行した結果として説明できるはずだ
    - あらゆるコードで成り立つわけではない！
- 3 つに別れる
  - 並行性のないプログラム
  - 並行性があり、逐次一貫モデルで説明できるプログラム
  - 並行性があり、逐次一貫性で説明できないプログラム
    - 書くべきではない
    - どうしたらこうなるのか？
    - data race のあるプログラム

### メモリーモデル

- 目標
  - 観測可能性: メモリーモデルのゴール
  - happens-before 関係
  - 並行性
  - data race
- happens-before
  - 先に起こる、とは解釈しない！！
    - 「観測可能性」を定めるためのもの
  - 3つの可能性
    - a happens before b
      - a < b
    - b happens before a
      - b < a
    - どちらも成り立たない
      - a と b は並行（concurrent）
  - A から B に向かって到達できる時
    - A happens before B
- data race (演算 a, b)
  - a と b は並行
  - a と b の対象メモリ位置が重なっている
  - 少なくともどちらかが書き込み演算
  - どちらかが同期処理
    - atomic ではない
- 観測可能性
  - どちらかが成り立つ時、観測可能である！
    - r < w, r < w' < w となるような他の書き込み演算 w' が存在しない
    - r と w が並行

### 1.19

- 8年ぶりのメモリーモデルのアップデート！
  - sync atomic の仕様がわかるようになった！
- sync/atomic
  - sync/atomic だけで書かれた go プログラムは、逐次一貫モデルで説明できる！
    - 観測可能性モデルでは、happens before 関係を書き加える作業に対応する
    - 単にアトミックなだけではなく、同期処理的な一面も持つ


## 可読性

- Go Style Guide
  - https://google.github.io/styleguide/go/
  - Go によってサポートされてるわけではなく、Google のもの？
- Effective Go
  - https://go.dev/doc/effective_go
  - 更新はない
- Uber Go Style Guide
  - https://github.com/uber-go/guide/blob/master/style.md

### Style Guide

- Style Guide 基礎になるもの
- Decisions 特定のポイントについて
- Best Practice 具体的な例や補足

### memo

- `%q` を使う
- buf -> bufio
  - パッケージ名と変数名を被らせないための工夫
  - url -> urlpkg


## メモリ管理

- 現代のアプリケーションにおいて、CPU がボトルネクになることはない
- I/O がボトルネックになることが多い
- go メモリ管理
- プロセスが直接物理領域を触ることはない
  - プロセスからは**仮想メモリ**しか見えない！
- バイトスライスを書いた時に、どこのメモリ領域に用意してるの？
  - 基本的には、go に任せていればいい
- ヒープ
  - m ヒープ
    - 全ての go ランタイムで共有されている
  - arena
  - mcentral
    - mspan っていう双方向リスト
- スタック
  - goroutine のスタックはめっちゃ小さい、2kb
    - os のスタックは 8mb
    - os スレッドの中にザバザバって詰めていくことを go がしている
- ヒープがしばしば問題になる
  - gc とか
  - メモリのアロケートはコストが高すぎる

