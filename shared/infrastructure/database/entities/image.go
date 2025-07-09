package entities

import "gorm.io/gorm"

type Image struct {
	ID     uint
	UserID uint
	Name   string
	Path   string
	User   User
	gorm.Model
}
