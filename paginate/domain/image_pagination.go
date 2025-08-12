package domain

type ImagePagination interface {
	Username() string
	Page() int
	Limit() int
}

type imagePagination struct {
	page     int
	limit    int
	username string
}

func (i imagePagination) Username() string {
	return i.username
}

func (i imagePagination) Page() int {
	if i.page <= 0 {
		i.page = 1
	}
	return i.page
}

func (i imagePagination) Limit() int {
	switch {
	case i.limit > 100:
		i.limit = 100
	case i.limit <= 0:
		i.limit = 10
	}
	return i.limit
}

func NewImagePagination(username string, page int, limit int) ImagePagination {
	return imagePagination{
		page:     page,
		limit:    limit,
		username: username,
	}
}
