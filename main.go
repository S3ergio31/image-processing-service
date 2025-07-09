package main

import (
	setup "github.com/S3ergio31/image-processing-service/init"
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	upload "github.com/S3ergio31/image-processing-service/upload/application"
	"github.com/S3ergio31/image-processing-service/upload/infrastructure"
)

func main() {
	/*router.GET("/test", func(c *gin.Context) {
		buffer, err := bimg.Read("image.jpg")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		newImage, err := bimg.NewImage(buffer).Resize(800, 600)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		size, err := bimg.NewImage(newImage).Size()
		if size.Width == 800 && size.Height == 600 {
			fmt.Println("The image size is valid")
		}

		bimg.Write("new.jpg", newImage)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})*/

	/*images := router.Group("/images")
	images.POST("/", uploader.UploadController)
	images.POST("/:id/transform", transformer.TransformController)
	images.GET("/:id", Finder.FindController)
	images.GET("/", Paginator.PaginateController)*/

	bus := events.New()
	bus.Subscribe(events.ImageUploaded, func(e events.Event) {
		if evt, ok := e.(events.ImageUploadedEvent); ok {
			register := upload.Register{RegisterImageRepository: infrastructure.NewSqliteRegisterImageRepository()}
			register.Store(evt)
		}
	})
	database.GetDatabase()
	setup.LoadEnv(".env")
	setup.Router().Run()
}
