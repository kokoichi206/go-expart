package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func heavyCalc(wg *sync.WaitGroup) {
	a := 3
	b := 0
	res := a / b
	log.Println(res)
	wg.Done()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[5])
	return

	engine := gin.New()
	// Recovery は gin.Default() にも含まれる。
	engine.Use(gin.Recovery())
	engine.Handle(http.MethodGet, "/calc", func(context *gin.Context) {
		a := 3
		b := 0
		res := a / b
		context.JSON(http.StatusOK, gin.H{
			"results": res,
		})
	})
	engine.Handle(http.MethodGet, "/goroutine", func(context *gin.Context) {
		var wg sync.WaitGroup
		wg.Add(3)
		for i := 0; i < 3; i++ {
			go func() {
				defer func() {
					if err := recover(); err != nil {
						// 何かしらの処理。
						fmt.Printf("err: %v\n", err)
					}
				}()
				heavyCalc(&wg)
			}()
		}
		wg.Wait()
	})

	srv := &http.Server{
		Addr:    ":21829",
		Handler: engine,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
