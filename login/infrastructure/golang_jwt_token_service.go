package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type GolangJwtTokenService struct{}

func (t GolangJwtTokenService) Generate(secret string, username string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "image-processing-service-app",
		"aud": "user",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
