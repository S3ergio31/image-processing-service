package domain

type ImageRepository interface {
	Save(image Image)
}
