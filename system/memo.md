## sec 1

### はじめに
システムプログラミングとは？
さまざまな定義

- C言語によるプログラミング
- アセンブリ言語を意識したC言語によるプログラミング
- 言語処理系の開発
- OS自身のプログラミング
- OSの提供する機能を使ったプログラミング

OS の機能とは、おおよそ次のもの

- メモリの管理
- プロセスの管理
- プロセス間通信
- ファイルシステム
- ネットワーク
- ユーザー管理
- タイマー

### WHY Go for system programming
C言語の性能とPythonの書きやすさ/読みやすさを両立させ、モダンな言語の特徴をうまく取り入れた言語となることを目標に開発されたもの。

Go言語は多くのOSの機能を直接扱えて、なおかつ少ない行数で動くアプリケーションが作れる！

### デバッガーを使ってシステムコールを「見る」
Goでは処理系に全環境用のソースコードも全てバンドルされている。そのため、デバッガーで処理を追いかけていくだけで、OSの機能を直接呼び出すコードまで簡単に見ることができる！

デバッガーは実行中のプログラムの全ての情報にアクセスできてしまうので、悪意のあるユーザーが下手に実行できないように、初回起動時は認証がある。


``` sh
go get github.com/derekparker/delve/cmd/dlv
```

``` go
func (f *File) write(b []byte) (n int, err error) {
	n, err = f.pfd.Write(b)
	runtime.KeepAlive(f)
	return n, err
}


// Write implements io.Writer.
func (fd *FD) Write(p []byte) (int, error) {
	if err := fd.writeLock(); err != nil {
		return 0, err
	}
	defer fd.writeUnlock()
	if err := fd.pd.prepareWrite(fd.isFile); err != nil {
		return 0, err
	}
	var nn int
	for {
		max := len(p)
		if fd.IsStream && max-nn > maxRW {
			max = nn + maxRW
		}
		n, err := ignoringEINTRIO(syscall.Write, fd.Sysfd, p[nn:max])
		if n > 0 {
			nn += n
		}
		if nn == len(p) {
			return nn, err
		}
		if err == syscall.EAGAIN && fd.pd.pollable() {
			if err = fd.pd.waitWrite(fd.isFile); err == nil {
				continue
			}
		}
		if err != nil {
			return nn, err
		}
		if n == 0 {
			return nn, io.ErrUnexpectedEOF
		}
	}
}
```


## sec 2
Go 言語がOS直上の低レイヤーを扱いやすくするために提供している抽象化レイヤーを紹介する。

- io.Writer
  - 出力の抽象化
- io.Reader
  - 入力の抽象化
- channel
  - 通知の抽象化

### io.Writer
syscall.Write() の呼び出しを確認したが、OSではこのシステムコールを、**ファイルディスクリプタ**といわれるものに対して呼ぶ。

ファイルディスクリプタは一種の識別子（数値）であり、システムコール呼び出し時に数値に対応するものにアクセスできる。

ファイルディスクリプタは、OSがカーネルのレイヤーで用意している抽象化の仕組み。

OSはプロセスが起動されるとまず３つの擬似ファイルを作成し、それぞれにファイルディスクリプタを割り当てる。0が標準入力、1が標準出力、2が標準エラー出力。以降は、そのプロセスでファイルをオープンしたりするたびに、一ずつ大きな数が割り当てられる。

POSIX系OSでは、可能な限りさまざまなものが「ファイル」として抽象化されている！

### Interface
Go では Interface に対して、`~er, ~or` のような名前にすることが多い。（Java では ~able）

func と メソッドのシグネチャの定義の間にレシーバーを置くと、構造体にメソッドを定義したことになる！

副作用のあるメソッドではレシーバーの方をポインタ型にする！

「*File は io.Writer インタフェースを満たす」といえる。という表現をする。

### 書かれた内容を記憶しておくバッファ: bytes.Buffer
Write() で書き込まれた内容を淡々と溜めておいて後でまとめて結果を受け取れる機能もある。

それが bytes.Buffer

1.10 からは strings.Builder も同様の役割をするものとして導入されている。

### インターネットアクセスの送信
io.Writer のような抽象化は他の言語でも実装がある。

### interface の実装状況・利用状況を調べる
```
godoc -http ":6060" -analysis type
```

ただし、GOPATH にパッケージがたくさん入っている場合は、解析に時間がかかり過ぎてしまう。


## sec 3

### io.Reader
``` go
type Reader interface {
  func Read(p []byte) (n int, err error)
}
```

引数である p は、読み込んだ内容を一時的に入れておくバッファ。あらかじめメモリを用意しておいて、それを使う。

``` go
// 1024 バイトのバッファを make で作る
buffer := make([]byte, 1024)
// size は実際に読み込んだバイト数、err はエラー
size, err := r.Read(buffer)
```

読み込み処理の方がめんどくさい。

### io.Reader の補助関数
Python, Ruby 等では、補助的なメソッドもファイルオブジェクトが持ってたりするが、Go言語では特別なもの以外はこのような外部のヘルパー関数を使って実現する。

#### コピーの補助関数
io.Reader から io.Writer にそのままデータを渡したい時！ファイルを開いてそのまま HTTP で転送したいとか、ハッシュ値を計算したいとか、いろいろなケースで使える

``` go
writeSize, err := io.Copy(writer, reader)
writeSize, err := io.CopyN(writer, reader, size)
```

### エンディアン変換
主流のCPUはリトルエンディアンであり、小さい桁数からメモリに格納される（0x2710 -> 10,27,0,0）。

しかし、ネットワーク上で転送されるデータの多くは、大きい桁からメモリに格納されるビックエンディアンであるため、修正がが必要となる！

### png ファイルを見てみる
PNGファイルはバイナリフォーマット。先頭の 8byte がシグニチャとなっている。

### format print
Go言語は方情報をデータが持っているため、全て「%v」と書いておけば変数の型を読み取って変換してくれる。






