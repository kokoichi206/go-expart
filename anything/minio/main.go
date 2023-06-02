package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// debug.Stack()

	env := env()
	fmt.Printf("env: %v\n", env)

	// useMinioSDK(env)

	useAWSSDK(env)
}

func useMinioSDK(env Env) {
	// Initialize minio client object.
	minioClient, err := minio.New(
		env.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(env.AccessKeyID, env.SecretAccessKey, ""),
			Secure: env.UseSSL,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient)

	ctx := context.Background()
	bis, err := minioClient.ListBuckets(ctx)
	if err != nil {
		log.Fatal(err)
	}

	target := "mybucket"
	found := false
	for _, bi := range bis {
		log.Printf("%#v\n", bi)
		if bi.Name == target {
			found = true

			break
		}
	}
	if !found {
		if err := minioClient.MakeBucket(ctx, target, minio.MakeBucketOptions{Region: "us-east-1"}); err != nil {
			log.Fatalln(err)
		}
	}
	info, err := minioClient.FPutObject(ctx, target, "myobject", "testdata/test.txt", minio.PutObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("info: %v\n", info)

	obj, err := minioClient.GetObject(ctx, target, "myobject", minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
}

// see: https://github.com/aws/aws-sdk-go-v2/blob/main/example/service/s3/listObjects/listObjects.go
func useAWSSDK(env Env) {
	// aws sdk での default-client は環境変数を参照する。
	os.Setenv("AWS_ACCESS_KEY_ID", env.AccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", env.SecretAccessKey)

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               env.URL,
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		log.Fatal(err)
	}

	// **aws の** s3 client を作成する。
	client := s3.NewFromConfig(cfg)

	// ========= list objects =========
	lo, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("mybucket"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, o := range lo.Contents {
		fmt.Printf("o: %v\n", o)
		fmt.Printf("o.Size: %v\n", o.Size)
	}

	// ========= Get an object =========
	obj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("mybucket"),
		Key:    aws.String("myobject"),
	})
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(obj.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("string(b): %v\n", string(b))
}

type Env struct {
	URL             string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func env() Env {
	endpoint := os.Getenv("MINIO_ENDPOINT")

	// ローカルにおいては init.go で設定する。
	return Env{
		URL: fmt.Sprintf("http://%s", endpoint),
		// MinioSDK が設定値としてドメイン名を要求するため URL とは別に用意する。
		// aws sdk のみ使うのであれば不要。
		Endpoint:        endpoint,
		AccessKeyID:     os.Getenv("MINIO_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		UseSSL:          os.Getenv("c") == "true",
	}
}
