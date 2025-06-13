package application

import (
	"github.com/S3ergio31/image-processing-service/register/domain"
)

type Register struct {
	Repository domain.UserRepository
}

func (r Register) Save(username string, password string) []error {
	if r.Repository.UsedUsername(username) {
		return []error{domain.UserAlreadyExists{}}
	}

	userToRegister, errors := domain.NewUser(username, password)

	if len(errors) != 0 {
		return errors
	}

	r.Repository.Save(userToRegister)

	return []error{}
}
