package tasks

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	// sqlite3 driver conforming to the built-in database/sql interface
	_ "github.com/mattn/go-sqlite3"
)

// InitDB retourne la database todosDB.db
func InitDB() (*sql.DB, error) {
	dir := os.Getenv("GOBIN")
	db, err := sql.Open("sqlite3", filepath.Join(filepath.Dir(dir), "bin", "todosDB.db"))
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
