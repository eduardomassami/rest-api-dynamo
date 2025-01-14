package bucket

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket struct {
	client     *s3.Client
	bucketName string
	objectKey  string
}

func NewBucket(client *s3.Client, bucketName string, objectKey string) *Bucket {
	return &Bucket{
		client:     client,
		bucketName: bucketName,
		objectKey:  objectKey,
	}
}

func (b *Bucket) Save(user []byte) error {

	log.Printf("Conteúdo do JSON enviado: %s", string(user))
	_, err := b.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(b.bucketName),
		Key:         aws.String(b.objectKey),
		Body:        bytes.NewReader(user),
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		return fmt.Errorf("erro ao salvar objeto no S3: %w", err)
	}

	return nil
}
