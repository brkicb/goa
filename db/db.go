package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createListingsTable := `
	CREATE TABLE IF NOT EXISTS listings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        slug TEXT NOT NULL,
		address TEXT NOT NULL,
		price INTEGER NOT NULL
	)
	`

	_, err := DB.Exec(createListingsTable)

	if err != nil {
		panic("Could not create events table")
	}
}
