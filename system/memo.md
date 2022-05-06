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


## sec 4
低レベルアクセスを抽象化する３つ目のGoの機能がチャネル。

### チャネル
キューは FIFO 型のデータ構造。Go言語のチャネルは、このキューに並行処理用の「並列でアクセスされても正しく処理される」ことを保証する機能を組み合わせたもの！

CSP（Communicating Sequential Processes）というモデルを、チャネルという形で実現したものになる！並行で動くコードにおいて、お互いのプロセスが同じデータを直接触るのではなく、コミュニケーションを行いつつ協調する構造にする！

- チャネルとは、データを順序よく受け渡すためのデータ構造である
- チャネルは、並列処理されても正しくデータを受け渡す同期機構である
- チャネルは、読み込み・書き込みで準備ができるまでブロックする機能である

チャネルは、クローズしなくてもガベージコレクタに回収される。

チャネルは単なるデータ構造ではなく、言語のコアに深く組み込まれた機能である。

for 文と組み合わせるのがよく使われる。

for ループのイテレーターに chan を使うと、チャネルがオープンしている間は回り続けるが、チャネルがクローズしたら止まるものとなる。
つまり、「値が来るたびに for ループが回る、個数が未定の動的配列」のように扱っていることになる（Python, JS のジェネレータのようなこと）。

### チャネルと select 文
``` go
for {
	select {
	case data := <-reader:
		// 読み込んだ data を利用
	case <-exit:
		// ループを抜ける
		break
	}
}
```

### コンテキスト
コンテキストは、深いネストの中、あるいは派生ジョブなどがあって複雑なロジックの中でも、正しく終了やキャンセル、タイムアウトが実装できるようにするための仕組み。

### システムからの通知
OSの仕事の中には、時間のかかるものや、いつ返ってくるかわからないものがいくつかある。

- サーバーのプロセスに、クライアントからつなげてくるのを待つ
- 巨大なファイルを読み込んで、読み込み完了まで待つ
- ユーザーがマウスをクリックするまで待つ
- 他のスレッドがロックを解除するまで待つ

これらの仕事を実現するために、システムの一番下のカーネルのレイヤーでは、大きく次の３つの方式が採用されている

- OSが何かを準備するとき、それを待っているプロセスがどれかを把握し、準備が終わるまではプロセスを止め、準備ができたらプロセスに処理を戻す
  - ファイルやソケットなどのブロッキング入力
- OSが何かを準備するとき、終わっていなくても即座に処理を返す。返すものが一部だけ準備できていたら、その一部のデータとまだ続きがあることを返す
  - ノンブロッキング入力
- プロセスが実行中であればプロセスを一時停止し、あらかじめ設定していたコールバック関数を呼ぶ
  - シグナル


## sec 5
システムコールとは「特権モードでOSの機能を呼ぶ」こと。

### CPU の動作モード
プロセスは自分のことだけに集中し、メモリや時間の管理はプロセスの外からOSが全て行う方式が主流。

実行してよいハードウェアとしての機能が、ソフトウェアの種類に応じて制限されており、それを動作モードによって区別している。サポートしている動作モードの種類はCPUによって異なるが、OSが動作する特権モードと、一般的なアプリケーションが動作するユーザーモードの2種類はある。

OSの機能も、アプリケーションの機能も、バイナリレベルで見れば同じようなアセンブリコードですが、CPUの動作モードが異なる。

### システムコールでモードの壁を越える
システムコールを介して、特権モードでのみ許されている機能をユーザーモードのアプリケーションから利用できるようにしている。

ん？それって、ユーザーモードに最初から許可してたらダメなん？特別な使い方のみ、にしたいってこと？ぽい。

mac の場合に、ファイルオープンの syscall が呼ばれる。

