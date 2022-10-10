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

``` sh
# 以下のように target オプションを指定することで、deploy にリリース用のバイナリしか含まれなくなる
docker build -t kokoichi0206/todo:${DOCKER_TAG} --target deploy ./

docker compose build --no-cache
docker compose up

curl localhost:18000/hello
```
