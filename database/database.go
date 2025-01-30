package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB(dbFile string) (*sql.DB, error) {
	var err error
	once.Do(func() {

		db, err = sql.Open("sqlite3", dbFile)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
			return
		}

		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS tasks (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				status TEXT NOT NULL,
				result TEXT,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);
		`)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
			return
		}
	})
	return db, err
}

func GetDB() *sql.DB {
	return db
}
