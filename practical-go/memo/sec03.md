## 構造体

オブジェクト指向のクラスから**一歩シンプルにした**構造体が主役となる

## 特徴

- フィールドをまとめて持てる
- 他の構造体を埋め込める
  - 親クラスの継承のようにも見えるが、アップキャスト、ダウンキャストはできない
- 関数の中や式の中でも、動的に構造体を定義できる
  - ?

## レシーバー

- 値型にすると、レシーバーの属性を書き換えても変更は保存されない！
- **このメソッドは状態を変更しない**という情報は重要なので、変更の有無によって使い分けるのが大事

## 関数

Go において、関数は一級市民である。

## タグ

構造体のタグは、Java などのアノテーションに相当している。

```
フィールド名 型 `json:"field"`
```

入出力をマッピングする多くの処理系では、CamelCase と snake_case の自動変換等をサポートしてくれている。しかし、**実際にはコードを探索するためのキーとなり得るため、極力名前の変換は省略せずにコードとして表現するのが良い！**

暗黙の変換ルールを想像するのはなかなか難しいものがあるので。。。

```go
// 以下のように短縮表記を使うこともできる！（Go 1.16~）
type User struct {
  Name string `json xml:"name"`
}
```

## 設計のポイント

- ポインター型でのみ扱うかどうか
  - 内部にスライスや map, ポインターなど参照型の要素を持っている場合には、基本的にポインター型でのみ扱う構造体にする。
- 値として扱える場合
  - インスタンス全体がコピーされることになる
  - ポインターの場合と異なり、**代入したり引数として渡すたびにコピーされる**
  - **インスタンスを作り、その関数のライフサイクルでのみ消費される場合、スタックメモリ上にインスタンスが確保される！**
    - メモリ割り当てのコストがほぼないし、GC の仕事も減る
    - パフォーマンス改善どころ

Go の場合、エンティティと呼ばれるような構造体はミュータブルのほうが良さそう。

## 空の構造体

- 並行処理のゴルーチン間での最低限の情報のやり取り
- map を他の言語の集合として扱う
- 何もフィールドがないメソッド集の実装

## メモリ割り当ての高速化

`sync.Pool` などで、OS にメモリをリクエストする回数を減らすことが可能。

## 構造体について補足・オブジェクト指向と対比して

Go の埋め込みは継承ではなく、その逆。埋め込んだ構造体が子供としてぶら下がる。単なるメンバーであり、埋め込まれた構造体は自分のことしか知らない。

Is-A **ではなく** Has-A の関係である。

## メモ

- Go はプログラマーの成長、という点に重きを置いてると何回も述べられてる
  - フレームワーク重視の言語設計との対比
  - main 関数から辿れる

## TODO

- reflect に慣れる
- noCopy の部分理解する
- ストラテジーパターン