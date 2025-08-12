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

func TestPaginateImages(t *testing.T) {
	router := setup.Router()
	database.Refresh()
	setup.LoadEnv("testdata/.env.testing")
	image1Uuid := "ed8671d9-7129-49b2-84c1-7c6e4d51f716"
	image2Uuid := "6f4a7da6-1761-4a0f-b364-d69ffbb3ddd7"
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
		Uuid: image1Uuid,
		Type: imageType,
		Name: imageName,
		Path: fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, image1Uuid, imageName, imageType),
		User: user,
	})

	database.GetDatabase().Create(&entities.Image{
		Uuid: image2Uuid,
		Type: imageType,
		Name: imageName,
		Path: fmt.Sprintf("images/%s/uploads/%s_%s.%s", username, image2Uuid, imageName, imageType),
		User: user,
	})

	auth := application.Auth{
		UserRepository: login.NewSqliteUserRepository(),
		TokenService:   domain.TokenService{Secret: os.Getenv("JWT_SECRET")},
	}

	token, _ := auth.Login(username, password)

	w := httptest.NewRecorder()

	path := fmt.Sprintf("/images?page=1&limit=1")
	req, _ := http.NewRequest("GET", path, strings.NewReader(""))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	router.ServeHTTP(w, req)

	ok := assert.Equal(t, 200, w.Code)

	assert.Contains(t, w.Body.String(), image1Uuid)

	if !ok {
		log.Println(w.Body)
		t.Errorf("expected status 200 given %d\n", w.Code)
	}

}
