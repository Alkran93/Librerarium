package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "./cart.db")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS cart_items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id TEXT NOT NULL,
			quantity INTEGER NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("DB table creation failed:", err)
	}
}
