package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	// connectionString := "postgres://db-username@localhost/go-mvc?sslmode=disable"
	// connectionString := os.Getenv("DATABASE_URL")
	connectionString := "dbname=go-mvc sslmode=disable"

	// Open connection to database
	Db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	// Ping database
	if err = Db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
