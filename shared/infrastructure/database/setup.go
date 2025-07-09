package database

import (
	"log"

	models "github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB
var entities []interface{}

func getEntities() []interface{} {
	if entities != nil {
		return entities
	}

	entities = []interface{}{
		&models.User{},
		&models.Image{},
	}

	return entities
}

func GetDatabase() *gorm.DB {
	if _db != nil {
		return _db
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("database: test.db loaded")

	db.AutoMigrate(getEntities()...)

	_db = db

	return _db
}

func Refresh() {
	GetDatabase().Migrator().DropTable(getEntities()...)
	GetDatabase().AutoMigrate(getEntities()...)
}
