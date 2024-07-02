package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// The file we'll be using to store the counter data.
	const dbfile string = "laskuri.db"

	// The sql for creating the necessary table at program
	// start up if it doesn't exist.
	const create string = `
    CREATE TABLE IF NOT EXISTS laskuri (
	    id INTEGER NOT NULL PRIMARY KEY,
      time DATETIME NOT NULL,
      description TEXT
    );`

	// Open the database connection.
	db, err := sql.Open("sqlite3", dbfile)

	// Initialize the HTTP listener
	r := gin.Default()

	// And run it for the sole end-point we provide.
	r.GET("/laskuri", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
