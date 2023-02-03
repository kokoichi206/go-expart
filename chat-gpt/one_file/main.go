package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("postgres", "user=root password=rootpassword dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"greeting": fmt.Sprintf("hello %s", name),
		})
	})

	router.GET("/group", func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM groups")
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var groups []Group
		for rows.Next() {
			var group Group
			if err := rows.Scan(&group.ID, &group.Name); err != nil {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			groups = append(groups, group)
		}

		c.JSON(http.StatusOK, groups)
	})

	router.Run(":8080")
}
