package main

import (
	"benchmark/gen/component"
	"benchmark/gen/component_710"
	"benchmark/gen/component_method"
	"encoding/json"
	"os"
	"testing"
)

// ==================== Unmarshal ====================
func BenchmarkJsonUnMarshal(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res component.Pet
		json.Unmarshal(bytes, &res)
	}
}

func BenchmarkJsonUnMarshalMethod(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res component_method.Pet
		json.Unmarshal(bytes, &res)
	}
}

func BenchmarkJsonUnMarshal710(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res component_710.Pet
		json.Unmarshal(bytes, &res)
	}
}

// ==================== Marshal ====================
func BenchmarkJsonMarshal(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")
	var res component.Pet
	json.Unmarshal(bytes, &res)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(res)
	}
}

func BenchmarkJsonMarshalMethod(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")
	var res component_method.Pet
	json.Unmarshal(bytes, &res)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(res)
	}
}

func BenchmarkJsonMarshal710(b *testing.B) {
	bytes, _ := os.ReadFile("res.json")
	var res component_710.Pet
	json.Unmarshal(bytes, &res)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(res)
	}
}
