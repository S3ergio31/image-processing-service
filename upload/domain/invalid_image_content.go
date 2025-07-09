package domain

type InvalidImageContent struct {
}

func (InvalidImageContent) Error() string {
	return "image content cannot be empty"
}
