package main

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

// func countGoroutine() {
// 	for range time.Tick(3 * time.Second) {
// 		println("goroutine count:", runtime.NumGoroutine())
// 	}
// }

type IncomparableStruct struct {
	SomeFields int

	_ [0]struct{ _ []byte }
}

type Tea struct {
	Name string
}

// recover() の練習
// recover() は panic() で発生したエラーをキャッチする
func recoverPractice() {
	runtime.GC()

	// a := IncomparableStruct{
	// 	SomeFields: 3,
	// }
	// b := IncomparableStruct{}
	// if a == b {
	// }
	// aa := Tea {}
	// bb := Tea {}
	// if aa == bb {
	// }
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover")
		}
	}()

	fmt.Println("goroutine count: ", runtime.NumGoroutine())

	panic("panic")
}

func main() {
	envTest()
	return
	// recoverPractice()
	// return

	// go countGoroutine()

	// time.NewTicker(1 * time.Second)

	// os.ReadFile("main.go")

	// // e := gin.Default()
	// e := gin.New()
	// // e.RedirectFixedPath = false

	// e.GET("/", func(c *gin.Context) {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println("goroutine count: ", runtime.NumGoroutine())

	// 		// fmt.Println("hello world")
	// 		time.Sleep(1 * time.Second)
	// 	}

	// 	c.JSON(200, gin.H{})
	// })
	// e.GET("/pien", func(c *gin.Context) {
	// 	fmt.Println("hello world")
	// 	time.Sleep(199 * time.Second)
	// 	c.JSON(200, gin.H{})
	// })

	// auth := e.Group("/auth", func(c *gin.Context) {
	// 	fmt.Println("auth middleware")
	// 	c.Next()
	// })
	// auth.GET("/login", func(c *gin.Context) {
	// 	fmt.Println("login")
	// 	c.JSON(200, gin.H{})
	// })

	// fmt.Println("goroutine count: ", runtime.NumGoroutine())

	// if err := e.Run(":7772"); err != nil {
	// 	panic(err)
	// }

	// ここで goroutine が 1 つ増える
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
		}
		w.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
	})

	fmt.Println("http://localhost:8080/time")

	server := &http.Server{Addr: ":8080", Handler: nil}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}

	// 1
	fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())

	// ここで goroutine が ２ つ増えるのはなぜ？
	// 1. http.HandleFunc() で goroutine が 1 つ増える
	// 2. server.Serve() で goroutine が 1 つ増える
	// 3. net.Listen() で goroutine が 1 つ増える
	// 4. この main() 関数の goroutine が 1 つ増える
	// 5. 以上の合計で 4 つ増える
	if err := server.Serve(ln); err != nil {
		panic(err)
	}
}

func hoeg(pien interface{}) {
	switch pien.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	}
}
