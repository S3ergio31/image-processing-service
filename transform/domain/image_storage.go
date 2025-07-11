package domain

type ImageStorage interface {
	Get(path string) ([]byte, *ImageStorageGetError)
	Storage(transformation Transformation) *ImageStorageSaveError
}
