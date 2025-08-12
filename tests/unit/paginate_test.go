package unit

import (
	"fmt"
	"testing"

	"github.com/S3ergio31/image-processing-service/paginate/application"
	"github.com/S3ergio31/image-processing-service/paginate/domain"
)

type PaginateImageRepository struct {
	images []domain.Image
}

func (u PaginateImageRepository) Get(imagePagination domain.ImagePagination) []domain.Image {
	return []domain.Image{u.images[imagePagination.Page()-1]}
}

func TestPaginateImages(t *testing.T) {
	username := "test"
	imageName := "image"
	imageType := "png"
	image1Uuid := "ed8671d9-7129-49b2-84c1-7c6e4d51f716"
	image2Uuid := "6f4a7da6-1761-4a0f-b364-d69ffbb3ddd7"
	image1, _ := domain.NewImage(
		image1Uuid,
		fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, image1Uuid, imageName, imageType),
		imageName,
		imageType,
	)
	image2, _ := domain.NewImage(
		image2Uuid,
		fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, image2Uuid, imageName, imageType),
		imageName,
		imageType,
	)
	paginator := application.Paginator{
		ImageRepository: PaginateImageRepository{
			images: []domain.Image{image1, image2},
		},
	}

	imagesUuid := map[int]string{
		1: image1Uuid,
		2: image2Uuid,
	}

	for _, page := range []int{1, 2} {
		imagePagination := domain.NewImagePagination("Test", page, 1)

		images := paginator.Paginate(imagePagination)

		if images[0].Uuid().String() != imagesUuid[page] {
			t.Errorf("expected image=%s current=%s images=%s", imagesUuid[page], images[0].Uuid(), images)
		}
	}
}
