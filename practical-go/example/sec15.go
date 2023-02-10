package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

func S3Example() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load sdk configuration %v", err)
	}

	client := s3.NewFromConfig(cfg)

	var token *string
	for {
		resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("kubernetes-test-kokoichi"),
			// Prefix: aws.String(""), // 特定の文字から始まるファイルに絞る設定
			ContinuationToken: token,
		})
		if err != nil {
			log.Fatalf("list objects, %v", err)
		}

		for _, c := range resp.Contents {
			fmt.Printf("Name:%s LastModified:%s\n", *c.Key, c.LastModified.Format(time.RFC3339))
		}

		if resp.ContinuationToken == nil { // ページング
			break
		}
		token = resp.ContinuationToken
	}
}

func GoCDK() {
	ctx := context.TODO()

	bucket, err := blob.OpenBucket(ctx, "s3://kubernetes-test-kokoichi")
	if err != nil {
		log.Fatal(err)
	}
	defer bucket.Close()

	var token = blob.FirstPageToken
	for {
		// opts := &blob.ListOptions{
		// 	Prefix: "",
		// }
		objs, nextToken, err := bucket.ListPage(ctx, token, 10, nil)
		if err != nil {
			log.Fatalf("list objects, %v", err)
		}
		for _, obj := range objs {
			fmt.Printf("Name: %s, LastModified: %s\n", obj.Key, obj.ModTime.Format(time.RFC3339))
		}

		if nextToken == nil {
			break
		}
		token = nextToken
	}
}

type Item struct {
	ID            string `dynamodbav:"id"`
	ProcessDate   string `dynamodbav:"process_date"`
	Text          string `dynamodbav:"text"`
	TextOmitEmpty string `dynamodbav:"text_omit_empty,omitempty"`
}

func DynamoDB() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load SDK config: %v", err)
	}
	db := dynamodb.NewFromConfig(cfg)

	item := Item {
		ID: "0001",
		ProcessDate: time.Now().Format("2020-02-06"),
		Text: "EXAMPLE TEXT",
		TextOmitEmpty: "",
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("failed to marshal, item = %v, %v", item, err)
	}
	in := &dynamodb.PutItemInput{
		TableName: aws.String("example"),
		Item: av,
	}

	_, err = db.PutItem(ctx, in)
	if err != nil {
		log.Fatalf("failed to put item %v", err)
	}
}
