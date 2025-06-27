package entities

import "gorm.io/gorm"

type User struct {
	ID       uint
	Username string
	Password string
	Token    string
	gorm.Model
}
