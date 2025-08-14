package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	setup "github.com/S3ergio31/image-processing-service/init"
	"github.com/S3ergio31/image-processing-service/login/application"
	login "github.com/S3ergio31/image-processing-service/login/infrastructure"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"github.com/S3ergio31/image-processing-service/transform/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	router := setup.Router()
	database.Refresh()
	setup.LoadEnv("testdata/.env.testing")
	imageUuid := "e9c1a020-7077-49e3-a217-72c395d371fc"
	username := "Test"
	password := "Test12345*"
	user := entities.User{
		Username: username,
		Password: "$2a$10$idfO21767DpicmjfFMBVoOUaufaZztlqZcbABAOE0gTHnPH0b151a",
	}

	database.GetDatabase().Create(&user)

	database.GetDatabase().Create(&entities.Image{
		Uuid: imageUuid,
		Type: "jpg",
		Name: "image",
		Path: "testdata/image.jpg",
		User: user,
	})

	auth := application.Auth{
		UserRepository: login.NewSqliteUserRepository(),
		TokenService:   login.GolangJwtTokenService{},
		Secret:         os.Getenv("JWT_SECRET"),
	}

	token, _ := auth.Login(username, password)

	w := httptest.NewRecorder()

	rotate := 0
	transformBody := infrastructure.TranformBody{
		Transformations: infrastructure.Transformations{
			Rotate: &rotate,
			Format: "png",
			Resize: &infrastructure.Resize{
				Width:  200,
				Height: 200,
			},
			Crop: &infrastructure.Crop{
				Width:  100,
				Height: 100,
				X:      10,
				Y:      20,
			},
		},
	}

	transformJson, _ := json.Marshal(transformBody)

	path := fmt.Sprintf("/images/%s/transform", imageUuid)
	req, _ := http.NewRequest("POST", path, strings.NewReader(string(transformJson)))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	router.ServeHTTP(w, req)

	ok := assert.Equal(t, 200, w.Code)

	if !ok {
		log.Println(w.Body)
	}

}
