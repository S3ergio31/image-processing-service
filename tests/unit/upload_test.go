package unit

import (
	"testing"

	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/upload/application"
	"github.com/S3ergio31/image-processing-service/upload/domain"
	"github.com/google/uuid"
)

type ImageRepository struct{}

func (u ImageRepository) Save(user domain.Image) {}

func TestUploadImage(t *testing.T) {
	eventBus := events.New()

	eventBus.Subscribe(events.ImageUploaded, func(e events.Event) {
		if _, ok := e.(events.ImageUploadedEvent); !ok {
			t.Errorf("event ImageUploadedEvent not fired")
		}
	})

	uploader := application.Uploader{
		ImageRepository: ImageRepository{},
		EventBus:        eventBus,
	}

	err := uploader.Upload(domain.UploadImage{
		Uuid:     uuid.NewString(),
		Username: "Test",
		Name:     "test.png",
		Content:  make([]byte, 1),
	})

	if err != nil {
		t.Errorf("expected non-error")
	}
}

func TestInvalidImage(t *testing.T) {
	eventBus := events.New()

	eventBus.Subscribe(events.ImageUploaded, func(e events.Event) {
		if _, ok := e.(events.ImageUploadedEvent); ok {
			t.Errorf("event ImageUploadedEvent fired")
		}
	})

	uploader := application.Uploader{
		ImageRepository: ImageRepository{},
		EventBus:        eventBus,
	}

	errors := uploader.Upload(domain.UploadImage{
		Uuid:     "",
		Username: "",
		Name:     "",
		Content:  make([]byte, 0),
	})

	if len(errors) != 4 {
		t.Errorf("expected 4 error")
	}
}
