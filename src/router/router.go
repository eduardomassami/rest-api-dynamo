package router

import (
	"log"
	"os"

	"github.com/eduardomassami/rest-api-dynamo/adapter/http"
	"github.com/gin-gonic/gin"
)

func Initialize(handler *http.Handler) {
	router := gin.Default()

	initializeRoutes(router, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
