package infrastructure

import (
	"net/http"

	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/responses"
	"github.com/S3ergio31/image-processing-service/upload/application"
	"github.com/S3ergio31/image-processing-service/upload/domain"
	"github.com/gin-gonic/gin"
)

func UploadController(c *gin.Context) {
	uploadRequest, err := BuildUploadRequest(c)

	if err != nil {
		responses.WriteBadRequestResponse([]error{err}, c)
		return
	}

	uploader := application.Uploader{
		ImageRepository: LocalDiskImageRepository{},
		EventBus:        events.New(),
	}

	uploadErrors := uploader.Upload(domain.UploadImage{
		Uuid:     uploadRequest.Uuid,
		Username: uploadRequest.Username,
		Name:     uploadRequest.Filename,
		Content:  uploadRequest.Content,
	})

	if len(uploadErrors) != 0 {
		responses.WriteBadRequestResponse(uploadErrors, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid":     c.PostForm("uuid"),
		"filename": uploadRequest.Filename,
		"size":     uploadRequest.Size,
	})
}
