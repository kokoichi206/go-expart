## REST API を構築していく

- RESE API での考え方は gRPC などでも使える
  - Middleware パターンは、gRPC では Interceptor
  - Graceful Shutdown
  - Exponential Backoff

## リファクタリングとテスト

1st ステップの問題点

- テスト完了後に終了する術がない
- 出力を検証しにくい
- 異常時に「os.Exit」
- ポート番号が固定されており、テスト起動に失敗する可能性がある

[「Go net/http タイムアウト」の完全ガイド](https://blog.cloudflare.com/ja-jp/the-complete-guide-to-golang-net-http-timeouts-ja-jp/)

## Docker

```sh
# 以下のように target オプションを指定することで、deploy にリリース用のバイナリしか含まれなくなる
docker build -t kokoichi0206/todo:${DOCKER_TAG} --target deploy ./

docker compose build --no-cache
docker compose up

curl localhost:18000/hello
```

## テスト

- t.Helper
- t.Skip
- t.Cleanup
- t.Parallel

テストの入力や期待値を別ファイルとして保存したテストのことを**ゴールデンテスト**と呼ぶ！
テストコードとは別に保存するデータは例えばデータなら `*.json.golden` というファイル名とかにする。

https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3

golden ファイルを json として認識させる

https://khigashigashi.hatenablog.com/entry/2019/04/27/150230

```sh
curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden
curl -i -XGET localhost:18000/tasks
```

## マイグレーションツール

標準パッケージや Go 自体に RDMBS のマイグレーションを管理する機能は提供されてないので、OSS を利用する。

```sh
go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
```
