package domain

import "github.com/S3ergio31/image-processing-service/shared/domain"

type RegisterImageRepository interface {
	Save(event domain.ImageUploadedEvent)
}
