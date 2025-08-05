package setup

import (
	"log"

	find "github.com/S3ergio31/image-processing-service/find/infrastructure"
	login "github.com/S3ergio31/image-processing-service/login/infrastructure"
	register "github.com/S3ergio31/image-processing-service/register/infrastructure"
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	shared "github.com/S3ergio31/image-processing-service/shared/infrastructure"
	transform "github.com/S3ergio31/image-processing-service/transform/infrastructure"
	application_upload "github.com/S3ergio31/image-processing-service/upload/application"
	"github.com/S3ergio31/image-processing-service/upload/infrastructure"
	upload "github.com/S3ergio31/image-processing-service/upload/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/register", register.RegisterController)
	router.POST("/login", login.LoginController)

	images := router.Group("/images", shared.Authenticate)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	images.POST("/", upload.UploadController)
	images.POST("/:id/transform", shared.RateLimiter(5), transform.TransformController)
	images.GET("/:id", find.FinderController)
	images.GET("/", func(ctx *gin.Context) {})

	return router
}

func LoadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		panic("Cannot load .env file: " + err.Error())
	}
	log.Println(".env file loaded")
}

func RegisterEvents() {
	bus := events.New()
	bus.Subscribe(events.ImageUploaded, func(e events.Event) {
		if evt, ok := e.(events.ImageUploadedEvent); ok {
			register := application_upload.Register{RegisterImageRepository: infrastructure.NewSqliteRegisterImageRepository()}
			register.Store(evt)
		}
	})
	bus.Subscribe(events.ImageTransformed, func(e events.Event) {
		if evt, ok := e.(events.ImageTransformedEvent); ok {
			log.Println(evt)
		}
	})
}
