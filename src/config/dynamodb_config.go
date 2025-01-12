package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InitializeDynamoDB() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	if os.Getenv("LOCALSTACK") == "true" {
		return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(os.Getenv("AWS_ENDPOINT"))
		})
	}

	return dynamodb.NewFromConfig(cfg)

}
