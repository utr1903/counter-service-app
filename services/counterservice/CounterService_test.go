package counterservice

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"testing"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
)

func TestGetCounter(t *testing.T) {

	db := initDb()
	s := initService(db)

	r := s.GetCounter()

	if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
		t.Log("True usage")
	} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
		t.Log("Exception caugth")
	} else {
		t.Error("Failed")
	}
}

func TestIncreaseCounter(t *testing.T) {

	db := initDb()
	s := initService(db)

	inputs := []string{"1", "5", "x"}

	for _, input := range inputs {
		r := s.IncreaseCounter(&input)

		if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
			t.Log("True usage")
		} else if r.Counter == nil && r.Code == http.StatusBadRequest && r.Error != nil {
			t.Log("Exception caugth")
		} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
			t.Log("Exception caugth")
		} else {
			t.Error("Failed")
		}
	}
}

func TestDecreaseCounter(t *testing.T) {

	db := initDb()
	s := initService(db)

	inputs := []string{"1", "5", "x"}

	for _, input := range inputs {
		r := s.IncreaseCounter(&input)

		if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
			t.Log("True usage")
		} else if r.Counter == nil && r.Code == http.StatusBadRequest && r.Error != nil {
			t.Log("Exception caugth")
		} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
			t.Log("Exception caugth")
		} else {
			t.Error("Failed")
		}
	}
}

func TestResetCounter(t *testing.T) {

	db := initDb()
	s := initService(db)

	r := s.ResetCounter()

	if r.Counter != nil && r.Code == http.StatusOK && r.Error == nil {
		t.Log("True usage")
	} else if r.Counter == nil && r.Code == http.StatusInternalServerError && r.Error != nil {
		t.Log("Exception caugth")
	} else {
		t.Error("Failed")
	}
}

func initDb() *sql.DB {

	dbUser := os.Getenv("MYSQL_USERNAME")
	dbPass := os.Getenv("MYSQL_PASSWORD")

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

	return db
}

func initService(db *sql.DB) *CounterService {
	return &CounterService{
		Db:     db,
		IsProd: false,
	}
}
