package config

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func InitializeDynamoDB(cfg *aws.Config) *dynamodb.Client {
	if os.Getenv("LOCALSTACK") == "true" {
		return dynamodb.NewFromConfig(*cfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(os.Getenv("AWS_ENDPOINT"))
		})
	}

	return dynamodb.NewFromConfig(*cfg)

}
