package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func server1() error {
	r := gin.Default()
	r.Use(requestid.New())
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})
	r.GET("/connect2", func(c *gin.Context) {
		req, _ := http.NewRequest("GET", "http://localhost:2023/hello", nil)
		req.Header.Set("X-Request-ID", requestid.Get(c))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		defer res.Body.Close()
		io.Copy(c.Writer, res.Body)
	})
	return r.Run(":1919")
}

func server2() error {
	r := gin.Default()
	r.Use(requestid.New())
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
		fmt.Printf("requestid.Get(c): %v\n", requestid.Get(c))
		fmt.Printf("c.Request.Header.Get(\"X-Request-ID\"): %v\n", c.Request.Header.Get("X-Request-ID"))
	})
	return r.Run(":2023")
}

func main() {
	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return server1()
	})
	eg.Go(func() error {
		return server2()
	})
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
