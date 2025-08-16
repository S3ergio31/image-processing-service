package domain

import "github.com/S3ergio31/image-processing-service/shared/domain"

type User interface {
	Username() string
	Check(password string) bool
	SetToken(token string)
	Token() string
}

type user struct {
	username string
	password HashedPassword
	token    string
}

func (u user) Username() string {
	return u.username
}

func (u user) Check(password string) bool {
	return u.password.Check(password)
}

func (u *user) SetToken(token string) {
	u.token = token
}

func (u user) Token() string {
	return u.token
}

func NewUser(username string, hashedPassword string, hasher Hasher) (User, []error) {
	errors := []error{}

	hashedPasswordValueObject := HashedPassword{
		HashedPassword: domain.BuildHashedPassword(hashedPassword),
		Hasher:         hasher,
	}

	validUsername, usernameErr := domain.BuildUsername(username).Value()
	_, hashedPasswordErr := hashedPasswordValueObject.Value()

	if usernameErr != nil {
		errors = append(errors, usernameErr)
	}

	if hashedPasswordErr != nil {
		errors = append(errors, hashedPasswordErr)
	}

	if len(errors) != 0 {
		return nil, errors
	}

	return &user{username: validUsername, password: hashedPasswordValueObject}, errors
}
