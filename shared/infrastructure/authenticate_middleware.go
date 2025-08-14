package infrastructure

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/S3ergio31/image-processing-service/login/infrastructure"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	tokenValidator := infrastructure.GolangJwtTokenValidator{}
	secret := os.Getenv("JWT_SECRET")

	if !strings.HasPrefix(bearerToken, "Bearer ") {
		log.Println("Invalid Authorization header")
		unauthorized(c)
		return
	}

	token := bearerToken[7:]

	if !tokenValidator.IsValid(secret, token) {
		unauthorized(c)
		return
	}

	c.Set("username", tokenValidator.Username(secret, token))

	c.Next()
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		gin.H{"error": "Authorization header required"},
	)
}
