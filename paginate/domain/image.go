package domain

import (
	"github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/google/uuid"
)

type Image interface {
	Uuid() uuid.UUID
	Name() string
	Path() string
	Type() string
}

type image struct {
	uuid      uuid.UUID
	name      string
	imageType string
	path      string
}

func (i image) Uuid() uuid.UUID {
	return i.uuid
}

func (i image) Name() string {
	return i.name
}

func (i image) Path() string {
	return i.path
}

func (i image) Type() string {
	return i.imageType
}

func NewImage(id string, path string, name string, imageType string) (Image, []error) {
	errors := []error{}
	uuid, uuidErr := uuid.Parse(id)
	imagePath, pathErr := domain.BuildFilePath(path).Value()
	imageName, nameErr := domain.BuildImageName(name).Value()
	imageType, typeErr := domain.BuildImageType(imageType).Value()

	if uuidErr != nil {
		errors = append(errors, uuidErr)
	}

	if pathErr != nil {
		errors = append(errors, pathErr)
	}

	if nameErr != nil {
		errors = append(errors, nameErr)
	}

	if typeErr != nil {
		errors = append(errors, typeErr)
	}

	if len(errors) != 0 {
		return nil, errors
	}

	return &image{
		uuid:      uuid,
		path:      imagePath,
		name:      imageName,
		imageType: imageType,
	}, nil
}
