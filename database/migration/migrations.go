package migration

import (
	"github.com/firebase-golang/app/models"
	"github.com/firebase-golang/database/connections"
	"github.com/firebase-golang/database/seeder"

	"github.com/jinzhu/gorm"
)

//Migrate all tables and seeders
func Migrate() {
	// Creating Tables
	connections.QueryPG(tables)
	// Seeders
	connections.QueryPG(seeders)
}

func tables(db *gorm.DB) {
	db.AutoMigrate(models.Role{})
	db.AutoMigrate(models.User{})
	db.Model(&models.User{}).AddForeignKey("role", "roles(name)", "RESTRICT", "RESTRICT")
}

func seeders(db *gorm.DB) {
	seeder.RoleSeeder(db)
}
