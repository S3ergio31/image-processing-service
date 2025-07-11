package transformers

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
)

type Resizer struct {
}

func (c Resizer) Transform(transformations *transformations.Transformations, imageEditor ImageEditor, content []byte) []byte {
	if transformations.Resize == nil {
		return nil
	}

	return imageEditor.Resize(*transformations.Resize, content)
}
