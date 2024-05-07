package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func HandleDBclose() int {
	DbHandler.db.Close()
	return 1
}

func (DbHandler Database) RunMigrations() {
	_, err := DbHandler.db.Exec("CREATE TABLE IF NOT EXISTS `employees`" +
		" (`id` INTEGER PRIMARY KEY AUTOINCREMENT,	`name` VARCHAR(64) NULL," +
		"	`position` VARCHAR(200) NULL,`salary` FLOAT NULL);")
	if err != nil {
		log.Fatal("Failed preparing migration statement ", err)
	}
	// os.Exit(HandleDBclose())
}
