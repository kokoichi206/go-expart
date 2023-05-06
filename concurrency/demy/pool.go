package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

func poolSample() {
	fmt.Println("========== poolSample ==========")
	log("hello")
	log("world")
}

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes.Buffer")

		return new(bytes.Buffer)
	},
}

func log(log string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString(log)
	b.WriteString("\n")

	w := os.Stdout
	w.Write(b.Bytes())

	bufPool.Put(b)
}
