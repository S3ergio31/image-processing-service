package transformers

import "github.com/S3ergio31/image-processing-service/transform/domain/transformations"

type ImageEditor interface {
	Resize(transformations.Resize, []byte) []byte
	Crop(transformations.Crop, []byte) []byte
	Filter(transformations.Filters, []byte) []byte
	Rotate(int, []byte) []byte
	Format(string, []byte) []byte
}
