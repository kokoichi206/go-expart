## File

### Close

「書き込み」のような、元ファイルに影響を与えるような操作の場合、その操作が正常終了しないと Close できなかったりする。
→ Close メソッドのエラーハンドリングを、defer の中でやった方が良さげ

`syscall.Open` の時に `O_CLOEXEC` のフラグも渡しているので、`syscall.Close` は呼び出していない。



## Links

- [Go から学ぶ I/O](https://zenn.dev/hsaki/books/golang-io-package/viewer/intro)
