package domain

import (
	shared "github.com/S3ergio31/image-processing-service/shared/domain"
)

type HashedPassword struct {
	shared.HashedPassword
	Hasher
}

func (p HashedPassword) Check(password string) bool {
	value, _ := p.Value()
	err := p.CompareHashAndPassword(value, password)

	return err == nil
}
