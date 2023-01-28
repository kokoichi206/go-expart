## 環境

- VSCode
  - Restart Language Server
  - Generate Interface Stubs
  - Fill Struct
  - Add Tags To Struct Fields
  - Go To Definition
  - Generate Unit Tests For Function
  - Toggle Test File

## ランタイム

2 つ後のメジャーバージョンがリリースされるまでがサポート対象！

```sh
# 既存 Go モジュールのバージョンを上げる方法
go mod tidy -go=1.17
```

## lint

```sh
go vet
```

### golangcli-lint

```sh
go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
```

- golangcli-lint, errcheck とかあるのいいね！
- reviewdog と組み合わせると、新規に発生した差分に対してのみ Linter を実行できる！

## build

### バイナリサイズを小さくする

`-ldflags` を付与し `-s -w` などとするだけで、バイナリサイズを小さくできる！

``` sh
# 3.8M
go build main.go

# 2.9M
go build -ldflags '-s -w -X main.version=1.0.0' main.go
## -s -w はシンボルテーブルを省略する引数
## go tool objdump main でバイナリから逆アセンブルできる
```
