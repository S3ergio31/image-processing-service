package transformers

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
)

type Rotator struct {
}

func (c Rotator) Transform(transformations *transformations.Transformations, imageEditor ImageEditor, content []byte) []byte {
	return imageEditor.Rotate(transformations.Rotate, content)
}
