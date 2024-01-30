詳細 Go 言語アプリケーション開発を読ませてもらってます。
ありがとうございます。

## context.Context

context.Context 型の値には関数のロジックに関わる値を含めてはいけない。
アプリケーションを実装する上で、認証・認可についてビジネスロジックの中で検証しないようにするのは、ビジネスロジックと認証・認可ロジックを疎にしておくための 1 つのアプローチ方法。

context.TODO 関数

## [database/sql](https://pkg.go.dev/database/sql)

- **sql.Open を使うのは一度だけ**
  - sql パッケージには DB へのコネクションをプールする機能がある
  - Xxx メソッドと XxxContext が存在するときは XxxContext メソッドを利用する！
- sql.ErrNoRows は単一のレコードを取得する時のみ発生しうる
- トランザクションを使うときは defer で rollback を呼ぶ

## オブジェクト指向

- カプセル化（Encapsulation）
- 多態性（Polymorphism）
  - Japanese クラスが Person クラスの変数に代入できる
- 継承（Inheritance）
  - 埋め込みはオブジェクト指向で定義される**継承ではなく、コンポジションにすぎない**
    - 所有

[Go はオブジェクト指向言語か](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)
⇨ yes and no

[FAQ](https://go.dev/doc/faq)

### 継承

- Java では 2 つの継承がある
  - 実装継承
    - カプセル化を破壊する危険も大きい
    - 深い継承構造は、クラスの構成把握を困難にする
  - インターフェース継承
    - Go サポート

## インターフェース

Go は具象的にインターフェース名を記載しない暗黙的インターフェース実装を採用している。
→ 一見不便に見えるが、疎結合で柔軟な設計を可能にする！

「-er」インターフェース。
インターフェース定義に Xxx メソッドを 1 つだけ持つインターフェースは慣例として Xxxer という名前にする（Doer）。

- 契約による設計。
  - C++ 等では assert
  - Go ではコードコメント
- API インターフェースのコメントに何を書くべきか → 13.5 interface documentation
- 不要な抽象化を行うインターフェースの定義はコードの可読性を低下させる

## エラー

- エラーはただの値である！
- エラー作成方法 2 つ
  - メッセージに %s, %d といった Verb を使って動的な情報を埋め込むときは fmt.Errorf
  - errors.New
- fmt.Errorf を利用すれば別の error 型の値を使って新たな error 型の値も生成できる
  - チェーンをしてエラーが発生した一連の流れを表現できる
    - スタックトレースを含まないため
  - スタックトレースは「ひとまず出しておく情報」という面が強く、冗長なことが多い
    - プログラマーがエラー設計ときちんと向き合うべきだ〜〜〜っていう考えにより、
- もとのエラーオブジェクトの情報を保持するにはエラーを埋め込む **`%w`** を利用する
  - エラーチェーンに特定のオブジェクトが含まれているかを見るには、`errors.Is` 関数を使う
  - エラーのラップが漏れていないかのチェックとして、静的解析ツールも存在する
  - エラーオブジェクトを比較するときは == ではなく errors.Is を使う
- errors.As を使った独自情報の取得
  - err 引数のメソッドチェーンの中に target 引数で指定された方があれば true を戻す
- 独自エラー宣言
  - 慣習として Err- というプレフィックスをつける
    - `ErrNotExist = errors.New("xxx")`
  - 独自エラーを定義する場合でも関数やメソッドの戻り値は error インターフェースを使う
    - その型情報まで必要かどうかは呼び出し側の都合

## 関数

- 関数に状態を持たせるときは構造体を用意してメソッドを作成する
  - 構造体自体に意味がないならばクロージャの作成ですむ！

## 環境変数

- 環境変数を読み込む操作はシステムコールを呼ぶ操作なのでコストが高い
- os.Getenv
- os.LookupEnv

## DIP

- Go はシンプルであることが言語思考にあるため、高度な DI ツールはあんまり使われてない印象
- 他言語で DI ツールが重宝されるのは次の理由
  - DLL ファイルや Jar ファイルから実行時に動的にクラスをロードできる
  - フレームワークや UI に依存した実装を避けたい
- DB に依存させたくないなどの理由がある場合など、DI の利用は適切に用法量量を守って！！
  - インターフェースを使った過剰な抽象化には注意が必要
    - 暗黙的インターフェース
- **下位の実装の詳細が上位概念の中小へ依存していること**

## ミドルウェアパターン

- 共通処理
  - 多数のエンドポイントで同じ処理を行いたい！
  - オブザーバビリィティツールの対応
  - アクセスログの処理
- 具体例
  - リカバリーミドルウェア
    - テストをいくら書いても、配列操作などで panic を発生させてまう可能性はゼロにはできない
  - アクセスログミドルウェア
  - リクエストボディをログに残すミドルウェア
    - \*http.Request のリクエストボディはストリームのデータ構造であり、一度しか読み取れない！

## memo

- 現在は特別な理由がない限り Go Modules を使う。
  - `go get -u`: 依存パッケージの更新
  - `go get -u example.com/pkg`: 特定のパッケージの更新
  - go mod tidy は、go.mod 修正後は必ず行う
- GO111MODULE 環境変数によって挙動が変わるが、Go 1.16 からはデフォルトで ON
- 依存先のコードにデバッグコードを差し込む
  - go mod vendor
  - go.mod ファイルに replace ディレクティブを記述する
  - Workspace モードを使う（Go 1.18 で追加！）！
- テストが書きにくいコードは書き直しの対象になるだけではなく、凝集度が低いことや結合度が高い疑惑がある！
  - テストコードを書こうとしてはじめてわかる設計のまずさが存在する！
- [Twelve-Factor App](https://12factor.net/ja/)
- [O'Reilly Online Learning](https://www.oreilly.com/online-learning/individuals.html)
- **迷ったらシンプルを選ぶ**
- 構造体のフィールドに『context.Context』インタフェースを含めない！
  - **対象とするスコープが曖昧になるため**
  - [Contexts and structs](https://go.dev/blog/context-and-structs)
- github でコードジャンプできるのは [navigation code on github](https://docs.github.com/ja/repositories/working-with-files/using-files/navigating-code-on-github) のおかげ！
- [セマンティックバージョニング 2.0.0](https://semver.org/lang/ja/) に従う！
  - 一回ちゃんと読んでみる
- FAQ:
  - https://go.dev/doc/faq#Is_Go_an_object-oriented_language
- **SOLID**
  - **インタフェース分離の原則**
    - **クライアントに、クライアントが利用しないメソッドへの依存を強制してはならない！**
  - 1 つのみのメソッドを持つインタフェースは『~~er』という名前にする
- エラーの原因がわからなくなるのは情報隠蔽というより情報欠落。。。
- `t.Setenv` は `t.Parallel` との併用ができない
- テストを書くことは
  - **API の使いやすさを最初の利用者として確認できる**

## TODO

- TODO vs Background
  - [差分はない！](https://github.com/golang/go/blob/f719d5cffdb8298eff7a5ef533fe95290e8c869c/src/context/context.go#L195-L205)
    - じゃあどう使い分ける？
  - [テストでは Background でいいらしい](https://github.com/golang/go/blob/f719d5cffdb8298eff7a5ef533fe95290e8c869c/src/context/context.go#L209)
- Go には何があればオブジェクト指向になる？
  - [From FAQ](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)
- Go らしさっていっぱいあると思うんですが、どこが一番？
- [pkg.go.dev](https://pkg.go.dev/std) を読む
- Go が Go で描かれてる、の意味があんまりわかってない
  - 一番最初ってどうやってできてるん？
    - 流石にアセンブリ言語とか、でできてる？
    - もしくは C とか？
- internal パッケージを使いこなしたい
  - パッケージの1つ上とその階層以下ならアクセス可能
- デバッグしてみる！
  - go mod vendor
  - go.mod -> replace
  - Workspace
- go の migration ツールに馴染む
  - sqldef
    - https://github.com/sqldef/sqldef
- ファクトリー関数では実装とインタフェースのどっちを返すべき？
  - クライアントが利用しないメソッドへの依存を強制させないって意味で、構造体を返すべきなのかな？
    - store の KVS でも構造体を返している
  - **インタフェースを返すべき時ってのはどんな時？**
    - 本を読んだけどよくわからず、
    - 契約による設計？
    - github-sdk みたいな、ある API のラッパー的なのを作ってる、それはインタフェースでもいい？
- Go の情報として追ってるものはありますか？
  - 自分はリリース情報の rss とか
    - https://googlegroups2rss.ferdypruis.dev/rss/golang-announce.xml.rss
- やってみたいこと
  - Go scheduler
    - https://morsmachine.dk/go-scheduler
- `go:embed` を使えるようになる！
  - シングルバイナリで実行可能のまま保てる
- インタフェースについて
  - クライアント側に利用するインタフェースを絞らせるべきではないってことで、
  - usecase, repository にそれぞれインタフェースつけてる
    - クライアント側じゃなくて、実装側のパッケージにインタフェース作ってる
    - usecase, repository がそれぞれ1つだから助かってる気がする？
    - サービスが細かく分かれてるから、そうしていた
- 標準パッケージもしくはコードリーディング
  - net/http -> port を bind するってところとか serve の時に 1 req ごとに accept していく部分をみた
    - 大雑把な流れは掴めた
    - context が複雑だったような
  - json tag がどう解釈されるかっていうところで json パッケージ見たりとか
    - プールしてる部分がいまいち掴めなかったイメージ
  - 興味あるところ
    - 最近 k8s の勉強してるんで、そことかも気になってはいる
      - service とか dns の周りとか
