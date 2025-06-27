package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	Secret string
}

func (t TokenService) Generate(username string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "image-processing-service-app",
		"aud": "user",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(t.Secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
