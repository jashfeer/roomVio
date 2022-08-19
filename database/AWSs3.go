package database

import (
	"fmt"
	"log"
	"mime/multipart"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func UploadImage(bucket string, key string, body multipart.File) string {
	godotenv.Load()

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		//return
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("roomvio"),
		Key:    aws.String(key),
		Body:   body,
		ACL:    "public-read",
	})
	if uploadErr != nil {
		fmt.Println("Failed to upload : ")
		//return
	}

	return result.Location
}
