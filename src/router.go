package main

import (
	"expense-management-images/src/controllers"
	"expense-management-images/src/handlers"
	"expense-management-images/src/middlewares"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Controllers structure used to handle requests
type Controllers struct {
	ImageController controllers.ImageCtl
}

func createRouter() *gin.Engine {
	router := gin.New()

	// Attach logger middleware
	router.Use(gin.Logger())

	// Attach recovery middleware
	router.Use(gin.Recovery())

	// Configure CORS
	router.Use(middlewares.CorsMiddleware())

	apiv1 := router.Group("/api/v1")

	// Initialize controllers
	controller := &Controllers{
		ImageController: &controllers.ImageController{},
	}

	imageDirectory := os.Getenv("IMAGE_DIRECTORY")

	router.Handle(http.MethodGet, "/lifecheck", handlers.LifeCheckHandler())
	router.StaticFS("/images", gin.Dir(imageDirectory, true))
	apiv1.Handle(http.MethodPost, "/images", handlers.UploadImageHandler(controller.ImageController))

	return router
}
