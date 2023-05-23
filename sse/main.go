package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func currentGoroutine() {
	for range time.Tick(time.Second * 10) {
		fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
	}
}

func main() {
	go currentGoroutine()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/stream", addSseHeaders(), func(c *gin.Context) {
		for range time.Tick(time.Second * 3) {
			select {
			case <-c.Request.Context().Done():
				fmt.Println("client closed")
			default:
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			fmt.Println("sent")

			c.SSEvent("message", currentTime)

			c.Writer.Flush()
		}
	})

	router.Run(":8192")
}

func addSseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Writer.Header().Set("X-Accel-Buffering", "no")
		c.Next()
	}
}
