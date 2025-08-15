package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
}

func (b BcryptHasher) Hash(value string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
