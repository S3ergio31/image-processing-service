package infrastructure

import (
	"fmt"
	"net/http"
	"path/filepath"

	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/responses"
	"github.com/S3ergio31/image-processing-service/transform/application"
	"github.com/S3ergio31/image-processing-service/transform/domain"
	transformations "github.com/S3ergio31/image-processing-service/transform/domain"
	"github.com/gin-gonic/gin"
)

func TransformController(c *gin.Context) {
	var tranformBody TranformBody
	bindError := c.ShouldBindJSON(&tranformBody)

	if bindError != nil {
		responses.WriteBadRequestResponse([]error{bindError}, c)
		return
	}

	transformation, transformErr := application.Transformer{
		ImageEditor:     BimgImageEditor{},
		ImageStorage:    LocalDiskImageStorage{},
		ImageRepository: NewSqliteRegisterImageRepository(),
		EventBus:        events.New(),
	}.Transform(&transformations.Transformations{
		ImageUuid: c.Param("id"),
		Username:  c.GetString("username"),
		Rotate:    *tranformBody.Rotate,
		Format:    tranformBody.Format,
		Resize:    (*transformations.Resize)(tranformBody.Resize),
		Crop:      (*transformations.Crop)(tranformBody.Crop),
	})

	if transformation != nil {
		download(transformation, c)
		return
	}

	switch err := transformErr.(type) {
	case *transformations.ImageStorageGetError:
		responses.WriteConflictResponse([]error{err}, c)
	case *transformations.ImageStorageSaveError:
		responses.WriteConflictResponse([]error{err}, c)
	case *transformations.ImageNotFound:
		responses.WriteNotFoundRequestResponse([]error{err}, c)
	}
}

func download(transformation *domain.Transformation, c *gin.Context) {
	filename := filepath.Base(transformation.Path())
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	contentType := fmt.Sprintf("image/%s", transformation.Format)
	c.Data(http.StatusOK, contentType, transformation.Content)
}
