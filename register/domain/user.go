package domain

import (
	"github.com/S3ergio31/image-processing-service/shared/domain"
)

type User interface {
	Username() string
	Password() string
}

type user struct {
	username string
	password string
}

func (u user) Username() string {
	return u.username
}

func (u user) Password() string {
	return u.password
}

func NewUser(username string, password string, hasher Hasher) (User, []error) {
	errors := []error{}
	hashedPassword, hashedPasswordErr := Password{value: password, Hasher: hasher}.Value()
	validUsername, usernameErr := domain.BuildUsername(username).Value()
	hashedPassword, passwordErr := domain.BuildHashedPassword(hashedPassword).Value()

	if hashedPasswordErr != nil {
		errors = append(errors, hashedPasswordErr)
	}

	if usernameErr != nil {
		errors = append(errors, usernameErr)
	}

	if passwordErr != nil {
		errors = append(errors, passwordErr)
	}

	if len(errors) != 0 {
		return nil, errors
	}

	return &user{username: validUsername, password: hashedPassword}, errors
}
