package domain

import (
	"log"

	"github.com/golang-jwt/jwt"
)

type TokenValidator struct {
	Secret string
}

func (t TokenValidator) getJwtToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})
}

func (t TokenValidator) Username(token string) string {
	jwtToken, _ := t.getJwtToken(token)
	claims := jwtToken.Claims.(jwt.MapClaims)

	return claims["sub"].(string)
}

func (t TokenValidator) IsValid(token string) bool {
	jwtToken, err := t.getJwtToken(token)

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
