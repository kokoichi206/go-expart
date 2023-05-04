## sec 0

### コンポジット型

複数のデータを１つの集合として表す型

- 構造体: struct
- 配列: array
- スライス: slice
- マップ: map
- チャネル: channel, chan

### slice

[Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)

``` go
src := []int{1, 2, 3, 4, 5}

dst := make([]int, len(src))
copy(dst, src)
```

スライスの任意の要素を削除する


``` go
src := []int{1, 2, 3, 4, 5}

// ３番目の要素を削除する
i := 2
dst := append(src[:i], src[i+1:]...)
```

### エクスポート

アクセス修飾子はなく、パッケージ外からアクセスできるかできないかの二択。ユーザー定義型の型名や、フィールドの最初の文字が大文字かどうかで判断。

構造体の方をエクスポートして、フィールドをエクスポートしない例として、カウンターがある。

``` go
package syncutil

import "sync"

type Counter struct {
    Name String

    // エクスポートされないフィールドがある場合は、空行を入れることが多い
    // ミューテックスを利用する際は対象となるフィールドらの先頭で定義することが多い
    m       sync.RWMutex
    count   int
}

func (c *Counter) Increment() int {
    c.m.Lock()
    defer c.m.Unlock()
    c.count++
    return c.count
}

func (c *Count) View() int {
    c.m.RLock()
    defer c.m.RUnlock()
    return c.count
}

c := &syncutil.Counter{
    Name: "Access",
}

fmt.Println(c.Increment())
fmt.Println(c.View())
```

「型自体はパッケージ外で再利用を制限したいが、JSONへの変換やログへの出力を行いたい」というケースには、フィールドを公開するかどうかを検討する！

### pointer vs value

- https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
- https://go.dev/doc/effective_go#pointers_vs_values
- json.RawMessage は使い分けがきっちりしてるので、読んでみるといいかも

``` go
// RawMessage is a raw encoded JSON value.
// It implements Marshaler and Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.
type RawMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ Marshaler = (*RawMessage)(nil)
var _ Unmarshaler = (*RawMessage)(nil)
```

- english
  - pitapat
  - sup?

### メソッド

レシーバに紐づけられた関数のことをメソッド。式としても扱うことができる

### インタフェース
Goでは、インタフェースを使うことでのみ抽象化できる。抽象化とは、実装そのものを隠蔽し、振る舞いのみに注目する共通化を行う手法。

``` go
type Crier interface {
    Cry() string
}

type Duck struct{}

func (d Duck) Cry() string {
    return "Quack"
}

// Duck型はCryメソッドを実装しているのでCrierインタフェースを満たしている
var Crier = Duck{}
```

**関数にメソッドを実装し、インタフェースを満たす場合もある**（関数にメソッド。。。）

``` go
type ParrotFunc func() string

func (p ParrotFunc) Cry() string {
    return p()
}

var c Crier = ParrotFunc(func() string {
    return "Squawk"
})
```

空インタフェースは、満たすべきメソッドが存在しないインタフェースで、どんな型でも受け入れることができる。

### インタフェースの設計

**極力メソッドリストは少なくする！**

> The bigger the interface, the weaker the abstraction.

構造体の埋め込み

### 型キャスト・型アサーション

``` go
var i int32 = 189
var j int64

// compile error
j = i

// cast
j = int64(i)

i := interface{}("test sample")
// assertion
s := i.(string)
n, ok := i.([]byte)
```

### 型スイッチ

``` go
i := interface{}("Go Expert")

switch i.(type) {
case int, int8, int16, int32, int64:
    fmt.Println("This is integer,", i)
case string:
    fmt.Println("This is string,", i)
default:
    fmt.Printf("This is unknown type, %T\n", i)
}
```

### 入出力の標準化

io.Reader と io.Writer は、非常によく抽象化されたインタフェース

``` go
func IsPNG(r io.Reader) (bool, error) {
    magicnum := []byte{137, 80, 78, 71}
    buf := make([]byte, len(magicnum))
    _, err := io.ReadAtLeast(r, buf, len(buf))
    if err != nil {
        return false, err
    }
    return bytes.Equal(magicnum, buf), nil
}
```

