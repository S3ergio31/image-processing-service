package infrastructure

import (
	"errors"

	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"github.com/S3ergio31/image-processing-service/transform/domain"
	"gorm.io/gorm"
)

type sqliteImageRepository struct {
	*gorm.DB
}

func (repository sqliteImageRepository) Find(uuid string, username string) (domain.Image, *domain.ImageNotFound) {
	var image *entities.Image
	result := repository.InnerJoins("User").Where("uuid = ?", uuid).Where("username = ?", username).First(&image)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &domain.ImageNotFound{}
	}

	return domain.NewImage(image.Uuid, image.Path, image.User.Username, image.Name), nil
}

var repository *sqliteImageRepository

func NewSqliteRegisterImageRepository() sqliteImageRepository {
	if repository != nil {
		return *repository
	}

	repository = &sqliteImageRepository{DB: database.GetDatabase()}

	return *repository
}
