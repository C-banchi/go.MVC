package models

import "gorm.io/gorm"

// gotta add the new models class to the db.go
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password []byte
	Active   string
	Role     string
}
