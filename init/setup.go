package setup

import (
	"github.com/S3ergio31/image-processing-service/register/infrastructure"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/register", infrastructure.RegisterController)

	return router
}
