package unit

import (
	"testing"

	shared "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/upload/application"
)

const testUuid = "d3e984d3-3098-404d-a71b-f2d25b26cfbe"

type RegisterImageRepository struct {
	t *testing.T
}

func (u RegisterImageRepository) Save(event shared.ImageUploadedEvent) {
	if event.Uuid != testUuid {
		u.t.Errorf("invalid event, expected uuid : %s given: %s", testUuid, event.Uuid)
	}
}

func TestUploadImageRegister(t *testing.T) {
	register := application.Register{
		RegisterImageRepository: RegisterImageRepository{t: t},
	}

	register.Store(shared.ImageUploadedEvent{
		Uuid:     testUuid,
		Username: "Test",
		Name:     "test.png",
		Path:     "images/Test/d3e984d3-3098-404d-a71b-f2d25b26cfbe_test.png",
	})
}
