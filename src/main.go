package main

import (
	"log"
	"os"

	"github.com/eduardomassami/rest-api-dynamo/adapter/db"
	"github.com/eduardomassami/rest-api-dynamo/adapter/http"
	"github.com/eduardomassami/rest-api-dynamo/application"
	"github.com/eduardomassami/rest-api-dynamo/config"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	dbClient := config.GetDynamoDB()

	repo := db.NewDynamoDBRepository(dbClient, "TestTable")

	service := application.NewUserService(repo)
	handler := http.NewHandler(service)

	r := gin.Default()
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
