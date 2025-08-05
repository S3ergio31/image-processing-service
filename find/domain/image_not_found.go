package domain

type ImageNotFound struct {
}

func (ImageNotFound) Error() string {
	return "image not found"
}
