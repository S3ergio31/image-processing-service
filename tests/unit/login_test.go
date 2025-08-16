package unit

import (
	"errors"
	"testing"

	"github.com/S3ergio31/image-processing-service/login/application"
	"github.com/S3ergio31/image-processing-service/login/domain"
)

type LoginUserRepository struct {
	users map[string]domain.User
}

func (u LoginUserRepository) Save(user domain.User) {
	u.users[user.Username()] = user
}

func (u LoginUserRepository) FindByUsername(username string) (domain.User, *domain.UserNotFound) {
	user, ok := u.users[username]

	if !ok {
		return nil, &domain.UserNotFound{}
	}

	return user, nil
}

type MockTokenService struct {
	token string
}

func (m MockTokenService) Generate(secret string, username string) (string, error) {
	return m.token, nil
}

type LoginMockHasher struct {
	hashError bool
}

func (b LoginMockHasher) CompareHashAndPassword(hashedPassword string, password string) error {
	if b.hashError {
		return errors.New("hash error")
	}
	return nil
}

func mockLoginUserRepository(users map[string]string, hashError bool) LoginUserRepository {
	userEntities := make(map[string]domain.User, 0)

	for username, password := range users {
		userEntity, _ := domain.NewUser(username, password, LoginMockHasher{hashError: hashError})
		userEntities[username] = userEntity
	}

	return LoginUserRepository{users: userEntities}
}

func auth(hashError bool) application.Auth {
	return application.Auth{
		UserRepository: mockLoginUserRepository(map[string]string{"Test": "$2a$10$idfO21767DpicmjfFMBVoOUaufaZztlqZcbABAOE0gTHnPH0b151a"}, hashError),
		TokenService:   MockTokenService{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"},
	}
}

func assertInvalidCredentials(t *testing.T, err error) {
	if _, ok := err.(domain.InvalidCredentials); !ok {
		t.Errorf("Expected error: 'InvalidCredentials', given: %s", err)
	}
}

func TestLoginUser(t *testing.T) {

	token, err := auth(false).Login("Test", "Test12345*")

	if token == "" {
		t.Errorf("Expected a valid token, given: %s", err)
	}
}

func TestLoginInvalidUsername(t *testing.T) {
	_, err := auth(false).Login("Test1", "Test12345*")

	assertInvalidCredentials(t, err)
}

func TestLoginInvalidPassword(t *testing.T) {
	_, err := auth(true).Login("Test", "Sergio31")

	assertInvalidCredentials(t, err)
}
