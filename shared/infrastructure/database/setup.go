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

func initDb(dsn string) *gorm.DB {
	if _db != nil {
		return _db
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Printf("database: %s loaded\n", dsn)

	db.AutoMigrate(getEntities()...)

	_db = db

	return _db
}

func GetDatabase() *gorm.DB {
	return initDb("test.db")
}

func getTestDatabase() *gorm.DB {
	return initDb("testdata/test.db")
}

func Refresh() {
	getTestDatabase().Migrator().DropTable(getEntities()...)
	getTestDatabase().AutoMigrate(getEntities()...)
}
