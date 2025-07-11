package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/utils"
	"github.com/S3ergio31/image-processing-service/upload/domain"
)

type LocalDiskImageRepository struct {
}

func (LocalDiskImageRepository) Save(image domain.Image) {
	utils.SaveOnDisk(image.Path(), image.Content())
}
