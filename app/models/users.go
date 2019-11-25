package models

import (
	"github.com/firebase-golang/database/connections"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User info
type User struct {
	gorm.Model
	FirstName string `json:"first-name" gorm:"type:varchar(50)"`
	LastName  string `json:"last-name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password"`
	Age       uint   `json:"age"`
	Role      string `json:"role" gorm:"default:'slow'"`
}

// Create new user
func (u User) Create(db *gorm.DB) bool {
	u.Password, _ = hashPassword(u.Password)
	return db.Create(&u).RowsAffected > 0
}

//CheckLogin verify if the user is within the DB
func (u *User) CheckLogin() {
	match := User{}
	var rowReturned int64
	connections.QueryPG(func(db *gorm.DB) {
		rowReturned = db.Where("email = ?", u.Email).First(&match).RowsAffected
	})
	if rowReturned == 1 && checkPasswordHash(u.Password, match.Password) {
		*u = match
	}
}

// IsValid confirm the data in user Model
func (u User) IsValid() bool {
	return u.Email != "" && u.FirstName != "" && u.LastName != "" && u.Password != "" && u.Age != 0
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
