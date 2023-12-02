## コンパイラ

- 独自のコンパイラ
  - gcc や llvm を使わない
- コンパイル型
  - CPU 命令を生成 → コード実行
  - cf: インタプリタ型
    - コード実行 → CPU 命令を生成
    - 仮想マシンが CPU/メモリへの負担となるらしい
    - Write Once, Run Everywhere
- C と Go の比較
  - バイナリサイズ
    - Go → 2 MB
  - Ojbdump でアセンブリを見る
    - Go → 14 万行... (hello world)
  - Go でシステムコールが多い理由
    - Runtime があるから
- [ref/spec](https://go.dev/ref/spec)
- 様々な Go コンパイラ
  - gc (go compiler)
  - gcc (gccgo)
  - llvm (llgo)
- コンパイラ
  - フロントエンド
    - ソースコード → AST
  - バックエンド
    - AST → 機械言語
    - [SST: Single Static Assignment](https://pkg.go.dev/golang.org/x/tools/go/ssa)
      - 低レベルな中間言語
      - 全ての変数が不可変

## CallGraph

- 関数どうしの呼び出し関係を表現した有向グラフ
- https://github.com/ondrajz/go-callvis
- NetworkX
  - https://networkx.org/documentation/stable/reference/generated/networkx.drawing.nx_pydot.read_dot.html

## parser

- 字句解析
  - bufio.Scanner
    - bufio.SplitFunc を書く
- ケース
  - 標準出力に流れてくるものに色付けをする
  - db に流れてくるパケットを解析し構造化ログを作る
    - 複数クライアントからくるため、バッファリングが必要
    - TCP の開始・終端
- bufio.Scanner
  - go-yaml/scanner
  - 読み込むバイト数の上限がわかってる時は便利
    - わからない時は bufio.Reader
- memo
  - https://github.com/prometheus/procfs

## 生成 AI と静的解析

- 静的解析
  - プログラムを**実行せずに**解析すること
  - lint, コード補完, format ...
- 作りやすい
  - go/analysis
- [GopherCon](https://www.gophercon.com/home)

## コンパイラのきもち

- シンタックスの暗記
  - ゼロサイズのメモリアロケーション回避
    - struct{}, bool?
  - キャスト最適化
    - string(b) とかをまとめてしまうと、余計にアロケーションが語る
- 知識があれば推論可能
  - ゼロクリア
    - 配列埋めは高速化が図られがち
    - clear 関数も最適化がかかる
      - 非ゼロ値以外は用途が限定的なため入ってない
  - インライン展開
    - [ghidra](https://ghidra-sre.org/)
    - `//go:noinline`
    - -l のビルドオプションで off にできる
      - めっちゃパフォーマンス変わる
  - エスケープ解析
    - 関数に閉じるのが明らかな変数はスタックに積まれる
    - 参照が関数に閉じない可能性があるオブジェクトはヒープにエスケープされる
    - **言語特性**の方が近い
  - マークスキャン回避
    - GC のスキャン
    - 内部にポインタがなければその分スキャンする必要がなくなる
