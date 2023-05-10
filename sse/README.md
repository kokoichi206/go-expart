## Server Sent Event: SSE

- サーバーからクライアントへの一方通行
- HTTP 上のプロトコル
  - gzip 圧縮での転送が可能
- テキストのみ
  - バイナリデータは Base64 エンコードして送る
- HTTP の chunk を利用
  - response の長さを不定にする

## Links

- mdm
  - https://developer.mozilla.org/ja/docs/Web/API/Server-sent_events/Using_server-sent_events
- gin-sse
  - https://github.com/gin-gonic/examples/blob/master/server-sent-event/main.go
- sse と ws の比較記事
  - https://qiita.com/suin/items/e33af700ceb678d40a67
- sse: RFC 6455
  - https://www.rfc-editor.org/rfc/rfc6455
- 標準パッケージで sse
  - https://qiita.com/taqm/items/e132a1aa55690a22b655
