package domain

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type HashedPassword struct {
	value string
}

func (p HashedPassword) Value() (string, error) {
	_, err := bcrypt.Cost([]byte(p.value))

	if err != nil {
		log.Println("HashedPassword.Value", err)
		return "", errors.New("invalid hashed password")
	}

	return p.value, nil
}

func BuildHashedPassword(value string) HashedPassword {
	return HashedPassword{value: value}
}
