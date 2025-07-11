package transformers

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
)

type Croper struct {
}

func (c Croper) Transform(transformations *transformations.Transformations, imageEditor ImageEditor, content []byte) []byte {
	if transformations.Crop == nil {
		return nil
	}

	return imageEditor.Crop(*transformations.Crop, content)
}
