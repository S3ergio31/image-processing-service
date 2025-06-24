package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/register/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type userModel struct {
	ID       uint
	Username string
	Password string
	gorm.Model
}

func (userModel) TableName() string {
	return "users"
}

type sqliteUserRepository struct {
	db *gorm.DB
}

func (repository sqliteUserRepository) Save(user domain.User) {
	repository.db.Create(&userModel{
		Username: user.Username(),
		Password: user.Password(),
	})
}

func (repository sqliteUserRepository) UsedUsername(username string) bool {
	var userModel userModel
	repository.db.First(&userModel, "username = ?", username)

	return userModel.Username == username
}

var repository *sqliteUserRepository

func NewSqliteUserRepository() sqliteUserRepository {
	if repository != nil {
		return *repository
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("NewSqliteUserRepository: failed to connect database")
	}

	db.AutoMigrate(&userModel{})

	repository = &sqliteUserRepository{db: db}

	return *repository
}
