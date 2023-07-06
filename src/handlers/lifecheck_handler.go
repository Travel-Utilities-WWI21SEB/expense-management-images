package handlers

import (
	"expense-management-images/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LifeCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := &models.LifeCheckResponse{
			Alive:   true,
			Version: "1.0.0",
		}

		c.JSON(http.StatusOK, response)
	}
}
