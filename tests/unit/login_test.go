package unit

import (
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

func mockLoginUserRepository(users map[string]string) LoginUserRepository {
	userEntities := make(map[string]domain.User, 0)

	for username, password := range users {
		userEntity, _ := domain.NewUser(username, password)
		userEntities[username] = userEntity
	}

	return LoginUserRepository{users: userEntities}
}

func auth() application.Auth {
	return application.Auth{
		UserRepository: mockLoginUserRepository(map[string]string{"Test": "$2a$10$idfO21767DpicmjfFMBVoOUaufaZztlqZcbABAOE0gTHnPH0b151a"}),
		TokenService:   domain.TokenService{Secret: "test_secret"},
	}
}

func assertInvalidCredentials(t *testing.T, err error) {
	if _, ok := err.(domain.InvalidCredentials); !ok {
		t.Errorf("Expected error: 'InvalidCredentials', given: %s", err)
	}
}

func TestLoginUser(t *testing.T) {

	token, err := auth().Login("Test", "Test12345*")

	if token == "" {
		t.Errorf("Expected a valid token, given: %s", err)
	}
}

func TestLoginInvalidUsername(t *testing.T) {
	_, err := auth().Login("Test1", "Test12345*")

	assertInvalidCredentials(t, err)
}

func TestLoginInvalidPassword(t *testing.T) {
	_, err := auth().Login("Test", "Sergio31")

	assertInvalidCredentials(t, err)
}
