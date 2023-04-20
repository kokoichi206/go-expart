package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// sleep の間にリクエストをキャンセルしても内部処理は終わらない + return されない。
		time.Sleep(15 * time.Second)
		fmt.Println("sleep finished!")
		c.String(200, "Hello, World!")
	})

	r.GET("/cancel", func(c *gin.Context) {
		ctx := c.Request.Context()

		signal := make(chan struct{}, 1)

		go func() {
			time.Sleep(10 * time.Second)

			signal <- struct{}{}
		}()

		select {
		case <-ctx.Done():
			// 499 は go の http パッケージで定義されていないので、
			// 499 のステータスコードは使わないほうがよさそう。
			c.String(499, "context canceled!")
			fmt.Println("context canceled!")

			return
		case <-signal:
			fmt.Println("long process finished!")
		}

		close(signal)

		c.String(200, "Hello, World!")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
