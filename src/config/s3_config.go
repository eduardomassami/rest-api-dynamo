package config

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitializeS3(cfg *aws.Config) *s3.Client {
	if os.Getenv("LOCALSTACK") == "true" {
		return s3.NewFromConfig(*cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(os.Getenv("AWS_ENDPOINT"))
		})
	}

	return s3.NewFromConfig(*cfg)
}
