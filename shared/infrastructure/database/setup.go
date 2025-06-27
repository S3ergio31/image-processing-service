package database

import (
	"log"

	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func GetDatabase() *gorm.DB {
	if _db != nil {
		return _db
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("database: test.db loaded")

	db.AutoMigrate(&entities.User{})

	_db = db

	return _db
}

func Refresh() {
	GetDatabase().Migrator().DropTable(&entities.User{})
	GetDatabase().AutoMigrate(&entities.User{})
}
