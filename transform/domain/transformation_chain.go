package domain

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
	"github.com/S3ergio31/image-processing-service/transform/domain/transformers"
)

type TransformationChain struct {
	Transformers []transformers.Transformer
	ImageEditor  transformers.ImageEditor
}

func (t *TransformationChain) Transform(transformations *transformations.Transformations, content []byte) []byte {
	result := content

	for _, transformer := range t.Transformers {
		content = transformer.Transform(transformations, t.ImageEditor, content)
		if content == nil {
			continue
		}
		result = content
	}

	return result
}
