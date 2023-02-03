package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"chat-gpt/service"
)

func ApplyRoutes(r *gin.Engine) {
	db := service.ConnectDB()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{"greeting": "hello " + name})
	})

	r.GET("/group", func(c *gin.Context) {
		var groups []service.Group
		rows, err := db.Query("SELECT * FROM groups")
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var g service.Group
			err := rows.Scan(&g.ID, &g.Name)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{"error": "internal server error"})
				return
			}
			groups = append(groups, g)
		}
		c.JSON(200, groups)
	})
}
