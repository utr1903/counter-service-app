package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/utr1903/counter-service-app/controllers"
	"github.com/utr1903/counter-service-app/controllers/countercontroller"
)

// App : DB and Controllers
type App struct {
	Db     *sql.DB
	Router *mux.Router
}

// InitDb : Initializes the Db connection
func (a *App) InitDb(dbUser *string, dbPass *string) {

	connectionString := *dbUser + ":" + *dbPass + "@(127.0.0.1:3306)/counterdb?parseTime=true"

	db, err := sql.Open("mysql", connectionString)
	a.Db = db
	if err != nil {
		log.Fatal("Error by database initialization", err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Error by database connection", err)
		os.Exit(1)
	}

	a.createTableIfNotExists()
	a.initializeTableIfNotExists()
}

// Create a new table if not exists
func (a *App) createTableIfNotExists() error {

	q := `
		create table if not exists counter (
			id int not null,
			counter int not null,
			primary key (id)
		)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	if _, err := a.Db.ExecContext(ctx, q); err != nil {
		return err
	}

	return nil
}

// Initialize table with id = 1 and counter = 0
func (a *App) initializeTableIfNotExists() error {

	q := `
		insert into counterdb.counter
		(id, counter)
		values
		(1, 0)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := a.Db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

// InitControllers : Initializes the controllers
func (a *App) InitControllers() {
	r := mux.NewRouter()
	a.Router = r

	b := &controllers.ControllerBase{Db: a.Db}
	c := &countercontroller.CounterController{Base: b}

	a.Router.HandleFunc("/counter/GetCounter", c.GetCounter).Methods("GET", "OPTIONS")
	a.Router.HandleFunc("/counter/IncreaseCounter", c.IncreaseCounter).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/counter/DecreaseCounter", c.DecreaseCounter).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/counter/ResetCounter", c.ResetCounter).Methods("POST", "OPTIONS")
}

// RouterWithCORS : To prevent getting CORS errors from Angular UI
type RouterWithCORS struct {
	r *mux.Router
}

// Serve : Runs web server
func (a *App) Serve() {
	http.ListenAndServe(":8080", &RouterWithCORS{a.Router})
}

// ServeHTTP : A middleware to add necessary headers in order not to get CORS error
func (s *RouterWithCORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		// w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here for a Preflighted OPTIONS request.
	if r.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.r.ServeHTTP(w, r)
}
