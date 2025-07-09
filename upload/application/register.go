package application

import (
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/upload/domain"
)

type Register struct {
	domain.RegisterImageRepository
}

func (u *Register) Store(event events.ImageUploadedEvent) {
	u.Save(event)
}
