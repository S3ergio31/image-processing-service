package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/login/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/gorm"
)

type sqliteUserRepository struct {
	*gorm.DB
}

func (repository sqliteUserRepository) Save(user domain.User) {
	repository.Model(&entities.User{}).
		Where("username = ?", user.Username()).
		Update("token", user.Token())
}

func (repository sqliteUserRepository) FindByUsername(username string) (domain.User, *domain.UserNotFound) {
	var userEntity entities.User
	repository.First(&userEntity, "username = ?", username)

	if userEntity.Username != username {
		return nil, &domain.UserNotFound{}
	}

	user, _ := domain.NewUser(userEntity.Username, userEntity.Password)

	return user, nil
}

var repository *sqliteUserRepository

func NewSqliteUserRepository() sqliteUserRepository {
	if repository != nil {
		return *repository
	}

	repository = &sqliteUserRepository{DB: database.GetDatabase()}

	return *repository
}
