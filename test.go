package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "wiki.db")
	if err != nil {
		log.Println(err)
		return
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(30)

	defer db.Close()
}
