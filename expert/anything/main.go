package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
	"unsafe"
)

func main() {
	x := json.RawMessage(`{"foo":"bar"}`)
	fmt.Printf("x: %v\n", x)

	// io.Discard <= /dev/null
	f, err := os.CreateTemp("", "test-*.md")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("f.Name(): %v\n", f.Name())
	time.Sleep(1 * time.Second)

	// fs.FS
	// a := map[string]string{}

	// cast で変換すると、変換の際にデータのコピーが発生してしまう。
	s := "hello"
	fmt.Printf("s: %v\n", s)
	fmt.Printf("s address: %p\n", &s)
	fmt.Printf("unsafe.Pointer(&s): %v\n", unsafe.Pointer(&s))
	b := *(*[]byte)(unsafe.Pointer(&s))
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b address: %p\n", b)
	
	errors.Is(err, )

	return

	doneCh := make(chan struct{})
	for i := 0; i < 10; i++ {
		go do(i, doneCh)

		// go func() {
		// 	fmt.Printf("i in bad loop: %v\n", i)
		// }()
	}

	// broadcast !!!
	close(doneCh)
	time.Sleep(1 * time.Second)
}

func do(i int, doneCh <-chan struct{}) {
	for {
		select {
		case <-doneCh:
			fmt.Printf("finished %d\n", i)
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
