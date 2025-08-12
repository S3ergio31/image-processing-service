package domain

type ImageRepository interface {
	Get(imagePagination ImagePagination) []Image
}
