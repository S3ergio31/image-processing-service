package infrastructure

import (
	"net/http"

	"github.com/S3ergio31/image-processing-service/find/application"
	"github.com/S3ergio31/image-processing-service/find/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/responses"
	"github.com/gin-gonic/gin"
)

func FinderController(c *gin.Context) {
	finder := application.Finder{
		ImageRepository: NewSqliteRegisterImageRepository(),
	}

	image, errs := finder.Find(
		c.GetString("username"),
		c.Param("id"),
	)

	if len(errs) == 1 {
		if _, ok := errs[0].(*domain.ImageNotFound); ok {
			responses.WriteNotFoundRequestResponse(errs, c)
			return
		}
	}

	if errs != nil {
		responses.WriteBadRequestResponse(errs, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid": image.Uuid(),
		"path": image.Path(),
		"name": image.Name(),
		"type": image.Type(),
	})

}
