package domain

type imageName struct {
	value string
}

func (i imageName) Value() (string, error) {
	if i.value == "" {
		return "", InvalidImageName{}
	}
	return i.value, nil
}

func BuildImageName(value string) imageName {
	return imageName{value: value}
}
