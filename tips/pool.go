package main

import (
	"io"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		return make([]byte, 1024)
	},
}

// 呼び出す毎に新たな []byte スライスが作成されるわけではない。
func write(w io.Writer) {
	buf := pool.Get().([]byte)
	// リセットする。
	buf = buf[:0]
	defer pool.Put(buf)

	// 提供されたバッファを使って書き込む。
	getResponse(buf)
	_, _ = w.Write(buf)
}

func getResponse(buf []byte) {
	// buf に書き込む。
}
