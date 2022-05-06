package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	dirInfo()
	tilde()
	go selectAPI()
	traverseDirToFindImages()
	time.Sleep(1000 * time.Second)
}

func selectAPI() {
	// macOS の runtime パッケージ内で利用されている kqueue
	kq, err := syscall.Kqueue()
	if err != nil {
		panic(err)
	}
	// 監視対象のファイルディスクリプタ
	fd, err := syscall.Open("memo.md", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	// 監視したいイベントの構造体
	ev1 := syscall.Kevent_t{
		Ident:  uint64(fd),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE,
		Data:   0,
		Udata:  nil,
	}
	// イベント待ちの無限ループ
	for {
		events := make([]syscall.Kevent_t, 10)
		nev, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		if err != nil {
			panic(err)
		}
		// イベントを確認
		for i := 0; i < nev; i++ {
			fmt.Printf("Event [%v] -> %v\n", i, events[i])
		}
	}
}

func traverseDirToFindImages() {
	imageSuffix := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
		".gif":  true,
		".tiff": true,
		".eps":  true,
	}

	root := "."

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if info.Name() == "_build" {
					return filepath.SkipDir
				}
				return nil
			}
			ext := strings.ToLower(filepath.Ext(info.Name()))
			if imageSuffix[ext] {
				rel, err := filepath.Rel(root, path)
				if err != nil {
					return nil
				}
				fmt.Printf("%s\n", rel)
			}
			return nil
		})
	if err != nil {
		fmt.Println(1, err)
	}
}

func tilde() {
	fmt.Println(os.UserHomeDir())
	fmt.Println(Clean2("~/${LANG}/p/../hoge.txt"))
}

func Clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func dirInfo() {
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	fileInfos, err := dir.ReadDir(-1) // 負の数は全要素取得
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("[Dir] %s\n", fileInfo.Name())
		} else {
			fmt.Printf("[File] %s\n", fileInfo.Name())
		}
	}
	fmt.Printf("\n===== path/filepath =====\n")
	fmt.Printf("Temp File Path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))
	fmt.Printf("GOPATH Base: %s\n", filepath.Base(os.Getenv("GOPATH")))
	fmt.Printf("GOPATH Dir: %s\n", filepath.Dir(os.Getenv("GOPATH")))
	fmt.Printf("GOPATH Ext: %s\n", filepath.Ext(os.Getenv("GOPATH")))
}

func timer(sec int) {
	timeout := time.After(time.Duration(sec) * time.Second)

	// このforループを1秒間ずっと実行し続ける
	for {
		select {
		case <-timeout:
			fmt.Println("time out")
			return
		default:
			// fmt.Println("default")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func signalNotify() {
	signals := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C) を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルがくるまで待つ
	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}

func printPrimeNumbers() {
	pn := primeNumbers()
	for n := range pn {
		fmt.Println(n)
	}
}

func primeNumbers() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 1000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
			time.Sleep(500 * time.Millisecond)
		}
		close(result)
	}()
	return result
}

func chanel() {
	fmt.Println("start sub()")
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("sub is finished")
		done <- struct{}{}
	}()
	<-done
	fmt.Println("all tasks are finished")
}

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func goroutine() {
	fmt.Println("start sub()")
	go sub()

	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
