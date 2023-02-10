package example

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"time"
)

func Goroutine() {
	fmt.Println("Start goroutine")
	go func() {
		fmt.Println("running goroutine...")
	}()

	fmt.Println("wait")
	time.Sleep(time.Second)
	fmt.Println("Finish goroutine")
}

func Channel() {
	ic := make(chan int)

	// 送信
	ic <- 100

	// 受信1. 結果は捨てる
	<-ic
	// 受信2. 変数に入れる
	r := <-ic
	// 受信2. 結果とチャネルの状態を変数に入れる
	r, ok := <-ic
	fmt.Println(r, ok)

	go func() {
		ic <- 10
		ic <- 200
		close(ic)
	}()
	// close されるとループ解除
	for v := range ic {
		fmt.Println(v)
	}
}

type Task string
type Result struct {
	Value int64
	Task  Task
	Err   error
}

func worker(id int, tasks <-chan Task, results chan<- Result) {
	for t := range tasks {
		fmt.Printf("worker: %d and task: %s\n", id, t)
		s, err := os.Stat(string(t))
		if err == nil && s.IsDir() {
			err = fmt.Errorf("worker: %d and err: %s is dir", id, string(t))
		}

		result := Result{
			Task: t,
		}
		if err != nil {
			result.Err = err
		} else {
			fmt.Printf("worker: %d and path: %s and size: %d\n", id, string(t), s.Size())
			result.Value = s.Size()
		}
		results <- result
	}
}
func TotalFileSize() int64 {
	tasks := make(chan Task)
	results := make(chan Result)

	// ワーカーの起動
	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, tasks, results)
	}
	// タスクを非同期でチャネルに投入！
	in := make(chan struct{})
	var remainedCount int64
	go func() {
		filepath.Walk(runtime.GOROOT(), func(path string, info os.FileInfo, err error) error {
			// atomic のパッケージなんだっけ
			atomic.AddInt64(&remainedCount, 1)
			tasks <- Task(path)
			return nil
		})
		// https://qiita.com/tenntenn/items/dd6041d630af7feeec52
		close(in)
		close(tasks)
	}()

	// 結果の収集
	var size int64
	for {
		select {
		case result := <-results:
			if result.Err != nil {
				fmt.Printf("err %v for %s\n", result.Err, result.Task)
			} else {
				atomic.AddInt64(&size, result.Value)
			}
			atomic.AddInt64(&remainedCount, -1)
		case <-in:
			if remainedCount == 0 {
				return size
			}
		}
	}
}

func UnknownTasks() {
	total := TotalFileSize()
	fmt.Println("Total Size: ", total)
}

func fixedTasks(taskSrcs []Task) int64 {
	// 数がわかっている場合は、、ちょうどだけ作るのがよい！
	tasks := make(chan Task, len(taskSrcs))
	results := make(chan Result)
	for _, src := range taskSrcs {
		tasks <- src
	}
	close(tasks)

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, tasks, results)
	}

	var count int
	var size int64
	for {
		result := <- results
		count += 1
		if result.Err != nil {
			fmt.Printf("err %v for %s\n", result.Err, result.Task)
		} else {
			size += result.Value
		}
		if count == len(taskSrcs) {
			break
		}
	}
	return size
}
