package unit

import (
	"testing"

	"github.com/S3ergio31/image-processing-service/register/application"
	"github.com/S3ergio31/image-processing-service/register/domain"
	shared "github.com/S3ergio31/image-processing-service/shared/domain"
)

type UserRepository struct {
	users map[string]domain.User
}

func (u UserRepository) Save(user domain.User) {
	u.users[user.Username()] = user
}

func (u UserRepository) UsedUsername(username string) bool {
	_, ok := u.users[username]
	return ok
}

func mockUserRepository(users map[string]string) UserRepository {
	userEntities := make(map[string]domain.User, 0)

	for username, password := range users {
		userEntity, _ := domain.NewUser(username, password)
		userEntities[username] = userEntity
	}

	return UserRepository{users: userEntities}
}

func assertHasErrors(t *testing.T, errors []error) {
	if len(errors) == 0 {
		t.Errorf("Expected 1 error, given 0")
	}
}

func TestRegisterUser(t *testing.T) {
	register := application.Register{
		Repository: mockUserRepository(map[string]string{"AnotherUser": "Sergio1234!"}),
	}

	errors := register.Save("Test", "Sergio1234!")

	if len(errors) != 0 {
		t.Errorf("Expected 0 errors, given: %d - result: %s", len(errors), errors)
	}
}

func TestUserAlreadyExists(t *testing.T) {
	register := application.Register{
		Repository: mockUserRepository(map[string]string{"Test": "Sergio1234!"}),
	}

	errors := register.Save("Test", "Sergio1234!")

	assertHasErrors(t, errors)

	if _, ok := errors[0].(domain.UserAlreadyExists); !ok {
		t.Errorf("Expected error: 'UserAlreadyExists', given: %s", errors[0])
	}
}

func TestInvalidUsername(t *testing.T) {
	register := application.Register{
		Repository: mockUserRepository(map[string]string{}),
	}

	errors := register.Save("", "Sergio1234!")

	assertHasErrors(t, errors)

	if _, ok := errors[0].(shared.UsernameCannotBeEmpty); !ok {
		t.Errorf("Expected error: 'UsernameCannotBeEmpty', given: %s", errors[0])
	}
}

func TestInvalidPassword(t *testing.T) {
	register := application.Register{
		Repository: mockUserRepository(map[string]string{}),
	}

	errors := register.Save("Test", "Sergio1234")

	assertHasErrors(t, errors)

	if _, ok := errors[0].(domain.InvalidPassword); !ok {
		t.Errorf("Expected error: 'InvalidPassword', given: %s", errors[0])
	}
}
