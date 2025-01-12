package main

import (
	"log"

	"github.com/eduardomassami/rest-api-dynamo/adapter/db"
	"github.com/eduardomassami/rest-api-dynamo/adapter/http"
	"github.com/eduardomassami/rest-api-dynamo/application"
	"github.com/eduardomassami/rest-api-dynamo/config"
	"github.com/eduardomassami/rest-api-dynamo/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger = config.GetLogger("main")
	dbClient := config.GetDynamoDB()

	repo := db.NewDynamoDBRepository(dbClient, "TestTable")

	service := application.NewUserService(repo)
	handler := http.NewHandler(service)

	router.Initialize(handler)
}
