## クラウドのストレージ

## AWS SDK

- https://github.com/aws/aws-sdk-go-v2/

```sh
# 必要なものだけ get する必要があるっぽい
go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3

go get github.com/aws/aws-sdk-go-v2/service/dynamodb
```

- [s3 examples](https://github.com/aws/aws-sdk-go-v2/tree/main/example/service/s3)

## Go CDK

```
go get "gocloud.dev/blob"
go get "gocloud.dev/blob/s3blob"
```

- [gocloud.dev blob](https://gocloud.dev/howto/blob/)

## DynamoDB

- サーバーレスアプリケーションは RDB とあまり相性が良くない
  - NoSQL が採用されがち
  - https://www.suzu6.net/posts/36/#aws-lambda%E3%81%A8db%E3%81%AE%E7%9B%B8%E6%80%A7%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6
- DynamoDB
  - フルマネージド
    - サーバーの用意が不要
  - 高可用性
    - AWS 側で 3つの AZ に保存される
  - 低レイテンシー
- DynamoDB
  - パーティションキー
    - 同じパーティションキーを持つアイテムは格納できない
  - パーティションキー + ソートキー
- 設計？
  - 常に明示的にパーティションキー・ソートキーを指定する必要がある
  - テーブルへのアクセスパターンを洗い出してキーを設計する必要がある！
  - https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/best-practices.html
- コントロールプレーン
  - テーブル作成などの管理 API
- データプレーン
  - テーブルのデータを CRUD する API

``` sh
# モックサービス
docker run -it -p 4566:4566 -e SERVICES=dynamodb localstack/localstack:0.13.2

export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
```

