package health

import (
	"github.com/gin-gonic/gin"
)

// Check function returns "ok" status.
func Check(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
