package models

import (
	"github.com/jinzhu/gorm"
)

// Role info
type Role struct {
	Name string `gorm:"type:varchar(20);unique_index"`
	gorm.Model
}

func (role Role) create(db *gorm.DB) bool {
	return db.Create(&role).RowsAffected > 0
}
