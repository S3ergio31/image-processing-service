package application

import (
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/transform/domain"
)

type Transformer struct {
	domain.ImageEditor
	domain.ImageStorage
	domain.ImageRepository
	*events.EventBus
}

func (t Transformer) Transform(transformations *domain.Transformations) error {
	transformers := []domain.Transformer{
		domain.Croper{},
		domain.Formater{},
		domain.Resizer{},
		domain.Rotator{},
	}

	chain := domain.TransformationChain{
		Transformers: transformers,
		ImageEditor:  t.ImageEditor,
	}

	image, err := t.Find(transformations.ImageUuid, transformations.Username)

	if err != nil {
		return err
	}

	content, imageStorageError := t.Get(image.Path())

	if imageStorageError != nil {
		return imageStorageError
	}

	result := chain.Transform(transformations, content)

	imageSaveError := t.Storage(domain.Transformation{
		Transformations: transformations,
		Image:           image,
		Content:         result,
	})

	if imageSaveError != nil {
		return imageSaveError
	}

	t.Publish(events.ImageTransformed, events.ImageTransformedEvent{})

	return nil
}
