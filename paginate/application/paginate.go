package application

import (
	"log"

	"github.com/S3ergio31/image-processing-service/paginate/domain"
)

type Paginator struct {
	domain.ImageRepository
}

func (f Paginator) Paginate(imagePagination domain.ImagePagination) []domain.Image {
	log.Printf("Paginator.Paginate\n -> username=%s limit=%d page=%d", imagePagination.Username(), imagePagination.Limit(), imagePagination.Page())
	return f.ImageRepository.Get(imagePagination)
}
