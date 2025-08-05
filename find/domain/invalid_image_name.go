package domain

type InvalidImageName struct {
}

func (InvalidImageName) Error() string {
	return "Image name cannot be empty"
}
