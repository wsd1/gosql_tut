package app

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"os"
)

var (
	SqlDb *sql.DB
)

func Init() {

	db, err := sql.Open("sqlite3", "wiki.db")
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	SqlDb = db

	SqlDb.SetMaxIdleConns(20)
	SqlDb.SetMaxOpenConns(30)

}
