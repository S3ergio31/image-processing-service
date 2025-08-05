package unit

import (
	"fmt"
	"testing"

	"github.com/S3ergio31/image-processing-service/find/application"
	"github.com/S3ergio31/image-processing-service/find/domain"
	"github.com/google/uuid"
)

type FindImageRepository struct {
	Id        string
	Path      string
	Name      string
	ImageType string
	NotFound  bool
}

func (u FindImageRepository) FindByUsername(username string, uuid string) (domain.Image, []error) {
	if u.NotFound {
		return nil, []error{&domain.ImageNotFound{}}
	}
	return domain.NewImage(u.Id, u.Path, u.Name, u.ImageType)
}

func TestFindImage(t *testing.T) {
	username := "test"
	imageName := "image"
	imageUuid := uuid.NewString()
	imageType := "png"
	finder := application.Finder{
		ImageRepository: FindImageRepository{
			Id:        imageUuid,
			Path:      fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, imageUuid, imageName, imageType),
			Name:      imageName,
			ImageType: imageType,
		},
	}

	image, errors := finder.Find(username, imageUuid)

	if image == nil && errors != nil {
		t.Errorf("expected non-error")
	}
}

func TestImageCannotBeNotFound(t *testing.T) {
	username := "test"
	imageUuid := uuid.NewString()
	finder := application.Finder{
		ImageRepository: FindImageRepository{NotFound: true},
	}

	_, errors := finder.Find(username, imageUuid)

	if _, ok := errors[0].(*domain.ImageNotFound); !ok {
		t.Errorf("Expected domain.ImageNotFound, given %s", errors)
	}
}

func TestInvalidImageData(t *testing.T) {
	username := "test"
	imageName := ""
	imageUuid := ""
	imageType := ""
	finder := application.Finder{
		ImageRepository: FindImageRepository{
			Id:        imageUuid,
			Path:      fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, imageUuid, imageName, imageType),
			Name:      imageName,
			ImageType: imageType,
		},
	}

	_, errors := finder.Find(username, imageUuid)

	if len(errors) != 4 {
		t.Errorf("expected 4 error, given %d -> %s", len(errors), errors)
	}
}
