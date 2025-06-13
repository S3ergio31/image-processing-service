package domain

type Username struct {
	value string
}

func (u Username) Value() (string, error) {
	if u.value == "" {
		return "", UsernameCannotBeEmpty{}
	}
	return u.value, nil
}

func BuildUsername(value string) Username {
	return Username{value: value}
}