io.Pipe 面白い

``` go
func Post(m *Message) (rerr error) {
    pr, pw := io.Pipe()
    // 読み書きは単一のゴルーチンで同時には行えない。
    go func() {
        defer pw.Close()
        enc := json.NewEncoder(pw)
        err := enc.Encode(m)
        if err != nil { rerr = err }
    }()
    const url = "http://example.com"
    const contentType = "application/json"
    _, err := http.Post(url, contentType, pr)
    if err != nil { return err }
    return nil
}
```

他にも io.Copy を用いても、すべてのデータをメモリ上に展開せずにコピーが行える。しかし、外部から渡されたデータをコピーする場合には**ZIP爆弾**に注意し、制限が付与された io.CopyN 関数を使うといい。

io.ReadAll 関数の多用は避けるべき。（せっかく **io.Reader 型がストリームで扱えるように設計されている**にもかかわらず、メモリ上にダンプして無駄にメモリを消費しているため）

### ファイルシステムの抽象化

- `fs.File`
- `fs.FS`
  - embed.FS

### 並行処理

> Do not communicate by sharing memory; instead, share memory by communicating

ミューテックスは、複数のゴルーチン間でメモリを共有した領域（クリティカルセクション）を保護し、操作の原子性を担保するための仕組み。

Goではゴルーチン間で通信を行い、メモリの内容を共有することで平行処理を行うアプローチをとっている。チャネルはこのアプローチで並行処理を行うため、**基本的にはミューテックスではなくチャネルの使用が推奨される**。

go という予約語を用いた go 文。メイン関数自身もゴルーチンとして呼び出されており、メインゴルーチンと呼ぶ。

#### select-default

default をうまく使うことでブロックされずにチャネルを使った送信/受信処理が欠けるようになる

``` go
ch := make(chan int)

select {
case ch <- 100:
    fmt.Println("sent")
default:
}

select {
case <-ch:
    fmt.Println("received")
default:
}
```

#### for-select

定期実行されるような処理を簡潔にかける。

``` go
for {
    select {
    case <-time.Tick(1 * time.Second):
        fmt.Println("waiting...")
    case <-doneCh:
        // doneCh に値が送信されるまで、上記のメッセージが流れる
    }
}
```

#### nilチャネル

ch1 が受信できたら、以降の ch1 からの受信を無効化する

``` go
select {
case <-ch1:
    ch1 = nil
case <-ch2:
    //
}
```

### close を使ったブロードキャスト

クローズされたチャネルを受信するとブロックされずにゼロ値が返る。

### context パッケージ

ゴルーチン間での実行状況、生存時間、値を共有できる

contextパッケージの核心は、Contextインタフェースにある。コンテキスト木と呼ばれるものがあり、基本的に親コンテキストをコピーし、付加情報を加えた小コンテキストを生成する。

- 冪等性がある！
  - 最初に成功する操作があれば、２回目以降の呼び出しは処理されず終了する
  - Err メソッドは必ず最初にキャンセルされたり湯を返す！
- context 使用時に守るべきルール
  - https://pkg.go.dev/context#pkg-overview
- cancelCtx
  - **チャネルを閉じることで、それを待っているすべてのゴルーチンにブロードキャストしている！**

### ポインタ

言語仕様としてポインタを隠蔽している言語もあるなか、Goでは制約こそあるものの、ポインタと実体を使い分けながら利用するようにしている。

#### スタックとピープ

**スタック**はメモリの使い方や使用量が**コンパイル時に決定できる**場合に用いられる。必要な分だけ確保し、関数から抜ける時に解放されるため、メモリを効率よく使える。一方、**ヒープ**はメモリの使い方や使用量が**実行時にしかわからない**場合に用いる。ヒープに確保した変数の生存期間は用途によってバラバラであるため、GoではGCを用いてヒープのメモリを集中管理している。GCのアルゴリズムには様々なものがあるが、Go では GC 時に STW（Stop The World）が発生するものを採用しているので、GC が動いている間はプログラムの実行が止まってしまう。パフォーマンスを機にするプログラムを開発する場合は CG による STW の影響は無視できないため、メモリを確保する先がスタックになるかヒープになるかは重要！

