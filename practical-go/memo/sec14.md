## クラウド

- クロスコンパイルでビルドも早い
- ランタイムも内蔵したシングルバイナリが生成可能
- コンテナとの親和性が高く、クラウドネイティブ言語
- Docker がコンテナを一躍有名にした

## Docker

- docker image pull
- docker image ls
- docker image rm
- docker image build -t 
- 差分でイメージを管理している
  - 100M * 10 ≠ 1G
- Docker のマルチプラットフォーム機能

## オーケストレーション

コンテナオーケストレーション

- Docker compose
  - dc build, up, down,,
  - dc up では build されるのは初回のみ
  - 自前のイメージファイルを使う場合には build を忘れない
- Kubernetes

docker-compose.yaml はプロジェクトメンバーが用意できる最上級のおもてなし

## Dockerfile

- マルチステージビルド
  - https://docs.docker.jp/develop/develop-images/multistage-build.html
- Proxy の情報は、実行時のコマンドライン引数を使って渡す
  - セキュリティ情報もビルドしたコンテナに残る
- ko
  - https://github.com/ko-build/ko
  - Go に特化したイメージ作成ツール
