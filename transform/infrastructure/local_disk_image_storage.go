package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/utils"
	"github.com/S3ergio31/image-processing-service/transform/domain"
	"github.com/h2non/bimg"
)

type LocalDiskImageStorage struct {
}

func (LocalDiskImageStorage) Get(path string) ([]byte, *domain.ImageStorageGetError) {
	image, err := bimg.Read(path)

	if err != nil {
		utils.LogError(err)
		return nil, &domain.ImageStorageGetError{}
	}

	return image, nil
}

func (LocalDiskImageStorage) Storage(transformation domain.Transformation) *domain.ImageStorageSaveError {
	err := utils.SaveOnDisk(transformation.Path(), transformation.Content)

	if err != nil {
		return &domain.ImageStorageSaveError{}
	}
	return nil
}