``` go
syscall.Open(name, flag|syscall.O_CLOEXEC, syscallMode(perm))

syscall(funcPC(libc_open_trampoline), uintptr(unsafe.Pointer(_p0)), uintptr(mode), uintptr(perm))

// ---- asm_darwin_amd64.s の中のコード
// func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno);
TEXT	·Syscall(SB),NOSPLIT,$0-56
	CALL	runtime·entersyscall(SB)
	MOVQ	a1+8(FP), DI
	MOVQ	a2+16(FP), SI
	MOVQ	a3+24(FP), DX
	MOVQ	trap+0(FP), AX	// syscall entry
	ADDQ	$0x2000000, AX
	SYSCALL
	JCC	ok
	MOVQ	$-1, r1+32(FP)
	MOVQ	$0, r2+40(FP)
	MOVQ	AX, err+48(FP)
	CALL	runtime·exitsyscall(SB)
	RET
  ...
```

`runtime·entersyscall`から`runtime·exitsyscall`の中に、`SYSCALL`という命令がある。

POSIX（Portable Operating System Interface）は、OS間で共通のシステムコールを決めることで、アプリケーションの移植性を高めるために作られたIEEE規格のこと。最終的にOSに仕事を頼むのはシステムコールだが、POSIXで定められているのはそのインタフェースが決められている！

Go で syscall 以下の関数を使う場合、Go言語のドキュメントにはほとんどない。

### システムコールの内側
[linux のソースコード](https://github.com/torvalds/linux)の中の、[fs/read_write.c](https://github.com/torvalds/linux/blob/master/fs/read_write.c#L322)で`SYSCALL_DEFINE3`として定義されている。

[arch/x86/entry/entry_64.S#L115](https://github.com/torvalds/linux/blob/master/arch/x86/entry/entry_64.S#L115)
	call	do_syscall_64		/* returns with IRQs disabled */

### システムコールのモニタリング
main() 関数の呼び出し前の初期化シーケンスでも大量に呼ばれる。

Linux では strace, macOS では dtruss コマンド

### エラー処理
どのシステムコールも、大抵は正常の場合には0より大きい数値、エラーの場合には-1を返すようになっている。


## sec 6
TCPソケットとHTTPの実装

### プロトコルとレイヤー
インターネット通信で採用されているのはTCP/IP参照モデルであり、トランスポート層とアプリケーション層に気にする必要がある。

### RPC
RPC（Remote Procedure Calling）は、サーバーが用意しているさまざまな機能を、ローカルコンピューター上にある関数のように簡単に呼び出そう、という仕組み。Go言語では、JSON-RPCが標準ライブラリとして提供されている。

### REST
Resresentational State Transfer。RESTはHTTPのルールを最大限取り入れたプロトコルとなっている。

### ソケット
アプリケーション層からトランスポート層のプロトコルに対し、APIとして**ソケット**という仕組みを利用。TLSはソケットとHTTPの間に入って暗号化を行う！

一般に、他のアプリケーションとの通信をプロセス間通信という。ソケットは他と異なり、アドレスとポート番号がわかっていれば、外部のコンピュータとも通信が行える点。

ソケットにはいくつかあるが、主に次の３つ

- TCP
- UDP
- Unix ドメインソケット

### 速度改善
- HTTP/1.1 から、Keep-Alive が導入された。
- 圧縮を行う。
- チャンク形式でボディーを送信する。
- パイプライニング
  - HTTP/2 ではストリーム、という単位の名前


## sec 7
UDP はコネクションレスで、誰と繋がってるかを意識しない。複数のコンピュータに同時にメッセージを送ることが可能なマルチキャストとブロードキャストのサポート！DNS, NTP, ストリーミング動画, WebRTC などが UDP を利用

TCPには再送処理とフロー処理がある。また、ハンドシェイクに 1.5 RTT 分の時間がかかる。

## sec 8

### Unix ドメインソケット
外部インタフェースへの接続なし、代わりにカーネル内完結で高速。ウェブサーバーと NGINX などのリバースプロキシとの間、あるいはウェブサーバーとデータベースの間の接続を高速にできる場合がある。








