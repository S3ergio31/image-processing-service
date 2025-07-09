package domain

type InvalidImageName struct {
}

func (InvalidImageName) Error() string {
	return "invalid image name, should be like 'name.jpg|jpeg|png|gif|bmp|webp|tiff'"
}
