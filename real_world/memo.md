## sec 0
- [Microsoft REST API Guidelines](https://github.com/Microsoft/api-guidelines/blob/master/Guidelines.md)

## sec 1

### History
- RFCはIETFという組織が中心となって維持管理している、通信の相互接続性を維持するための共通化された仕様書集
- 通信プロトコルではなく、ファイルタイプの種類などの共通情報はIANAが管理

### HTTP0.9
0.9は現行のプロトコルと後方互換性がないため、まともには0.9のプロトコルを扱うことはできない

curl 

``` sh
curl --http1.0 http://localhost:18888/greeting
curl --http1.0 --get --data-urlencode "search world" http://localhost:18888
nslookup
```

### to 1.0
``` sh
curl -v --http1.0 http://localhost:18888/lgtm
curl --http1.0 -H "X-Test: Hi" http://localhost:18888/lgtm
```

0.9と比べて、以下が追加

- リクエスト
  - メソッド
  - HTTPバージョン
  - ヘッダー（Host, User-Agent, Accept）
- レスポンス
  - HTTPバージョン
  - 3桁のステータス
  - リクエストと同じ形式のヘッダー

### HTTPの先祖は電子メール！

HTTPの例
- リクエストヘッダー
  - User-Agent
  - Referer
  - Authorization
- Content-Type
  - MIMEタイプと呼ばれる識別子
    - MIMEタイプは電子メールのために作られた識別子
  - Content-Length
  - Content-Encoding
  - Date

ヘッダー（等？）の正規化の方法は特に指定されておらず、言語やフレームワークによっってことなる。

### MIMEタイプ（マイムタイプ）
ファイルの種類を区別するための文字列で、電子メールのために作られた。初出は1992。

Windowsは主にファイルの拡張子、Macはリソースフォークと呼ばれるメタ情報でファイルの種類を区別している。

『大項目/詳細』の形をしている。text/plain, image/jpeg などなど。

MIMEタイプのRFCはその後にも色々追加や更新があり、英語以外の多言語対応もこの延長線上に定義されている。image/svg+xml

意味が定義されていない単なるバイト列を表すMIMEタイプはapplication/octet-stream

IEはインターネットオプションによって、MIMEタイプではなく中身でファイルの種類を推測しようとする！この動作は Content Sniffing と呼ばれる。text/plain のつもりで送ったのに、HTMLとjsが書かれたために、ブラウザがそれを実行してしまうことがあった。
こうならないように、次のヘッダーを送信することで、ブラウザに推測を行わないよう指示できる

X-Content-Type-Options: nosniff

### HTTPの先祖はニュースグループ！
メソッドとステータスの２つの機能を導入。

curlでメソッドは、--request, or -X という形式を使う

``` sh
curl --http1.0 -X POST http://localhost:18888/lgtm
curl --http1.0 -X HEAD http://localhost:18888/lgtm
curl --http1.0 --head http://localhost:18888/lgtm
curl -v --http1.0 --head http://localhost:18888/lgtm
```

ステータスは、先頭３文字の数値を見てクライアントが動作を変更できるようにするべきもの！

### リダイレクト
Permanently 系のステータスでは、キャッシュをするようにしている。理にかなってそう。

リダイレクトはブラウザとの協調動作！サーバーはステータスとLocationヘッダーを返す。ブラウザはそれを見て、もう一度、ヘッダーで指定されたURLにリクエストし直している。

curlに-Lを付与すると、レスポンスが300番台でかつ、レスポンスヘッダーにLocationヘッダーがあった場合、そのヘッダーで指定されたURLに再度リクエストを送信する。

POSTでリダイレクトが返ってきた場合、GETで送信し直すケースもある、ということを覚えておくか。

1.1以降は、クライアントがリダイレクト無限ループを検知しなければならない。Go言語でのデフォルト設定でリダイレクトは10回に制限。

Googleのガイドラインでは、リダイレクトは5回以下、できれば3回以下、となっている。

アクセス先がディレクトリで末尾のスラッシュがない場合のリダイレクト、というものもある。

### URL (Uniform Resource Locators)
URIにはURN(Uniform Resource Name)という、名前の付け方のルールも含まれる。URLは住所、URNは名前そのもの。

Webを扱う限りURNが登場することはあまりないので、URLとURIはほぼ同一。その後の [RFC3305](https://datatracker.ietf.org/doc/html/rfc3305) でURLは慣用表現で、公式表記はURIということになったが、URLの方が一般的に広く使われている。

AWSのリソース名はURNを模して作られている。

URL
『スキーム://ホスト名/パス』

『スキーム://ユーザー:パスワード@ホスト名:ポート/パス?クエリー#フラグメント』

スキームを解釈するのはブラウザの仕事。ホスト名は大文字小文字を区別しない！へー。

フラグメントは、ページ内リンク先のアンカーを指定するもの。サーバーには送信されない。

### 正規URL
ほとんどのウェブサイトは www. を省略しても同じ内容が表示される。

Protocol-Relative URL と言って、スキームを省略して表記することもできる（例えば //example.com/image.png）。なるべく使わない方が良い。

### ボディ
0.9ではレスポンスはコンテンツそのものだったが、1.0ではリクエストとレスポンスにヘッダーが含まれるようになったため、ボディとヘッダーをきちんと分ける必要がある。

**ヘッダーとの間に空行を挟んで、それ以降が全てボディになる！**この構造は電子メールと全く同じ。

curl で送信時にボディを添付留守には、-d オプションを使う（--data）。

``` sh
# -d, --data, --data-ascii すでにエスケープされてる前提
# --data-urlencode, curl にエスケープしてもらう
# --data-binary, -T, -d @
curl --http1.0 -d "{\"hello\": \"world\"}" -H "Content-Type: application/json" http://localhost:18888/lgtm

```

リクエストボディはどのメソッドでも使えるが、推奨されていないメソッド（GETなど）もある。URLの文字数制限（約2000文字）を嫌ってあえて受け付けるように実装することもできるが、基本は避けるべき。









## curl
``` sh
curl -v --http1.0 http://localhost:18888/lgtm
curl --http1.0 -H "X-Test: Hi" http://localhost:18888/lgtm
curl --http1.0 -H "User-Agent: Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)" http://localhost:18888/lgtm

curl --http1.0 -X HEAD http://localhost:18888/lgtm
curl --http1.0 --head http://localhost:18888/lgtm

curl -L http://localhost:18888
```
