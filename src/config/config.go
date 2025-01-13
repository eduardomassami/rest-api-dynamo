package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	dbClient *dynamodb.Client
	s3Client *s3.Client
	logger   *Logger
)

func InitializeAWS() (*s3.Client, *dynamodb.Client) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	s3Client := InitializeS3(&cfg)
	dbClient := InitializeDynamoDB(&cfg)

	return s3Client, dbClient

}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger := NewLogger(p)
	return logger
}
