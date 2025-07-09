package domain

import (
	"fmt"

	"github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/google/uuid"
)

type UploadImage struct {
	Uuid     string
	Username string
	Name     string
	Content  []byte
}

type Image interface {
	Uuid() uuid.UUID
	Username() string
	Name() string
	Content() []byte
	Path() string
}

type image struct {
	uuid     uuid.UUID
	username string
	name     string
	content  []byte
}

func (i image) Uuid() uuid.UUID {
	return i.uuid
}

func (i image) Username() string {
	return i.username
}

func (i image) Name() string {
	return i.name
}

func (i image) Content() []byte {
	return i.content
}

func (i image) Path() string {
	return fmt.Sprintf("images/%s/%s_%s", i.username, i.uuid, i.name)
}

func NewImage(uploadImage UploadImage) (Image, []error) {
	errors := []error{}
	username, usernameErr := domain.BuildUsername(uploadImage.Username).Value()
	imageName, imageNameErr := BuildImageName(uploadImage.Name).Value()
	uuid, uuidErr := uuid.Parse(uploadImage.Uuid)
	content, contentErr := BuildImageContent(uploadImage.Content).Value()

	if usernameErr != nil {
		errors = append(errors, usernameErr)
	}

	if imageNameErr != nil {
		errors = append(errors, imageNameErr)
	}

	if uuidErr != nil {
		errors = append(errors, uuidErr)
	}

	if contentErr != nil {
		errors = append(errors, contentErr)
	}

	if len(errors) != 0 {
		return nil, errors
	}

	return &image{
		uuid:     uuid,
		username: username,
		name:     imageName,
		content:  content,
	}, nil
}
