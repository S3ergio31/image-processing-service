package infrastructure

import (
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/gorm"
)

type sqliteRegisterImageRepository struct {
	*gorm.DB
}

func (repository sqliteRegisterImageRepository) Save(event events.ImageUploadedEvent) {
	var user entities.User
	repository.First(&user, "username = ?", event.Username)

	repository.Create(&entities.Image{
		Uuid: event.Uuid,
		Name: event.Name,
		Path: event.Path,
		User: user,
		Type: event.Type,
	})
}

var repository *sqliteRegisterImageRepository

func NewSqliteRegisterImageRepository() sqliteRegisterImageRepository {
	if repository != nil {
		return *repository
	}

	repository = &sqliteRegisterImageRepository{DB: database.GetDatabase()}

	return *repository
}
