## zero value

未初期化変数の利用とバッファオーバーランはセキュリティホールの原因の中でもトップクラスの危険度。
Go ではどちらもないようになってる！

## スライス

配列を参照する窓のようなデータ構造。他の言語では「ビュー」と呼ばれることがある！

## マップ

``` go
hs := map[int]string{
  200: "OK",
  400: "Bad Request",
}

lang := make(map[string][]string)
lang["Go"] = []string{"a", "b"}

l := lang["Go"]

// あるかどうかも一緒に取得
l, ok := lang["Go"]
```

## 制御構文

- 他の言語にある falsy という概念はない
  - ものによっては条件文中で true, false として扱われるもの


## Links

- The Go Blog
  - https://go.dev/blog/
- Effective Go
  - https://go.dev/doc/effective_go
- Go Wiki
  - https://github.com/golang/go/wiki
- go-java-tutorial
  - https://yourbasic.org/golang/go-java-tutorial/
- A tour of Go
  - https://go-tour-jp.appspot.com/welcome/1
- Go by Example
  - https://gobyexample.com/

## メモ

- プログラムを書くとは名前をつけることです
- シャドーイング
- Go の関数は一級市民
  - 関数の引数に渡したり、変数に入れることができる
