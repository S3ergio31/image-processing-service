package domain

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func (p Password) Value() (string, error) {
	if !p.isValid() {
		return "", InvalidPassword{}
	}

	hash, err := p.hash()

	if err != nil {
		log.Println("Password.Value", err)
		return "", InvalidPassword{}
	}

	return hash, nil
}

func (p Password) isValid() bool {
	if len(p.value) < 8 {
		return false
	}

	hasMayus := regexp.MustCompile(`[A-Z]`).MatchString(p.value)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(p.value)
	hasSpecialChar := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\}\\|;:'",.<>\/?]`).MatchString(p.value)

	return hasMayus && hasNumber && hasSpecialChar
}

func (p Password) hash() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.value), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
