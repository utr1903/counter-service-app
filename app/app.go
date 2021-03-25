package app

import (
	"database/sql"
	"log"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/utr1903/counter-service-app/controllers"
)

// App : DB and Controllers
type App struct {
	Db     *sql.DB
	Router *mux.Router
}

// InitDb : Initializes the Db connection
func (a *App) InitDb() {
	db, err := sql.Open("mysql", "utr1903:utr1903@(127.0.0.1:3306)/counter?parseTime=true")
	a.Db = db
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create a new table if not exists
	q := `
		create table if not exists counter (
			id int not null,
			counter int not null,
			primary key (id)
		)`

	res, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize table
	q = `
		insert into counter
		(id, counter)
		values
		(1, 0)`

	res, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
}

// InitControllers : Initializes the controllers
func (a *App) InitControllers() {
	r := mux.NewRouter()
	a.Router = r

	b := &controllers.ControllerBase{Db: a.Db}
	c := &controllers.CounterController{Base: b}

	a.Router.HandleFunc("/counter/GetCounter", c.GetCounter).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/counter/IncreaseCounter", c.IncreaseCounter).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/counter/DecreaseCounter", c.DecreaseCounter).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/counter/ResetCounter", c.ResetCounter).Methods("POST", "OPTIONS")
}
