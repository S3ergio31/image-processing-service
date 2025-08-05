package integration

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	setup "github.com/S3ergio31/image-processing-service/init"
	"github.com/S3ergio31/image-processing-service/login/application"
	"github.com/S3ergio31/image-processing-service/login/domain"
	login "github.com/S3ergio31/image-processing-service/login/infrastructure"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"github.com/stretchr/testify/assert"
)

func TestFindImage(t *testing.T) {
	router := setup.Router()
	database.Refresh()
	setup.LoadEnv("testdata/.env.testing")
	imageUuid := "e9c1a020-7077-49e3-a217-72c395d371fc"
	username := "Test"
	password := "Test12345*"
	imageName := "image"
	imageType := "jpg"
	user := entities.User{
		Username: username,
		Password: "$2a$10$idfO21767DpicmjfFMBVoOUaufaZztlqZcbABAOE0gTHnPH0b151a",
	}

	database.GetDatabase().Create(&user)

	database.GetDatabase().Create(&entities.Image{
		Uuid: imageUuid,
		Type: imageType,
		Name: imageName,
		Path: fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, imageUuid, imageName, imageType),
		User: user,
	})

	auth := application.Auth{
		UserRepository: login.NewSqliteUserRepository(),
		TokenService:   domain.TokenService{Secret: os.Getenv("JWT_SECRET")},
	}

	token, _ := auth.Login(username, password)

	w := httptest.NewRecorder()

	path := fmt.Sprintf("/images/%s", imageUuid)
	req, _ := http.NewRequest("GET", path, strings.NewReader(""))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	router.ServeHTTP(w, req)

	ok := assert.Equal(t, 200, w.Code)

	if !ok {
		log.Println(w.Body)
		t.Errorf("expected status 200 given %d\n", w.Code)
	}

}
