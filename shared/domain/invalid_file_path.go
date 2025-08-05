package domain

type InvalidFilePath struct {
}

func (InvalidFilePath) Error() string {
	return "invalid file path, should be like 'path/name.jpg|jpeg|png|gif|bmp|webp|tiff'"
}
