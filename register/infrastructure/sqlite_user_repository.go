package infrastructure

import (
	"log"

	"github.com/S3ergio31/image-processing-service/register/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/gorm"
)

type sqliteUserRepository struct {
	*gorm.DB
}

func (repository sqliteUserRepository) Save(user domain.User) {
	repository.Create(&entities.User{
		Username: user.Username(),
		Password: user.Password(),
	})
}

func (repository sqliteUserRepository) UsedUsername(username string) bool {
	var userEntity entities.User
	repository.First(&userEntity, "username = ?", username)

	log.Println("UsedUsername", username, userEntity.Username)
	return userEntity.Username == username
}

var repository *sqliteUserRepository

func NewSqliteUserRepository() sqliteUserRepository {
	if repository != nil {
		return *repository
	}

	repository = &sqliteUserRepository{DB: database.GetDatabase()}

	return *repository
}
