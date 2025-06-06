package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
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
	})
	router.Run()
}
