package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello, World")
}

func standardServer() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}

func ginServer() {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/gin", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World Gin",
		})
	})

	if err := engine.Run(":21829"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	standardServer()

	// ginServer()
}
