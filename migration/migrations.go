package migrations

import (
	pg "github.com/firebase-golang/conections"
	model "github.com/firebase-golang/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type any interface{}

//Migrate all table
func Migrate() {
	// Creating Tables
	pg.Connect(tables)
}

func tables(db *gorm.DB) {
	db.AutoMigrate(model.User{})
}
