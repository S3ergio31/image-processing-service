package infrastructure

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type UploadRequest struct {
	Uuid     string
	Username string
	Content  []byte
	*multipart.FileHeader
	multipart.File
}

func BuildUploadRequest(c *gin.Context) (*UploadRequest, error) {
	fileHeader, err := c.FormFile("image")

	if err != nil {
		return nil, errors.New("fileHeader read error: " + err.Error())
	}

	file, _, err := c.Request.FormFile("image")

	if err != nil {
		return nil, errors.New("file read error: " + err.Error())
	}

	defer file.Close()

	buffer := bytes.NewBuffer(nil)

	if _, err := io.Copy(buffer, file); err != nil {
		return nil, errors.New("Copy file error: " + err.Error())
	}

	return &UploadRequest{
		Uuid:       c.PostForm("uuid"),
		Username:   c.GetString("username"),
		FileHeader: fileHeader,
		File:       file,
		Content:    buffer.Bytes(),
	}, nil
}
