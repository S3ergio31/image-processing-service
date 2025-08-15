package domain

import (
	"log"
	"regexp"
)

type Password struct {
	Hasher
	value string
}

func (p Password) Value() (string, error) {
	if !p.isValid() {
		return "", InvalidPassword{}
	}

	hash, err := p.Hash(p.value)

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
