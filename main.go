package main

import (
	m "github.com/firebase-golang/migration"
	"github.com/firebase-golang/routes"
)

func main() {
	m.Migrate()
	e := routes.SetRoutes()
	e.Logger.Fatal(e.Start(":1323"))
}
