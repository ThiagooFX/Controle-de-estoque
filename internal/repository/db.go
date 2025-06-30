package repository

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "estoque.db")
	if err != nil {
		log.Fatal(err)
	}

	createItensTable := `CREATE TABLE IF NOT EXISTS itens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		quantity INTEGER NOT NULL
	);`
	createLogsTable := `CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		action TEXT NOT NULL,
		ocorred_at TEXT NOT NULL,
		user TEXT NOT NULL,
		user_email TEXT NOT NULL,
		device TEXT NOT NULL,
		ip TEXT NOT NULL
	);`
	if _, err := DB.Exec(createItensTable); err != nil {
		log.Fatal(err)
	}
	if _, err := DB.Exec(createLogsTable); err != nil {
		log.Fatal(err)
	}
}
