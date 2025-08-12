package domain

import "regexp"

type imageType struct {
	value string
}

func (f imageType) Value() (string, error) {
	regex := regexp.MustCompile(`^(jpg|jpeg|png|gif|bmp|webp|tiff)$`)

	if !regex.MatchString(f.value) {
		return "", InvalidImageType{}
	}
	return f.value, nil
}

func BuildImageType(value string) imageType {
	return imageType{value: value}
}
