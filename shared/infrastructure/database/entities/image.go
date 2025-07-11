package entities

import "gorm.io/gorm"

type Image struct {
	ID     uint
	Uuid   string
	UserID uint
	Name   string
	Type   string
	Path   string
	User   User
	gorm.Model
}
