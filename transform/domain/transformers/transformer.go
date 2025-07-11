package transformers

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
)

type Transformer interface {
	Transform(*transformations.Transformations, ImageEditor, []byte) []byte
}
