package app

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
)

var dbUser string = "utr1903"
var dbPass string = "utr1903"

// InitDb : Initializes the Db connection
func TestInitDb(t *testing.T) {

	connectionString := dbUser + ":" + dbPass + "@(127.0.0.1:3306)/counterdb?parseTime=true"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Error by database initialization", err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error by database connection", err)
		os.Exit(1)
	}

	createTableIfNotExists(db)
	initializeTableIfNotExists(db)
}

// Create a new table if not exists
func createTableIfNotExists(db *sql.DB) error {

	q := `
		create table if not exists counter (
			id int not null,
			counter int not null,
			primary key (id)
		)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	if _, err := db.ExecContext(ctx, q); err != nil {
		return err
	}

	return nil
}

// Initialize table with id = 2 and counter = 0 for testing
func initializeTableIfNotExists(db *sql.DB) error {

	q := `
		insert into counterdb.counter
		(id, counter)
		values
		(2, 0)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}
