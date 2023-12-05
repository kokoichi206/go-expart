## スマホゲーム

- 1人プレイが主軸のスマホゲーム
  - 多くの木の王の主な処理が、データの更新処理
  - 1人のプレイユーザーのみ参照・更新するデータの操作で完結することが多い
- データの区分
  - マスターデータ
    - あまり更新されない
    - S3 などで管理
      - Google Sheets など
  - ユーザーデータ
    - クリア情報
    - 保持してるアイテム数など
- データ
  - 1人のユーザーから更新されるデータに関しては、排他ロックが不要
    - 他デバイスからのアクセスが見込まれるときは、ミドルウェア等でユーザ ID での制御とか
- サーバーの処理において、レスポンス遅延の原因として、データベースとサーバーの通信が比率を占めることが多い
  - AWS で AZ をまたぐ通信には数 ms の遅延が発生する

## ログ

- ログの活用
  - 不具合の問い合わせ
  - クエストやイベントでの、ログ分析を行い、レベルデザインや機能追加の意思決定に繋げる！
  - ランキング集計バッチなどで、途中結果の確認等
  - 誰がいつどの操作をおこなったか
- ログの種類
  - アプリケーションログ
  - アクセスログ
  - 行動ログ
    - 更新系の API で出力する
    - アイテム獲得・クエスト挑戦
    - トランザクション中のログ
      - context に格納しておき、トランザクション完了後に出力する
  - 操作ログ
    - マスターデータの更新・ゲーム内お知らせの配信

## イベント

- 複数のイベント・ミッション
  - 後から進捗判定が追加されることがある
  - 進捗管理のそれなりに重い処理 ＋ それなりに回数
- Pub/Sub パターン
  - publisher, subscriber がお互いのことを知らなくて良くなる
    - 疎結合！
  - 「〇〇という武器を強化したと発行する側」と「強化に関するイベントを購読する側」
  - context を利用して実現！
    - generics を利用して、型安全に！

## NG ワード

- スカンソープ問題
  - https://en.wikipedia.org/wiki/Scunthorpe_problem
  - NG: ABC
  - の時に ABCD もはじいちゃう問題

## 単体テスト

- assert.ElementMatch
  - 要素の順序が異なっていても、要素が等しければ判定は true になる
    - map, slice

## Links

- Colly: https://github.com/gocolly/colly
  - scraping
- connect-go: https://github.com/bufbuild/connect-go
  - grpc, grpc-web, connect 独自のプロトコルの3つに対応！
  - migrate from grpc: https://connect.build/docs/go/grpc-compatibility/
- https://github.com/grpc-ecosystem/go-grpc-middleware
