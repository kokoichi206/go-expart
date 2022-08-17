```sh
go run cmd/server/main.go

curl -X POST localhost:8080 -d '{"record": {"value": "TGFE3X0PKzIL"}}'
```

## プロトコルバッファ

非公開の API, 自分達でクライアントを開発するプロジェクトにおいては、JSON より生産性が高く、速く、多くの機能を持ち、バグの少ないサービスを作ることができるなら、それに越したことがない！（人間にとっての読みやすさは度外視できる）
それが、プロトコルバッファ（Protocol Buffers: protobuf）

- 型の安全性
- スキーマ違反を防ぐ
- 高速なシリアライズ
- 後方互換生

protobuf は２つのシステム間（マイクロサービスなど）の通信に適する。
Google は gRPC を構築した際に protobuf を使った。

### WHY プロトコルバッファ

- 一貫性のあるスキーマ
- バージョン管理
  - フィールドのバージョン管理が可能
- ボイラープレートコードの削減
- 拡張性
- 言語寛容性
- パフォーマンス

### プロトコルバッファのコンパイル

ある言語へコンパイルするには、その言語のランタイムが必要となる。
Go には protobuf を Go コードにコンパイルするための２つのランタイムがある。

```sh
# protobuf ランタイムをインストール
go get google.golang.org/protobuf/...@v1.28.0

# protobuf をコンパイル
protoc api/v1/*.proto \
  --go_out=. \
  --go_opt=paths=source_relative \
  --proto_path=.
```

## gRPC

gRPC ではプロトコルバッファを使って API を定義し、メッセージをシリアライズしている。

```sh
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-osx-aarch_64.zip

sudo unzip protoc-21.5-osx-aarch_64.zip -d /usr/local/protobuf

protoc --version
> libprotoc 3.21.4
```

[Go-standards](https://github.com/golang-standards/project-layout) によると、`api` ディレクトリに protobuf を置くこと、となっている。
