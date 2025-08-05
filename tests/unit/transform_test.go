package unit

import (
	"fmt"
	"testing"

	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/transform/application"
	"github.com/S3ergio31/image-processing-service/transform/domain"
)

type ImageEditor struct{}

func (i ImageEditor) Resize(resize *domain.Resize, content []byte) []byte {
	return []byte{}
}

func (i ImageEditor) Crop(crop *domain.Crop, content []byte) []byte {
	return []byte{}
}

func (i ImageEditor) Rotate(angle int, content []byte) []byte {
	return []byte{}
}

func (i ImageEditor) Format(imageType string, content []byte) []byte {
	return []byte{}
}

type ImageStorage struct {
	getError     bool
	storageError bool
}

func (i ImageStorage) Get(path string) ([]byte, *domain.ImageStorageGetError) {

	if i.getError {
		return nil, &domain.ImageStorageGetError{}
	}

	return []byte{}, nil
}

func (i ImageStorage) Storage(transformation domain.Transformation) *domain.ImageStorageSaveError {
	if i.storageError {
		return &domain.ImageStorageSaveError{}
	}

	return nil
}

type TransformImageRepository struct {
	image *domain.Image
}

func (r TransformImageRepository) Find(uuid string, username string) (domain.Image, *domain.ImageNotFound) {
	if r.image != nil {
		return *r.image, nil
	}
	return nil, &domain.ImageNotFound{}
}

func TestTransformOk(t *testing.T) {
	eventBus := events.New()

	eventBus.Subscribe(events.ImageTransformed, func(e events.Event) {
		if _, ok := e.(events.ImageTransformedEvent); !ok {
			t.Errorf("event ImageTransformedEvent not fired")
		}
	})

	imageUuid := "e9c1a020-7077-49e3-a217-72c395d371fc"
	name := "test"
	username := "Test"
	image := domain.NewImage(
		imageUuid,
		fmt.Sprintf("images/%s/transformations/%s_%s.png", username, imageUuid, name),
		username,
		name,
	)

	application.Transformer{
		ImageEditor:     ImageEditor{},
		ImageStorage:    ImageStorage{getError: false, storageError: false},
		ImageRepository: TransformImageRepository{image: &image},
		EventBus:        events.New(),
	}.Transform(&domain.Transformations{
		ImageUuid: imageUuid,
		Username:  username,
		Rotate:    0,
		Format:    "png",
		Resize: &domain.Resize{
			Width:  100,
			Height: 100,
		},
		Crop: &domain.Crop{
			Width:  50,
			Height: 50,
			X:      8,
			Y:      10,
		},
	})
}

func TestImageNotFound(t *testing.T) {
	eventBus := events.New()

	eventBus.Subscribe(events.ImageTransformed, func(e events.Event) {
		if _, ok := e.(events.ImageTransformedEvent); ok {
			t.Errorf("event ImageTransformed unexpected")
		}
	})

	imageUuid := "e9c1a020-7077-49e3-a217-72c395d371fc"
	username := "Test"

	_, err := application.Transformer{
		ImageEditor:     ImageEditor{},
		ImageStorage:    ImageStorage{},
		ImageRepository: TransformImageRepository{},
		EventBus:        events.New(),
	}.Transform(&domain.Transformations{
		ImageUuid: imageUuid,
		Username:  username,
		Rotate:    0,
		Format:    "png",
		Resize: &domain.Resize{
			Width:  100,
			Height: 100,
		},
		Crop: &domain.Crop{
			Width:  50,
			Height: 50,
			X:      8,
			Y:      10,
		},
	})

	if _, ok := err.(*domain.ImageNotFound); !ok {
		fmt.Printf("Type of x: %T\n", err)
		t.Errorf("expected error ImageNotFound, given: %s", err)
	}
}

func TestCannotRetrieveImageFromStorage(t *testing.T) {
	eventBus := events.New()

	eventBus.Subscribe(events.ImageTransformed, func(e events.Event) {
		if _, ok := e.(events.ImageTransformedEvent); ok {
			t.Errorf("event ImageTransformed unexpected")
		}
	})

	imageUuid := "e9c1a020-7077-49e3-a217-72c395d371fc"
	name := "test"
	username := "Test"
	image := domain.NewImage(
		imageUuid,
		fmt.Sprintf("images/%s/transformations/%s_%s.png", username, imageUuid, name),
		username,
		name,
	)

	_, err := application.Transformer{
		ImageEditor:     ImageEditor{},
		ImageStorage:    ImageStorage{getError: true},
		ImageRepository: TransformImageRepository{image: &image},
		EventBus:        events.New(),
	}.Transform(&domain.Transformations{
		ImageUuid: imageUuid,
		Username:  username,
		Rotate:    0,
		Format:    "png",
		Resize: &domain.Resize{
			Width:  100,
			Height: 100,
		},
		Crop: &domain.Crop{
			Width:  50,
			Height: 50,
			X:      8,
			Y:      10,
		},
	})

	if _, ok := err.(*domain.ImageStorageGetError); !ok {
		fmt.Printf("Type of x: %T\n", err)
		t.Errorf("expected error ImageStorageGetError, given: %s", err)
	}
}

func TestCannotSaveImageToStorage(t *testing.T) {
	eventBus := events.New()

	eventBus.Subscribe(events.ImageTransformed, func(e events.Event) {
		if _, ok := e.(events.ImageTransformedEvent); ok {
			t.Errorf("event ImageTransformed unexpected")
		}
	})

	imageUuid := "e9c1a020-7077-49e3-a217-72c395d371fc"
	name := "test"
	username := "Test"
	image := domain.NewImage(
		imageUuid,
		fmt.Sprintf("images/%s/transformations/%s_%s.png", username, imageUuid, name),
		username,
		name,
	)

	_, err := application.Transformer{
		ImageEditor:     ImageEditor{},
		ImageStorage:    ImageStorage{getError: false, storageError: true},
		ImageRepository: TransformImageRepository{image: &image},
		EventBus:        events.New(),
	}.Transform(&domain.Transformations{
		ImageUuid: imageUuid,
		Username:  username,
		Rotate:    0,
		Format:    "png",
		Resize: &domain.Resize{
			Width:  100,
			Height: 100,
		},
		Crop: &domain.Crop{
			Width:  50,
			Height: 50,
			X:      8,
			Y:      10,
		},
	})

	if _, ok := err.(*domain.ImageStorageSaveError); !ok {
		fmt.Printf("Type of x: %T\n", err)
		t.Errorf("expected error ImageStorageSaveError, given: %s", err)
	}
}
