package seeder

import (
	"github.com/firebase-golang/app/models"
	"github.com/jinzhu/gorm"
)

// RoleSeeder seed of roles
func RoleSeeder(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE roles")
	db.Create(&models.Role{Name: "admin"})
	db.Create(&models.Role{Name: "low"})
	db.Create(&models.Role{Name: "middle"})
	db.Create(&models.Role{Name: "high"})
}
