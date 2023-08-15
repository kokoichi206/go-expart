package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	<-ctx.Done()
	fmt.Println("handler end")
	return
}

func handlerWithout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithoutCancel(ctx)
	<-ctx.Done()
	fmt.Println("handlerWithout end")
	return
}

func main() {
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/handler1", handlerWithout)
	go http.ListenAndServe(":8818", nil)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	stopc := make(chan struct{})
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("stop")
		close(stopc)
	})
	time.Sleep(2 * time.Second)

	// stop を呼ぶことで、引数に渡した関数を実行する！
	// ctx.Done がそれより前に呼ばれていた場合、stop は false を返し、関数も実行されない！
	if !stop() {
		fmt.Println("waiting stopc...!")

		// wait for it's completion!
		<-stopc
		fmt.Println("done!")
	}
	fmt.Println("finished!")

	fmt.Printf("errors.ErrUnsupported: %v\n", errors.ErrUnsupported)
}
