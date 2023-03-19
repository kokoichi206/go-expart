## sec 1

- go の目指すところ
  - Java や c++ のような静的型付け言語で、巨大なシステムでもスケールする
  - 動的型付け言語のように生産性が高く、可読性も高い
  - IDE を必要としない
  - ネットワーク処理が多重実行できる
- Go の活用場所
  - CLI
  - TUI
    - GUI のようなユーザーインターフェースを提供
  - Web アプリケーション
- いいところ
  - 標準で UTF-8 をサポート
  - マルチプラットフォーム
  - 並行処理の扱いやすさ
  - ストリーム指向！
    - io interface など
  - シングルバイナリ
    - 実行環境と開発環境で同じバイナリを使える
- Why Go?
  - コンパイルの速さ
  - レビューのしやすさ
    - gofmt
  - パッケージの公開の簡単さ！
  - libc 非依存！
    - c で開発された言語や、ライブラリが c で開発されている場合、それらは libc に依存している！
    - 実行ファイル生成に c が必要 → クロスコンパイルの難易度が上がる

## sec 2

環境構築

- バージョン管理ソフトウェア (goenv など)
  - 各ソフトウェアに対して必要なランタイムが固定されていないと動作しない or 動作の保証がないという問題を解決するには有効
  - Go のようにビルドされてすでに静的にリンクされた実行ファイルに対してはほぼ意味がない
  - 下位互換性: The Go 1 Compatibility Rules

## sec 3

- Go の特徴
  - GC がある
  - 静的な型を使う
  - ポインタを使う
  - 三項演算子はない
  - 継承はない
  - 例外はない
- 命名規則
  - 名前の長さはその情報の中身を越えるべきではない
  - グローバルな名前は相対的により多くの情報を伝えなければならない
  - 全てを名前で伝える！
- レイアウト
  - 実行ファイルを提供する場合は cmd 配下に実行ファイル名のディレクトリ + main.go

## sec 4

``` sh
$ python
>>> import this
The Zen of Python, by Tim Peters

Beautiful is better than ugly.
Explicit is better than implicit.
Simple is better than complex.
Complex is better than complicated.
Flat is better than nested.
Sparse is better than dense.
Readability counts.
Special cases aren't special enough to break the rules.
Although practicality beats purity.
Errors should never pass silently.
Unless explicitly silenced.
In the face of ambiguity, refuse the temptation to guess.
There should be one-- and preferably only one --obvious way to do it.
Although that way may not be obvious at first unless you're Dutch.
Now is better than never.
Although never is often better than *right* now.
If the implementation is hard to explain, it's a bad idea.
If the implementation is easy to explain, it may be a good idea.
Namespaces are one honking great idea -- let's do more of those!
```

- Zen of Python のうち、以下の2つを重視している
  - Explicit is better than implicit
  - Simple is better than complex
- blank import
  - db のドライバの register など

``` sh
go install golang.org/x/tools/cmd/stringer@latest
```

- NDJSON: Newline Delimited JSON
- path vs filepath
  - path: URL などの仮想的なパスの操作
  - filepath: 物理的なパス
- filepath.Walk よりも、filepath.WalkDir の方がパフォーマンスが良くなる！
  - filepath.Walk は検索した各エントリに対して os.Lstat を呼び出す実装のため
- context package
  - cancel, timeout
- build constraints
- cgo: go から c 言語を扱うための仕組み！
- デフォルト引数
  - 可変個引数
  - Functional Options Pattern
- internal パッケージ
  - 意味はわかる
  - 使ってる例がわからない
  - ライブラリ的なのを作らないと分からないか
    - 複雑なライブラリ作って、その中でパッケージを切るけど、公開するのはトップのやつ（handler 的な部分）のみ、って感じかな
- 埋め込み
  - 委譲としての継承を実現してる！

``` sh
# モジュールを指定して run することもできる！
go run github.com/mattn/longcat@v0.0.4
```

### メモ

- 移譲と継承について
  - 委譲ってのは別クラスに切り出してメンバー変数で持つだけ、のものかな
    - わざわざ名前つけんなや

## sec 5

- routing
  - func.Handle, func.HandleFunc は引数の違いのみで、キャスト可能
- Go のメソッド呼び出しは、第一引数にレシーバを持った関数の呼び出しと同義！！

### SMTP

- gmail: smtp.gmail.com
  - https://support.google.com/a/answer/176600?hl=ja
- 2 段階認証プロセス > アプリ パスワード
