package application

import (
	"github.com/S3ergio31/image-processing-service/login/domain"
)

type Auth struct {
	domain.UserRepository
	domain.TokenService
	Secret string
}

func (a Auth) Login(username string, password string) (string, error) {
	user, notFoundErr := a.FindByUsername(username)

	if notFoundErr != nil || !user.Check(password) {
		return "", domain.InvalidCredentials{}
	}

	token, err := a.Generate(a.Secret, username)

	if err != nil {
		return "", domain.InvalidTokenGeneration{Err: err}
	}

	user.SetToken(token)

	a.Save(user)

	return token, nil

}
