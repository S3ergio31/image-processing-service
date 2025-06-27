package integration

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	setup "github.com/S3ergio31/image-processing-service/init"
	"github.com/S3ergio31/image-processing-service/register/infrastructure"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	router := setup.Router()
	database.Refresh()

	w := httptest.NewRecorder()

	registerBody := infrastructure.RegisterBody{
		Username: "test",
		Password: "Sergio1234!",
	}

	registerJson, _ := json.Marshal(registerBody)

	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(registerJson)))
	router.ServeHTTP(w, req)

	ok := assert.Equal(t, 201, w.Code)

	if !ok {
		log.Println(w.Body)
	}

}
