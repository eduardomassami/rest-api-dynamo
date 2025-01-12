package config

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

var (
	dbClient *dynamodb.Client
	logger   *Logger
)

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger := NewLogger(p)
	return logger
}

func GetDynamoDB() *dynamodb.Client {
	dbClient := InitializeDynamoDB()
	return dbClient
}
