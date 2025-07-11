package domain

type ImageStorageSaveError struct {
}

func (ImageStorageSaveError) Error() string {
	return "cannot save image to storage"
}
