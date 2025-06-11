package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	DB *sql.DB
}

func SQLiteINIT(DataBaseName string) *DataBase {
	db, err := sql.Open("sqlite3", DataBaseName)
	if err != nil {
		log.Fatal("Cannot open DB:", err)
	}
	create := `CREATE TABLE IF NOT EXISTS movies(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR,
		genre VARCHAR,
		year INTEGER,
		description VARCHAR
	);`
	db.Exec(create)
	return &DataBase{DB: db}
}
