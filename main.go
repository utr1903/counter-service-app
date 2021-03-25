package main

import (
	"log"
	"os"

	app "github.com/utr1903/counter-service-app/app"
)

func main() {
	a := &app.App{}
	a.InitDb(getDbCreditentials())
	a.InitControllers()
	a.Serve()
}

func getDbCreditentials() (*string, *string) {

	if len(os.Args) != 3 {
		log.Fatal("User and Password have to be given sequentially as command line arguments !")
		os.Exit(1)
	}

	dbUser := os.Args[1]
	dbPass := os.Args[2]

	return &dbUser, &dbPass
}
