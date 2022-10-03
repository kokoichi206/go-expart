## [サーバーオプションの選択](https://firebase.google.com/docs/cloud-messaging/server?hl=ja&authuser=0#choosing-a-server-option)

Firebase Admin SDK。Node、Java、Python、C#、Go をサポートします。

FCM サーバーとの対話には Firebase Admin SDK または基本的なプロトコルの 2 つの方法があり、どちらを使用するか決定する必要があります。一般的なプログラミング言語で広くサポートされていること、認証や認可を簡単に処理できることから、Firebase Admin SDK を使用することをおすすめします。

## 後はリポジトリを参考にする

https://github.com/firebase/firebase-admin-go

ドキュメント: https://pkg.go.dev/firebase.google.com/go/messaging
snippets: https://github.com/firebase/firebase-admin-go/blob/master/snippets/messaging.go

### [サーバーに Firebase Admin SDK を追加する](https://firebase.google.com/docs/admin/setup/#go)

``` sh
# Install as a module dependency
$ go get firebase.google.com/go/v4

# Install to $GOPATH
$ go get firebase.google.com/go

export GOOGLE_APPLICATION_CREDENTIALS="path_to.json"
```

[サービスアカウント用の秘密鍵ファイル（json）を取得](https://firebase.google.com/docs/admin/setup/#initialize-sdk)

#### 環境変数をセットする場合

``` go
app, err := firebase.NewApp(context.Background(), nil)
```

#### 環境変数をセットしない場合

``` go
import (
	option "google.golang.org/api/option"
)
opt := option.WithCredentialsFile("path_to.json")
app, err := firebase.NewApp(context.Background(), nil, opt)
```
