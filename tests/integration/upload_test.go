package integration

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	setup "github.com/S3ergio31/image-processing-service/init"
	"github.com/S3ergio31/image-processing-service/login/application"
	"github.com/S3ergio31/image-processing-service/login/domain"
	login "github.com/S3ergio31/image-processing-service/login/infrastructure"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database/entities"
	"github.com/stretchr/testify/assert"
)

func TestUploadImage(t *testing.T) {
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

	auth := application.Auth{
		UserRepository: login.NewSqliteUserRepository(),
		TokenService:   domain.TokenService{Secret: os.Getenv("JWT_SECRET")},
	}

	token, _ := auth.Login(username, password)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("uuid", imageUuid)
	fileWriter, _ := writer.CreateFormFile("image", "image.jpg")
	testFile, _ := os.Open("testdata/image.jpg")
	defer testFile.Close()
	io.Copy(fileWriter, testFile)
	writer.Close()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/images/", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	router.ServeHTTP(w, req)

	ok := assert.Equal(t, 200, w.Code)

	if !ok {
		log.Println(w.Body)
	}

}
