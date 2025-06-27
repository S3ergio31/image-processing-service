package integration

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	setup "github.com/S3ergio31/image-processing-service/init"
	"github.com/S3ergio31/image-processing-service/login/infrastructure"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := setup.Router()
	database.Refresh()
	setup.LoadEnv(".env.testing")

	database.GetDatabase().Create(&entities.User{
		Username: "Test",
		Password: "$2a$10$idfO21767DpicmjfFMBVoOUaufaZztlqZcbABAOE0gTHnPH0b151a",
	})

	w := httptest.NewRecorder()

	loginBody := infrastructure.LoginBody{
		Username: "Test",
		Password: "Test12345*",
	}

	loginJson, _ := json.Marshal(loginBody)

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(loginJson)))
	router.ServeHTTP(w, req)

	ok := assert.Equal(t, 200, w.Code)

	if !ok {
		log.Println(w.Body)
	}

}
