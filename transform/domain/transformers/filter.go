package transformers

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
)

type Filter struct {
}

func (c Filter) Transform(transformations *transformations.Transformations, imageEditor ImageEditor, content []byte) []byte {
	if transformations.Filters == nil {
		return nil
	}

	return imageEditor.Filter(*transformations.Filters, content)
}
