package domain

import (
	"errors"
	"log"
	"regexp"
)

type HashedPassword struct {
	value string
}

func (p HashedPassword) Value() (string, error) {
	regex := regexp.MustCompile(`^\$2[abxy]\$\d{2}\$[./A-Za-z0-9]{22}[./A-Za-z0-9]{31}$`)

	if !regex.MatchString(p.value) {
		err := "invalid hashed password"
		log.Println("HashedPassword.Value", err)
		return "", errors.New(err)
	}

	return p.value, nil
}

func BuildHashedPassword(value string) HashedPassword {
	return HashedPassword{value: value}
}
