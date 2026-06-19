package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(dbPath string) {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	// Ensure users table exists (shared schema)
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id         INTEGER  PRIMARY KEY AUTOINCREMENT,
			username   TEXT     NOT NULL UNIQUE,
			password   TEXT     NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	log.Println("Database connected:", dbPath)
}
