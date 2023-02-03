package hello

import (
	"github.com/gin-gonic/gin"
)

// Greet function greets the user with given name.
func Greet(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{"greeting": "hello " + name})
}
