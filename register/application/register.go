package application

import (
	"github.com/S3ergio31/image-processing-service/register/domain"
)

type Register struct {
	domain.UserRepository
	domain.Hasher
}

func (r Register) Store(username string, password string) []error {
	if r.UsedUsername(username) {
		return []error{domain.UserAlreadyExists{}}
	}

	userToRegister, errors := domain.NewUser(username, password, r.Hasher)

	if len(errors) != 0 {
		return errors
	}

	r.Save(userToRegister)

	return []error{}
}
