package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	bucket = "minio-compatibility-test"
	object = "92074.jpg"
	region = "ap-northeast-1"

	// minio も共通で使うためには、以下の値が必要（環境変数）。
	// むしろこの１つだけで良さそう。
	// for s3
	// see: https://docs.aws.amazon.com/ja_jp/general/latest/gr/s3.html#:~:text=%E3%82%A2%E3%82%B8%E3%82%A2%E3%83%91%E3%82%B7%E3%83%95%E3%82%A3%E3%83%83%E3%82%AF%20(%E6%9D%B1%E4%BA%AC)-,ap%2Dnortheast%2D1,-%E6%A8%99%E6%BA%96%E3%82%A8%E3%83%B3%E3%83%89%E3%83%9D%E3%82%A4%E3%83%B3%E3%83%88
	// url    = "https://s3.ap-northeast-1.amazonaws.com"
	url = "http://localhost:9001"
)

func minioCompatibility() {
	// dc 内のの mc で設定した key, secret を環境変数に設定する。
	// （sdk 内で勝手に使われる。）
	os.Setenv("AWS_ACCESS_KEY_ID", "3AFDH6SThcVDnI7FsACg")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "thG04kib7mgHwjfstUCpqsb1RQyzZitmbulZKWpI")

	// ============================================
	// =============== s3 のみとの差分===============
	// ============================================
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opts ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               url,
			HostnameImmutable: true,
			SigningRegion:     region,
		}, nil
	})
	// ============================================
	// ============================================
	// ============================================

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		log.Fatal(err)
	}

	// **aws の** s3 client を作成する。
	client := s3.NewFromConfig(cfg)

	// ========= Get an object =========
	obj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		log.Fatal(err)
	}

	writeFile(obj.Body)
}

func s3native() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// **aws の** s3 client を作成する。
	client := s3.NewFromConfig(cfg)

	// ========= Get an object =========
	obj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})
	if err != nil {
		log.Fatal(err)
	}

	writeFile(obj.Body)
}

func writeFile(r io.Reader) {
	file, _ := os.Create("test.jpg")
	b, _ := io.ReadAll(r)
	file.Write(b)
}
