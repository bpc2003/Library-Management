package models

import (
	"database/sql"
	"fmt"
	"log"
)

func MakeBook(db *sql.DB) {
	stmt := `CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		isbn TEXT NOT NULL,
		genre TEXT,
		UNIQUE (title, isbn)
);`

	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stmt)
}
