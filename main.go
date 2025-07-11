package main

import (
	setup "github.com/S3ergio31/image-processing-service/init"
	events "github.com/S3ergio31/image-processing-service/shared/domain"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
	upload "github.com/S3ergio31/image-processing-service/upload/application"
	"github.com/S3ergio31/image-processing-service/upload/infrastructure"
)

func main() {
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
