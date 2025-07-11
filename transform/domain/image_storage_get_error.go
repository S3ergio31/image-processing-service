package domain

type ImageStorageGetError struct {
}

func (ImageStorageGetError) Error() string {
	return "cannot retrieve image from storage"
}
