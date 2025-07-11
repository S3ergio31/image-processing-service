package application

import (
	"github.com/S3ergio31/image-processing-service/transform/domain"
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
	"github.com/S3ergio31/image-processing-service/transform/domain/transformers"
)

type Transformer struct {
	ImageEditor transformers.ImageEditor
}

func (a Transformer) Transform(transformations *transformations.Transformations) {
	transformers := []transformers.Transformer{
		transformers.Croper{},
		transformers.Filter{},
		transformers.Formater{},
		transformers.Resizer{},
		transformers.Rotator{},
	}
	chain := domain.TransformationChain{
		Transformers: transformers,
		ImageEditor:  a.ImageEditor,
	}
	chain.Transform(transformations, []byte{})
}
