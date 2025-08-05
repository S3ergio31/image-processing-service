package domain

type InvalidImageType struct {
}

func (InvalidImageType) Error() string {
	return "Invalid image type, should be: jpg|jpeg|png|gif|bmp|webp|tiff"
}
