package setup

import (
	"log"

	login "github.com/S3ergio31/image-processing-service/login/infrastructure"
	register "github.com/S3ergio31/image-processing-service/register/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/register", register.RegisterController)
	router.POST("/login", login.LoginController)

	return router
}

func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		panic("Cannot load .env file: " + err.Error())
	}
	log.Println(".env file loaded")
}
