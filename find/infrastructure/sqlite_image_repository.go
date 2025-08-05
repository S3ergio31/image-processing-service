package infrastructure

import (
	"errors"

	"github.com/S3ergio31/image-processing-service/find/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/gorm"
)

type sqliteImageRepository struct {
	*gorm.DB
}

func (repository sqliteImageRepository) FindByUsername(username string, uuid string) (domain.Image, []error) {
	var image *entities.Image
	result := repository.InnerJoins("User").Where("uuid = ?", uuid).Where("username = ?", username).First(&image)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, []error{&domain.ImageNotFound{}}
	}

	return domain.NewImage(image.Uuid, image.Path, image.Name, image.Type)
}

var repository *sqliteImageRepository

func NewSqliteRegisterImageRepository() sqliteImageRepository {
	if repository != nil {
		return *repository
	}

	repository = &sqliteImageRepository{DB: database.GetDatabase()}

	return *repository
}
