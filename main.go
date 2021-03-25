package main

import (
	app "github.com/utr1903/counter-service-app/app"
)

func main() {
	a := &app.App{}
	a.InitDb()
	a.InitControllers()
	a.Serve()
}