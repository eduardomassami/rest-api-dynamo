package main

import (
	"log"

	"github.com/eduardomassami/rest-api-dynamo/adapter/bucket"
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

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger = config.GetLogger("main")
	s3Client, dbClient := config.InitializeAWS()

	repo := db.NewDynamoDBRepository(dbClient, "TestTable")
	bucket := bucket.NewBucket(s3Client, "test-bucket", "last")

	service := application.NewUserService(repo, bucket)
	handler := http.NewHandler(service)

	router.Initialize(handler)
}
