package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
	"github.com/h2non/bimg"
)

type ImageEditor struct{}

func (i *ImageEditor) Resize(resize transformations.Resize, content []byte) []byte {
	result, _ := bimg.NewImage(content).Resize(resize.Width, resize.Height)

	return result
}

func (i *ImageEditor) Crop(crop transformations.Crop, content []byte) []byte {
	result, _ := bimg.NewImage(content).Extract(crop.Y, crop.X, crop.Width, crop.Height)

	return result
}

func (i *ImageEditor) Filter(filters transformations.Filters, content []byte) []byte {
	return nil
}

func (i *ImageEditor) Rotate(angle int, content []byte) []byte {
	image := bimg.NewImage(content)
	var result []byte

	switch angle {
	case 0:
		result, _ = image.Rotate(0)
	case 90:
		result, _ = image.Rotate(90)
	case 180:
		result, _ = image.Rotate(180)
	case 270:
		result, _ = image.Rotate(270)
	}

	return result
}

func (i *ImageEditor) Format(imageType string, content []byte) []byte {
	types := map[string]bimg.ImageType{
		"png":  bimg.PNG,
		"jpeg": bimg.JPEG,
		"gif":  bimg.GIF,
	}

	result, _ := bimg.NewImage(content).Convert(types[imageType])

	return result
}
