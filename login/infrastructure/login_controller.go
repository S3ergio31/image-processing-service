package infrastructure

import (
	"net/http"
	"os"

	"github.com/S3ergio31/image-processing-service/login/application"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var login LoginBody
	c.BindJSON(&login)

	auth := application.Auth{
		UserRepository: NewSqliteUserRepository(),
		TokenService:   GolangJwtTokenService{},
		Secret:         os.Getenv("JWT_SECRET"),
	}

	token, err := auth.Login(login.Username, login.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})

}
