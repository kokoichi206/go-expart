## Editor Tips

VSCode では、Fill Struct をすることで、構造体のフィールドを埋めることができる！

自分の環境では『Ctrl + Shift + P > Fill Struct』はなぜか動かなかったが、構造体にカーソルを合わせた時に出る**黄色いマーク**から埋めることができた。

![](./imgs/fill_struct.png)

## 環境

### 環境変数

クラウド時代になり、プログラムの挙動を制御する方法として環境変数が注目を集めている。

同一のバイナリを、本番環境やテスト環境など、個別の設定で動かす場合に活用できる。

## パフォーマンス

### スライスのメモリ確保を高速化する

Go のランタイム内部で OS とのやりとりを減らすため、確保済みのメモリを使い回すよう管理を行なっている、らしい！

**スライスの裏には固定長の配列があることを意識する**

```go
// 正確な長さが分かっている場合
sl := make([]int, 1000)
fmt.Println(len(sl))
fmt.Println(cap(sl))

// 正確な長さがわからないが、最大量の見込みがつく場合
// キャパシティだけ増やす？
sl2 := make([]int, 0, 1000)
fmt.Println(len(sl2))
fmt.Println(cap(sl2))
```

「**有効なデータが入っている要素数**と**キャパシティ**」

### マップのメモリ確保を高速化する

マップの裏には、バケットというデータ構造が複数ある。

1 つのバケットには 8 つのデータを保持でき、要素の数が増えるほどバケット数が増える。

**キャパシティのみ。**

```go
m := make(map[string]string, 1000)
// 0
fmt.Println(len(m))
```

## 日時

time.Time で**日時・タイムゾーンの全てを扱う！**

```go
now := time.Now()

tz, _ := time.LoadLocation("America/Los_angeles")
future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

fmt.Println(now.String())

past := tme.Date(1955, time.November, 12, 6, 3, 2, 1, time.UTC)
```

サマータイムの開始日時なども含まれ、これらは頻繁に更新されている！

### フロントに渡す

JS とやりとりをすることを考えると、**time.RFC3339Nano** がおすすめらしい。

### 翌月の計算

日付が範囲外になった時、Go はその値を存在する日付に変換する（正規化）！

例えば、6/31 を指定した場合、7/1 に自動的に変換される。

ここで、5/31 の 1 ヶ月後として次のようにしたとき、7 月が表示されることに注意。

```go
jst, _ := time.LoadLocation("Asia/Tokyo")
t := time.Date(2022, 5, 31, 00, 00, 00, 000, jst)
nextMonth := t.AddDate(0, 1, 0)
fmt.Println(nextMonth)

// 回避策？
y1, m1, d := t.Date()
first := time.Date(y1, m1, 1, 0, 0, 0, 0, time.UTC)
y2, m2, _ := first.AddDate(0, 1, 0).Date()
nextMonthTime := time.Date(y2, m2, d, 0, 0, 0, 0, time.UTC)
if m2 != nextMonthTime.Month() {
    // 翌月末
    return first.AddDate(0, 2, -1)
}
return nextMonthTime
```

## What is Go

- 多くの割り切り
- go fmt のオプションがないことで、誰が書いても似たようになる
- go modules というパッケージ管理で、ツール選定で悩まないように
- メタプログラミングを使って動的に挙動が変わる余地がほとんどない
  - なるべく静的に解決
-
