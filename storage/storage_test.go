package storage

import (
	"testing"
)

func TestInsertUser(t *testing.T) {
	InitDatabase()
	stmt, err := DbHandler.db.Prepare("INSERT INTO employees(name, position, salary) values(?,?,?)")
	if err != nil {
		t.Errorf("Failed executing migrations: %s ", err)
	}
	_, err = stmt.Exec("astaxie", "blockchain developer", "1209 USD")
	if err != nil {
		t.Errorf("Failed executing example query: %s ", err)
	}
}
