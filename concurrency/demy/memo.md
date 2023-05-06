## process and thread

- OS のスケジューラーはスレッドレベルでスケジューリングする
- Runnable → Executing
  - I/O event etc → Waiting
- コンテキストスイッチングはコスト高い！
  - コピーが走ったりするので
- スレッド
  - Fixed stack size
  - メモリーを共有するので複雑度が増す
    - Data Race
    - 適当にロックすると Dead Lock 発生したりする

## Goroutines

- CSP: Communicating Sequential Processes
  - sequential execution
  - メモリをシェアするのではなく、通信によってデータを共有する
- goroutine
  - go runtime によって管理される、ユーザースペースのスレッド
  - スタートのスタックサイズが 2kb とと、ばか小さい
  - CPU のオーバーヘッドが小さい
  - Go のランタイムが、ワーカー OS スレッド作成
  - OS スレッドのコンテキストの中で goroutine 実行される

## go scheduler

- M:N スケジューラー
- ユーザースペース
- asynchronous preemption
  - preemption
    - executing to runnable
- synchronous system call
  - read file . etc
  - I/O オペレーションが終わるまで待つ
  - 平行性を低下させる
  - go scheduler はそのとき持ってた他のタスクを、別のスレッドに割り当てる！
    - スレッドプールから
- asynchronous system call
  - network call
  - file descriptor is set to non-blocking mode
  - netpoller
    - **非同期システムコールを、ブロッキング市末t無コールに変える！**
    - epoll, kqueue, iocp などが OS によって提供されてる
- Work stealing!

## channels

- buffered channels
  - インメモリーの FIFO
- data structure
  - https://github.com/golang/go/blob/3e35df5edbb02ecf8efd6dd6993aabd5053bfc66/src/runtime/chan.go#L33-L52


## race detector

- go test -race mypkg
- go run -race main.go


## pipeline

- streams and batches of data
- a stage could consume and return the same type
- composability of pipeline
  - `square(square(generator(2, 3)))`
  - image processing pipeline

## Fan-out, Fan-in

- どこか一箇所が重い処理だったりするとき
  - そこだけ multiple goroutine に分解して、最後に合体させる！
  - 分解の Fan-Out
  - 統合の Fan-In


## context

- context.Context is an interface
- cancel とかは Done のチャネルをクローズしてるだけ
  - その close ノブロードキャスをと受け取っている！
  - WithTimeout is a wrapper over withDeadline
- I/O を伴う関数には ctx を渡そうか

## handler

- http.TimeoutHandler とかあるポピ


