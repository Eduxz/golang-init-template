package usermodel

import (
	"github.com/jinzhu/gorm"
)

// User info
type User struct {
	gorm.Model
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       uint   `json:"age"`
}

func (u User) create(db *gorm.DB) {
	db.NewRecord(u)
	db.Create(&u)
}
