package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/hello", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	if err := engine.Run(":21829"); err != nil {
		log.Fatal(err)
	}
}
