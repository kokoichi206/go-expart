## Proxy

- http client > DefaultTransport > Proxy フィールドで簡単に設定が可能
  - デフォルトでは ProxyFromEnvironment が使われる
    - `HTTP_PROXY`, `HTTPS_PROXY` 等の環境変数から読まれる

## 疑問

- ws コネクションではどうしたらいい？
