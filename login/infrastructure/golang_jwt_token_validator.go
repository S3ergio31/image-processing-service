package infrastructure

import (
	"log"

	"github.com/golang-jwt/jwt"
)

type GolangJwtTokenValidator struct {
}

func (t GolangJwtTokenValidator) getJwtToken(secret string, token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
}

func (t GolangJwtTokenValidator) Username(secret string, token string) string {
	jwtToken, _ := t.getJwtToken(secret, token)
	claims := jwtToken.Claims.(jwt.MapClaims)

	return claims["sub"].(string)
}

func (t GolangJwtTokenValidator) IsValid(secret string, token string) bool {
	jwtToken, err := t.getJwtToken(secret, token)

	if err != nil {
		log.Printf("tokenValidator -> IsValid -> Error: %s\n", err.Error())
		return false
	}

	if !jwtToken.Valid {
		log.Printf("tokenValidator -> IsValid -> Invalid token:\n")
		return false
	}

	return true
}
