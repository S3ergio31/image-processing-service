package domain

type imageContent struct {
	value []byte
}

func (u imageContent) Value() ([]byte, error) {
	if len(u.value) == 0 {
		return nil, InvalidImageContent{}
	}
	return u.value, nil
}

func BuildImageContent(value []byte) imageContent {
	return imageContent{value: value}
}
