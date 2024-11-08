package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB инициализирует базу данных SQLite и создает таблицы, если их нет.
func InitDB() error {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./myserver.db"
	}

	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tokens (
		id INTEGER PRIMARY KEY,
		authToken TEXT
	);
	CREATE INDEX IF NOT EXISTS idx_tokens_id ON tokens(id);
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
