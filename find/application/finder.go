package application

import (
	"github.com/S3ergio31/image-processing-service/find/domain"
)

type Finder struct {
	domain.ImageRepository
}

func (f Finder) Find(username string, uuid string) (domain.Image, []error) {
	return f.FindByUsername(username, uuid)
}
