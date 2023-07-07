package handlers

import (
	"expense-management-images/src/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UploadImageHandler Upload image endpoint
func UploadImageHandler(imageCtl controllers.ImageCtl) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := imageCtl.UploadImage(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Image uploaded successfully"})
	}
}
