package main

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var dbUser string = "utr1903"
var dbPass string = "utr1903"

func TestMain(m *testing.M) {

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

	os.Exit(m.Run())
}
