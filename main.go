package main

import (
	"github.com/firebase-golang/app/routes"
	"github.com/firebase-golang/database/migration"
)

func main() {
	migration.Migrate()
	e := routes.SetRoutes()
	e.Logger.Fatal(e.Start(":1323"))
}
