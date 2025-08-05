package domain

import "regexp"

type filePath struct {
	value string
}

func (f filePath) Value() (string, error) {
	regex := regexp.MustCompile(`^images/[a-zA-Z0-9_\-]+/uploads/[a-zA-Z0-9\-]+_[a-zA-Z0-9_\-\.]+\.(jpg|jpeg|png|gif|bmp|webp|tiff)$`)

	if !regex.MatchString(f.value) {
		return "", InvalidFilePath{}
	}
	return f.value, nil
}

func BuildFilePath(value string) filePath {
	return filePath{value: value}
}
