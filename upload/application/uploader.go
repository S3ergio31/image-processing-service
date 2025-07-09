package application

import (
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/upload/domain"
)

type Uploader struct {
	domain.ImageRepository
	*events.EventBus
}

func (u *Uploader) Upload(uploadImage domain.UploadImage) []error {
	image, errors := domain.NewImage(uploadImage)

	if len(errors) != 0 {
		return errors
	}

	u.Save(image)

	u.Publish(events.ImageUploaded, events.ImageUploadedEvent{
		Uuid:     image.Uuid().String(),
		Username: image.Username(),
		Name:     image.Name(),
		Path:     image.Path(),
	})

	return errors
}
