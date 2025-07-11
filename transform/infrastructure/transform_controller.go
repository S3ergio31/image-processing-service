package infrastructure

import (
	"github.com/S3ergio31/image-processing-service/transform/application"
	"github.com/S3ergio31/image-processing-service/transform/domain/transformations"
	"github.com/gin-gonic/gin"
)

func TransformController(c *gin.Context) {
	var tranformBody TranformBody
	c.BindJSON(&tranformBody)

	application.Transformer{}.Transform(&transformations.Transformations{
		Rotate:  tranformBody.Rotate,
		Format:  tranformBody.Format,
		Resize:  (*transformations.Resize)(tranformBody.Resize),
		Crop:    (*transformations.Crop)(tranformBody.Crop),
		Filters: (*transformations.Filters)(tranformBody.Filters),
	})
}
