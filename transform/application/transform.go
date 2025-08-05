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

func (t Transformer) Transform(transformations *domain.Transformations) (*domain.Transformation, error) {
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
		return nil, err
	}

	content, imageStorageError := t.Get(image.Path())

	if imageStorageError != nil {
		return nil, imageStorageError
	}

	result := chain.Transform(transformations, content)

	transformation := domain.Transformation{
		Transformations: transformations,
		Image:           image,
		Content:         result,
	}

	imageSaveError := t.Storage(transformation)

	if imageSaveError != nil {
		return nil, imageSaveError
	}

	t.Publish(events.ImageTransformed, events.ImageTransformedEvent{})

	return &transformation, nil
}
