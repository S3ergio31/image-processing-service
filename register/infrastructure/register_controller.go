package infrastructure

import (
	"net/http"

	"github.com/S3ergio31/image-processing-service/register/application"
	"github.com/gin-gonic/gin"
)

func RegisterController(c *gin.Context) {
	var user RegisterBody
	c.Bind(&user)
	register := application.Register{Repository: NewInMemoryUserRepository()}
	err := register.Save(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
	}

	c.Writer.WriteHeader(http.StatusCreated)
}
