package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "taskdb.db")
	if err != nil {

		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// Log the error or return it to the caller
		fmt.Println("Error pinging the database:", err)
		return nil, err
	} else {
		fmt.Println("Connected to database")
	}
	Db = db
	return db, nil
}