変数が実態で定義される場合、Goはその変数をスタック上に確保する。一方、変数の型をポインタにすると、メモリはヒープかスタックのどちらかに確保される。

`go build -gcflags "-m"' といったオプションをつけてビルドすると、エスケープ解析の結果を見ることができる。

#### 変数のコピーリスト

変数を実体にしていた場合は、そこに書かれているデータ自体のコピーになる！

レシーバはポインタとして定義していた方がパフォーマンスが良い

``` go
type Value struct {
    content [64]byte
}
// メソッドのレシーバをポインタにする
func (v *Value) Content() [64]byte {
    return v.content
}

// 関数の引数をポインタにする
func Content(v *Value) [64]byte {
    return v.content
}
```

### 構造体のフィールド

- time.Time 参考
  - 複数の変数から参照を共有されたくない場合 → 実体
  - 共有したい場合 → pointer
- 値が存在しない場合とゼロ地を区別したい場合 → pointer にすることも
  - int
- コピーコストの観点から、サイズの大きい方を実体で管理するよりは、ポインターで管理する

### unsafe.Pointer

- Go のポインタ型のキャストの制約を無視して、任意のポインタに変換するための仕組みを提供

### 型のメモリレイアウト

C のように文字列の最後に終端文字を配置せず、文字列がどこまで続いてるかは Len の値を参照する！

``` go
type String struct {
    Data untptr
    Len int
}
```

- slice: https://github.com/golang/go/blob/a82f69f60e976d1a99c477903f5de98839c24f70/src/runtime/slice.go#L15-L19
- map: https://github.com/golang/go/blob/a82f69f60e976d1a99c477903f5de98839c24f70/src/runtime/map.go#L117-L131

`var v map[int]struct{}` のように宣言すると、実際に `var v *runtime.hmap` のように hmap をポインタ付きで宣言したことになる！

つまり `map == *runtime.hmap` は**常にポインタでの扱いになる！**

### エラーハンドリング

``` go
type error interface {
    Error() string
}
```

- 一般に、エラーが発生するケースでは、その関数を完了できなかったことを意味する
  - → error 型以外の戻り値は利用しないことが強く推奨されている！

#### エラーの生成とハンドリング

``` go
func divide(x, y int) (int error) {
    if y == 0 {
        return 0, errors.New("divide by zero")
    }
    return x / y, nil
}
```

error インタフェースを満たした、独自のエラー型を生成することが可能

``` go
type PathError struct {
    Op      string
    Path    string
    Err     error
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

#### エラーのラッピング！

エラー発生箇所を特定しやすくするなどに役立つ

``` go
func handleSignupRequest(name string) error {
    // ~~~~
    if err := db.CreateUser(name); err != nil {
        return &Error {
            Op: "signup",
            err: err,
        }
    }
    
    return nil
}

type Error struct {
    Op string
    err error
}

func (e *Error) Error() string {
    return fmt.Sprintf("handle %s request: %s", e.Op, e.err.Error())
}
```

ラッピングを受け取った際に、db エラーなら 500 を返すようにしたい。しかし、そのままではラッピングされているので、型アサーションで見分けることができない。`errors.Unwrap`を用いて、ハンドリングする。

``` go
if err := handleSignupRequest(name); err != nil {
    if e, ok := errors.Unwrap(err).(*db.Error); ok {
        // status code 500 を返す処理
    }
}
func (e *Error) Unwrap() error {
    return e.err
}
```



## Tips??
- mapの初期化時の値をからの構造体にすることで、メモリのアロケーションをゼロにできる
    - `m := make(map[string]struct{})`
- *int のゼロ値は「nil」
    - ゼロ値という概念
    - JSON のパースで、null をどうするか
    - int のゼロ値は「0」
    - JSON のパースでは *int を使う
