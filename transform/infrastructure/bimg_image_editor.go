package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/utils"
	transformations "github.com/S3ergio31/image-processing-service/transform/domain"
	"github.com/h2non/bimg"
)

type BimgImageEditor struct{}

func (i BimgImageEditor) Resize(resize *transformations.Resize, content []byte) []byte {
	result, err := bimg.NewImage(content).Resize(resize.Width, resize.Height)

	utils.LogError(err)

	return result
}

func (i BimgImageEditor) Crop(crop *transformations.Crop, content []byte) []byte {
	result, err := bimg.NewImage(content).Extract(crop.Y, crop.X, crop.Width, crop.Height)

	utils.LogError(err)

	return result
}

func (i BimgImageEditor) Rotate(angle int, content []byte) []byte {
	image := bimg.NewImage(content)
	var result []byte
	var err error

	switch angle {
	case 0:
		result, err = image.Rotate(0)
	case 90:
		result, _ = image.Rotate(90)
	case 180:
		result, _ = image.Rotate(180)
	case 270:
		result, _ = image.Rotate(270)
	}

	utils.LogError(err)

	return result
}

func (i BimgImageEditor) Format(imageType string, content []byte) []byte {
	types := map[string]bimg.ImageType{
		"png":  bimg.PNG,
		"jpeg": bimg.JPEG,
		"gif":  bimg.GIF,
	}

	result, err := bimg.NewImage(content).Convert(types[imageType])

	utils.LogError(err)

	return result
}
