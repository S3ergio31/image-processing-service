package domain

import (
	shared "github.com/S3ergio31/image-processing-service/shared/domain"
	"golang.org/x/crypto/bcrypt"
)

type HashedPassword struct {
	shared.HashedPassword
}

func (p HashedPassword) Check(password string) bool {
	value, _ := p.Value()
	err := bcrypt.CompareHashAndPassword([]byte(value), []byte(password))

	return err == nil
}
