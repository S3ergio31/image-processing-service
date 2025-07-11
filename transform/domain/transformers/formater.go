package transformers

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
)

type Formater struct {
}

func (c Formater) Transform(transformations *transformations.Transformations, imageEditor ImageEditor, content []byte) []byte {
	return imageEditor.Format(transformations.Format, content)
}
