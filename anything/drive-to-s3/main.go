package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func upload(ctx context.Context, objectName string, object io.Reader, len int64) {
	bucket := os.Getenv("S3_BUCKET")

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	if _, err := client.PutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket:        aws.String(bucket),
			Key:           aws.String(objectName),
			Body:          object,
			ContentLength: &len,
		},
		s3.WithAPIOptions(
			v4.SwapComputePayloadSHA256ForUnsignedPayloadMiddleware,
		),
	); err != nil {
		log.Fatal(err)
	}
}

var driveURLRegex = regexp.MustCompile("https://drive.google.com/file/d/([^/]*)/view")

func fetchAndUpload(ctx context.Context, viewURL string, fileName string) {

	matched := driveURLRegex.FindStringSubmatch(viewURL)
	fmt.Printf("matched: %v\n", matched)
	if len(matched) < 2 {
		panic("No match found")
	}

	fileID := matched[1]
	downloadURL := fmt.Sprintf("https://drive.usercontent.google.com/download?id=%s&export=download&authuser=0", fileID)
	fmt.Printf("downloadURL: %v\n", downloadURL)

	res, err := http.Get(downloadURL)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	upload(ctx, "nanka.pdf", res.Body, res.ContentLength)
}

func main() {
	ctx := context.Background()
	fetchAndUpload(ctx, os.Getenv("DRIVE_URL"), "nanka.pdf")
}
