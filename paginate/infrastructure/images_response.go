package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/paginate/domain"
	"github.com/google/uuid"
)

type ImageResponse struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	ImageType string    `json:"type"`
	Path      string    `json:"path"`
}

func ToImageResponse(images []domain.Image) []ImageResponse {
	imageResponses := []ImageResponse{}

	for _, image := range images {
		imageResponses = append(imageResponses, ImageResponse{
			Uuid:      image.Uuid(),
			Name:      image.Name(),
			ImageType: image.Type(),
			Path:      image.Path(),
		})
	}

	return imageResponses
}
