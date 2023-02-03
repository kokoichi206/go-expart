package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saucelabs/example-clean-arch-gin/app/group"
	"github.com/saucelabs/example-clean-arch-gin/app/health"
	"github.com/saucelabs/example-clean-arch-gin/app/hello"
)

func main() {
	router := gin.Default()

	router.GET("/health", health.Check)
	router.GET("/hello", hello.Greet)
	router.GET("/group", group.GetGroups)

	router.Run()
}
