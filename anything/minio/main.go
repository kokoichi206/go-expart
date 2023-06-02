package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	env := env()
	fmt.Printf("env: %v\n", env)

	useMinioSDK(env)
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

type Env struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func env() Env {
	// ローカルにおいては init.go で設定する。
	return Env{
		Endpoint:        os.Getenv("MINIO_ENDPOINT"),
		AccessKeyID:     os.Getenv("MINIO_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		UseSSL:          os.Getenv("c") == "true",
	}
}
