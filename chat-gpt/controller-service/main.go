package main

import (
	"github.com/gin-gonic/gin"
	"chat-gpt/controller"
	"chat-gpt/service"
)

func main() {
	db := service.ConnectDB()
	defer db.Close()

	r := gin.Default()
	controller.ApplyRoutes(r)
	r.Run()
}
