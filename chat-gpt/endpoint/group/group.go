package group

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// GetGroups function retrieves all groups from the database.
func GetGroups(c *gin.Context) {
	db, err := sql.Open("postgres", "user=root password=rootpassword dbname=postgres sslmode=disable")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM groups")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var groups []map[string]interface{}
	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		groups = append(groups, map[string]interface{}{"id": id, "name": name})
	}

	c.JSON(200, groups)
}
