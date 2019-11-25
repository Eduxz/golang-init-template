package migrations

import (
	pg "github.com/firebase-golang/connections"
	"github.com/firebase-golang/models"
	"github.com/firebase-golang/seeder"

	"github.com/jinzhu/gorm"
)

//Migrate all tables and seeders
func Migrate() {
	// Creating Tables
	pg.QueryPG(tables)
	// Seeders
	pg.QueryPG(seeders)
}

func tables(db *gorm.DB) {
	db.AutoMigrate(models.Role{})
	db.AutoMigrate(models.User{})
	db.Model(&models.User{}).AddForeignKey("role", "roles(name)", "RESTRICT", "RESTRICT")
}

func seeders(db *gorm.DB) {
	seeder.RoleSeeder(db)
}
