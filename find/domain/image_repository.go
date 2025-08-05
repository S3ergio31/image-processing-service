package domain

type ImageRepository interface {
	FindByUsername(username string, uuid string) (Image, []error)
}
