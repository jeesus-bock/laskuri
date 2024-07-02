package main

import (
	"context"
	"database/sql"
	"hash/fnv"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// The file we'll be using to store the counter data.
	const dbFile string = "laskuri.db"

	// The sql for creating the necessary table at program
	// start up if it doesn't exist.
	const createSql string = `
    CREATE TABLE IF NOT EXISTS laskuri (
	    id INTEGER NOT NULL PRIMARY KEY
    );`

	// The SQL to insert a hash row to db. If the hash already exists in the db do nothing.
	const insertHashSql string = `INSERT OR IGNORE INTO laskuri ('id') VALUES(?);`

	// Open the database connection.
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run the initial db table creation SQL.
	statement, err := db.Prepare(createSql)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize our hasher to be used for masking the
	// IP-addresses of the visitors.
	fnvHash := fnv.New32a()

	// Initialize the HTTP listener
	r := gin.Default()

	// Define the sole end-point we provide.
	r.GET("/laskuri", func(c *gin.Context) {
		// The Nginx proxy fills this header with the original
		// requester IP-address
		origIp := c.Request.Header.Get("X-Real-IP")
		fnvHash.Write([]byte(origIp))
		ipHash := fnvHash.Sum32()
		fnvHash.Reset()

		// Insert the hash into database (ignore duplicates)
		_, err := db.ExecContext(context.Background(), insertHashSql, ipHash)
		if err != nil {
			log.Fatal(err)
		}

		// Get the total count of hashes in the db.
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM laskuri;").Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{
			"count": count,
		})
	})

	// And run it indefinitely
	r.Run() // listen and serve on 0.0.0.0:8080
}
