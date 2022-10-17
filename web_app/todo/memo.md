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

## mock tool

```sh
go get -u github.com/DATA-DOG/go-sqlmock
```

https://github.com/DATA-DOG/go-sqlmock

## 疎結合へ！

責務を複数の実装に分割。
インターフェースを挟むことで、他のパッケージの実装内容に影響しないテストコードが書ける。

handler パッケージからは、ビジネスロジックと永続化に関わる処理を取り除く

## generate でモックの自動生成

go generate
https://qiita.com/yaegashi/items/d1fd9f7d0c75b2bb7446

[moq 生成ツール](https://github.com/matryer/moq)
[このツールの思想: Meet Moq: Easily mock interfaces in Go](https://medium.com/@matryer/meet-moq-easily-mock-interfaces-in-go-476444187d10)

```
go get -u github.com/matryer/moq
```

## user

```sh
curl -X GET localhost:18000/health
curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden

# なぜこれで通らず、" をエスケープしたもので通るのかを調べる！！！！
## だめ
curl -X POST localhost:18000/register -d '{"name": "john doe", "password": "test", "role": "user"}'
curl -X POST localhost:18000/register -d '{"name":"john doe", "password":"test", "role":"user"}'
curl: (52) Empty reply from server

## おけ ？
curl -X POST localhost:18000/register -d '{\"name\":\"john doe\", \"password\":\"test\", \"role\":\"user\"}'
curl -X POST localhost:18000/register -d "{\'name\':\'john doe\', \'password\':\'test\', \'role\':\'user\'}"
curl -X POST localhost:18000/register -d "{'name':'john doe', 'password':'test', 'role':'user'}"

curl -i -XPOST localhost:18000/register -d @./handler/testdata/add_user/ok_req.json.golden

```

## Redis

Key value 型の**インメモリデータベース**

アクセストークンは有効期限が切れるとともに無効化すべき一時的なデータベースであるため、RDBMS を使った永続化は行わない。  
（誰がいつ発行した、とかいう情報はいらんっけ、いらんか）

スケールアウトで複数のコンテナが稼働している可能性や、そもそも数分前と同じ仮想サーバーが稼働していない可能性があるクラウドネイティブなアプリケーションにおいて、リクエストを処理する API サーバーがアクセストークンを払い出した同じ API サーバーだという前提を置いたらだめ！

**仮想サーバーやコンテナはステートレスである必要がある**ため、一時的なデータでも Redis などを利用して**ミドルウェア上で保管して共有**する！
なんでこれで解決できるんだっけ！！！！？？？？

```sh
# 開発環境の docker を停止
make down
# 開発環境の docker を開始
make up

make logs
```

## [Medis](https://github.com/luin/medis)

Redis の GUI クライアント

```sh
# リリースページから最新版をダウンロード
# https://github.com/luin/medis/releases

# ダウンロードのページに書いてあるように進める
## インストールで失敗
npm install
```

## go-redis

redis version によって違いそう！？
https://github.com/go-redis/redis#installation

```sh
go get github.com/go-redis/redis/v8
```

## JWT

JWT をアクセストークンとして利用する。

アクセストークンには、秘密鍵と公開鍵を利用した RS256 形式の署名にする！

```
brew install openssl
echo 'export PATH="/opt/homebrew/opt/openssl@3/bin:$PATH"' >> ~/.zshrc
openssl version
OpenSSL 3.0.5 5 Jul 2022 (Library: OpenSSL 3.0.5 5 Jul 2022)
```

秘密鍵と公開鍵の発行

```
openssl genrsa 4096 > secret.pem
openssl rsa -pubout < secret.pem > public.pem
```

JWT の生成には [lestrrat-go/jwx](https://github.com/lestrrat-go/jwx)。

```
go get github.com/lestrrat-go/jwx/v2
go get github.com/google/uuid
```

Json Web Token とは、Base64 URL エンコードされた JSON を使って、二者間で情報をやり取りするための仕様！！

[RFC7519](https://tex2e.github.io/rfc-translater/html/rfc7519.html) で定義。

JWT と署名・暗号化に関わる関連仕様をまとめたものは、JOSE（JSON Object Signing and Encryption）と呼ばれたりする。

jwt.io: https://jwt.io/
を使うと、正しい署名か確認してくれる。

### go:embed

ファイルパスを使う方法だと、シングルバイナリで実行可能、というメリットが使えなくなる。

そこで、`go:embed` ディレクティブを使い実行バイナリにファイルを埋め込む。

https://pkg.go.dev/embed

## TODO

疑問等

handler/service と service/interface の役割の違い（？）が分かってない
