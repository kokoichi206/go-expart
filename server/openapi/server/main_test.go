package main

import (
	"encoding/json"
	"openapi-generator-server/gen/component"
	"openapi-generator-server/gen/component_method"
	"os"
	"testing"
)

func BenchmarkJsonMarshal(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")
	var res component.ListCommentsResponse
	json.Unmarshal(bytes, &res)

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(res)
	}
}

func BenchmarkJsonMarshalMethod(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")
	var res component_method.ListCommentsResponse
	json.Unmarshal(bytes, &res)

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(res)
	}
}
