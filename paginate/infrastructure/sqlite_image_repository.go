package infrastructure

import (
	"log"

	"github.com/S3ergio31/image-processing-service/paginate/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"gorm.io/gorm"
)

type sqliteImageRepository struct {
	*gorm.DB
}

func (repository sqliteImageRepository) Get(imagePagination domain.ImagePagination) []domain.Image {
	var entities []entities.Image
	images := []domain.Image{}

	offset := (imagePagination.Page() - 1) * imagePagination.Limit()
	repository.DB.
		InnerJoins("User").
		Where("username = ?", imagePagination.Username()).
		Offset(offset).
		Limit(imagePagination.Limit()).
		Find(&entities)

	for _, entity := range entities {
		image, errors := domain.NewImage(entity.Uuid, entity.Path, entity.Name, entity.Type)

		if len(errors) != 0 {
			log.Println("Pagination errors: ", errors)
			continue
		}

		images = append(images, image)
	}

	return images
}

var repository *sqliteImageRepository

func NewSqlitePaginateImageRepository() sqliteImageRepository {
	if repository != nil {
		return *repository
	}

	repository = &sqliteImageRepository{DB: database.GetDatabase()}

	return *repository
}
