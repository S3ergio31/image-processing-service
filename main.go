package main

import (
	setup "github.com/S3ergio31/image-processing-service/init"
	"github.com/S3ergio31/image-processing-service/shared/infrastructure/database"
)

func main() {
	setup.RegisterEvents()
	database.GetDatabase()
	setup.LoadEnv(".env")
	setup.Router().Run()
}
