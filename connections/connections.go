package connections

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5400
	user     = "eduxz"
	dbname   = "system-manager"
	password = "123456"
)

func postgresData() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func connection(database string) *gorm.DB {
	db, err := gorm.Open("postgres", database)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}

// BooleanQueryPG with Posgrest
func BooleanQueryPG(callback func(database *gorm.DB) bool) bool {
	db := connection(postgresData())
	defer db.Close()

	return callback(db)
}

// QueryPG with Posgrest
func QueryPG(callback func(database *gorm.DB)) {
	db := connection(postgresData())
	defer db.Close()
	callback(db)
}
