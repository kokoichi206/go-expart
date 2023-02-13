package main

import (
	"context"
	"practical-go/example"
)

var version string

func main() {
	// example.TestCredential()

	// example.ReceiverTest()

	// example.TagTest()

	// example.Interface()

	// example.ErrorIs()

	// // example.FatalCheck()

	// example.ReadJson()

	// example.LogPackage()

	// example.S3Example()

	// example.GoCDK()

	// example.DynamoDB()

	// example.UnknownTasks()

	// example.TimeoutContext()

	example.ErrorGroupExample(context.Background())
}
