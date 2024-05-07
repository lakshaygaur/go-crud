package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbHandler Database

const DB_PATH = "./gocrud.db"

func InitDatabase() {
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	DbHandler = Database{db}

	DbHandler.RunMigrations()
	// return
}
