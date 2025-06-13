package infrastructure

import (
	"log"

	"github.com/S3ergio31/image-processing-service/register/domain"
)

type inMemoryUserRepository struct {
	users map[string]domain.User
}

func (repository inMemoryUserRepository) Save(user domain.User) {
	repository.users[user.Username()] = user
	log.Println(user)
}

func (repository inMemoryUserRepository) UsedUsername(username string) bool {
	_, ok := repository.users[username]

	return ok
}

var repository *inMemoryUserRepository

func NewInMemoryUserRepository() inMemoryUserRepository {
	if repository == nil {
		repository = &inMemoryUserRepository{users: make(map[string]domain.User)}
	}

	return *repository
}
