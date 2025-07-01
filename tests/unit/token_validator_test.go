package unit

import (
	"testing"

	"github.com/S3ergio31/image-processing-service/login/domain"
)

func getToken() string {
	token, _ := domain.TokenService{Secret: "test_secret"}.Generate("Test")

	return token
}

func TestValidToken(t *testing.T) {

	tokenValidator := domain.TokenValidator{Secret: "test_secret"}

	if !tokenValidator.IsValid(getToken()) {
		t.Errorf("Expected a valid token, given invalid")
	}
}

func TestInvalidToken(t *testing.T) {

	tokenValidator := domain.TokenValidator{Secret: "test_secret_2"}

	if tokenValidator.IsValid(getToken()) {
		t.Errorf("Expected a invalid token, given valid")
	}
}

func TestRetrieveUsernameFromToken(t *testing.T) {

	tokenValidator := domain.TokenValidator{Secret: "test_secret"}

	username := tokenValidator.Username(getToken())
	if username != "Test" {
		t.Errorf("Expected username Sergio, given %s", username)
	}
}
