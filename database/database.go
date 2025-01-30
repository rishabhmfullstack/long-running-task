package database

import (
	"database/sql"
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
			return
		}

		_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS tasks (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				status TEXT NOT NULL,
				result TEXT
			);
		`)
	})
	return db, err
}

func GetDB() *sql.DB {
	return db
}
