// db/db.go
package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite" // Используем чистый Go SQLite
)

var DB *sql.DB

// InitDB инициализирует базу данных SQLite и создает таблицы, если их нет.
func InitDB() error {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./myserver.db"
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath) // Обновили параметр на "sqlite"
	if err != nil {
		return err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS tokens (
        id TEXT PRIMARY KEY,
        club_id TEXT,
		extra_info TEXT
    );
	CREATE INDEX IF NOT EXISTS idx_tokens_id ON tokens (id);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

// CloseDB закрывает соединение с базой данных.
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
