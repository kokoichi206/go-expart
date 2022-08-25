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

[protobuf と gRPC](<https://docs.wantedly.dev/fields/system/apis#:~:text=protobuf%20(Protocol%20Buffers)%20%E3%81%AF%E3%83%87%E3%83%BC%E3%82%BF,%E5%BD%B9%E5%89%B2%E3%82%92%E7%BD%AE%E3%81%8D%E6%8F%9B%E3%81%88%E3%82%8B%E3%82%82%E3%81%AE%E3%81%A7%E3%81%99%E3%80%82>)

> protobuf (Protocol Buffers) はデータフォーマットで、JSON の役割を置き換えるものです。 一方 gRPC は通信プロトコルで、HTTP の役割を置き換えるものです。

gRPC ではプロトコルバッファを使って API を定義し、メッセージをシリアライズしている。

```sh
wget https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-osx-aarch_64.zip

sudo unzip protoc-21.5-osx-aarch_64.zip -d /usr/local/protobuf

protoc --version
> libprotoc 3.21.4
```

[Go-standards](https://github.com/golang-standards/project-layout) によると、`api` ディレクトリに protobuf を置くこと、となっている。

## ログパッケージの作成

- ログは分散サービスを構築する上で最も重要なツールキット！！
- ログを構築することで様々なことを学べる

### ログは強力なツール

変更内容をジャーナルに記録（ext）したり、WAL と呼ばれるログに記録（PostgreSQL）したりしている。
データベースの複製、分散サービスの連携、フロントエンドのアプリケーションの状態管理に、ログは役立っている。

完全なログは、最新の状態だけでなく、過去の全ての状態を保持している。

### ログの仕組み

ログは、追加専用のレコード列。ログは、レコードを常に時間順に並べ、オフセットと作成時間で各レコードにインデックスをつけるテーブルのようなもの。

ディスクサイズは有限なので、ログをセグメントに分割するなどが必要（**ログローテーション**）。

ストアファイルとインデックスファイルからなる（DB の感じか）。インデックスファイルは十分に小さいので、メモリへマップして高速化できる。

下から順に、ストアファイル・インデックス、セグメント、ログ。
ログという言葉は、レコード、レコードを保存するファイル、セグメントをまとめる抽象データ型という、少なくとも３つの異なるものを指す。

### 定義

- レコード: ログに保存されるデータ
- ストア: レコードを保存するファイル
- インデックス: インデックスエントリを保存するファイル
- セグメント: ストアとインデックスをまとめているものの抽象的概念
- ログ: セグメントをすべてまとめているものの抽象的概念

## gRPC

- gRPC は protobuf と HTTP/2 という強固な基盤の上に構築
  - protobuf はシリアライズに優れた性能
  - HTTP/2 は、gRPC が利用する長く維持されるコネクションを提供

```sh
go get google.golang.org/grpc
#?? go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

- internal パッケージとは、隣接するディレクトリのコードからしかインポートできない Go の魔法のパッケージ

### そのほか

- インタフェースによる依存性逆転
