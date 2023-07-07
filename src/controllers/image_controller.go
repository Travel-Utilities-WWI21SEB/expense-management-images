package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageCtl interface {
	UploadImage(c *gin.Context) error
}

type ImageController struct {
}

var mimeExtensions = map[string]string{
	"image/jpeg": ".jpg",
	"image/jpg":  ".jpg",
	"image/png":  ".png",
}

func (ctl *ImageController) UploadImage(c *gin.Context) error {
	// Get Multipart Form
	err := c.Request.ParseMultipartForm(3 << 20) // 3 MB
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		return err
	}

	// Get user id
	userId := c.Request.FormValue("userId")

	// Get File
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		log.Printf("Error getting file: %v", err)
		return err
	}
	defer func(file multipart.File) error {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
			return err
		}

		return nil
	}(file)

	// Create and Write to a new file at /app/images/avatar_<userId>.<fileExtension>
	fileExtension, err := getFileExtension(file)

	var directoyPrefix string
	if os.Getenv("ENVIRONMENT") == "DEV" {
		directoyPrefix = "."
	}
	directoryPath := fmt.Sprintf("%s./app/images/", directoyPrefix)

	out, err := os.Create(fmt.Sprintf("%savatar_%s%s", directoryPath, userId, fileExtension))
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return err
	}
	defer func(out *os.File) error {
		err := out.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
			return err
		}

		return nil
	}(out)

	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Error copying file: %v", err)
		return err
	}

	return nil
}

func getFileExtension(file multipart.File) (string, error) {
	fileHeader := make([]byte, 512)
	_, err := file.Read(fileHeader)
	if err != nil {
		log.Printf("Error reading file header: %v", err)
		return "", err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Printf("Error seeking file: %v", err)
		return "", err
	}
	mime := http.DetectContentType(fileHeader)
	ext, ok := mimeExtensions[mime]
	if !ok {
		log.Printf("Unsupported file type: %s", mime)
		return "", fmt.Errorf("unsupported file type: %s", mime)
	}
	return ext, nil
}
