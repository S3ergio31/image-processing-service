package domain

import (
	"fmt"
	"strings"

	"github.com/S3ergio31/image-processing-service/shared/domain"
)

type UploadImage struct {
	Uuid     string
	Username string
	Name     string
	Content  []byte
}

type Image interface {
	Uuid() string
	Username() string
	Name() string
	Content() []byte
	Path() string
	Type() string
}

type image struct {
	uuid     string
	username string
	name     string
	content  []byte
}

func (i image) Uuid() string {
	return i.uuid
}

func (i image) Username() string {
	return i.username
}

func (i image) Name() string {
	return strings.Split(i.name, ".")[0]
}

func (i image) Content() []byte {
	return i.content
}

func (i image) Path() string {
	return fmt.Sprintf("images/%s/uploads/%s_%s.%s", i.username, i.uuid, i.Name(), i.Type())
}

func (i image) Type() string {
	return strings.Split(i.name, ".")[1]
}

func NewImage(uploadImage UploadImage) (Image, []error) {
	errors := []error{}
	username, usernameErr := domain.BuildUsername(uploadImage.Username).Value()
	imageName, imageNameErr := BuildImageName(uploadImage.Name).Value()
	uuid, uuidErr := domain.BuildUuid(uploadImage.Uuid).Value()
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
