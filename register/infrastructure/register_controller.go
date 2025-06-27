package infrastructure

import (
	"net/http"

	"github.com/S3ergio31/image-processing-service/register/application"
	"github.com/gin-gonic/gin"
)

func RegisterController(c *gin.Context) {
	var user RegisterBody
	c.BindJSON(&user)
	register := application.Register{UserRepository: NewSqliteUserRepository()}
	errors := register.Store(user.Username, user.Password)

	if len(errors) != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"errors": ErrorsToStrings(errors),
		})
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}

func ErrorsToStrings(errs []error) []string {
	out := make([]string, len(errs))
	for i, e := range errs {
		out[i] = e.Error()
	}
	return out
}
