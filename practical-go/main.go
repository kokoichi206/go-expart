package main

import (
	"practical-go/example"
)

var version string

func main() {
	example.TestCredential()

	example.ReceiverTest()

	example.TagTest()

	example.Interface()

	example.ErrorIs()

	// example.FatalCheck()

	example.ReadJson()
}
