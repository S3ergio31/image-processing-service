package domain

type ImageRepository interface {
	Find(uuid string, username string) (Image, *ImageNotFound)
}
