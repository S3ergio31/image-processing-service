package domain

import "regexp"

type Uuid struct {
	value string
}

func (u Uuid) Value() (string, error) {
	regex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	if !regex.MatchString(u.value) {
		return "", InvalidUuid{}
	}

	return u.value, nil
}

func BuildUuid(value string) Uuid {
	return Uuid{value: value}
}
