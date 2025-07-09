package domain

import "regexp"

type imageName struct {
	value string
}

func (u imageName) Value() (string, error) {
	regex := regexp.MustCompile(`(?i)^[a-zA-Z0-9_\-\.]+\.(jpg|jpeg|png|gif|bmp|webp|tiff)$`)

	if !regex.MatchString(u.value) {
		return "", InvalidImageName{}
	}
	return u.value, nil
}

func BuildImageName(value string) imageName {
	return imageName{value: value}
}
