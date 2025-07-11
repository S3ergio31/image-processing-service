package domain

type Image interface {
	Uuid() string
	Path() string
	Name() string
	Username() string
}

type image struct {
	uuid     string
	path     string
	name     string
	username string
}

func (i image) Uuid() string {
	return i.uuid
}

func (i image) Path() string {
	return i.path
}

func (i image) Name() string {
	return i.name
}

func (i image) Username() string {
	return i.username
}

func NewImage(uuid string, path string, username string, name string) Image {
	return image{
		uuid:     uuid,
		path:     path,
		username: username,
		name:     name,
	}
}
