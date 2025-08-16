package infrastructure

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {
}

func (b BcryptHasher) CompareHashAndPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
