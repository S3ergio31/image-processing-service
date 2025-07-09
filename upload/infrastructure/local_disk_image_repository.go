package infrastructure

import (
	"os"
	"path/filepath"

	"github.com/S3ergio31/image-processing-service/upload/domain"
)

type LocalDiskImageRepository struct {
}

func (LocalDiskImageRepository) Save(image domain.Image) {
	if err := os.MkdirAll(filepath.Dir(image.Path()), 0750); err != nil {
		return
	}

	if err := os.WriteFile(image.Path(), image.Content(), 0644); err != nil {
		return
	}
}
