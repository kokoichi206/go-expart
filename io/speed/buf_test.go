package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark_CopyToFile(b *testing.B) {
	f, _ := os.Open("large_text")
	defer f.Close()

	for n := 0; n < b.N; n++ {
		dst, _ := os.Create("tmp")
		io.Copy(dst, f)
	}

	os.Remove("tmp")
}

func Benchmark_CopyToFile2(b *testing.B) {
	f, _ := os.Open("large_text")
	defer f.Close()

	dst, _ := os.Create("tmp")
	for n := 0; n < b.N; n++ {
		io.Copy(dst, f)
	}

	os.Remove("tmp")
}

func Benchmark_CopyToBuf(b *testing.B) {
	f, _ := os.Open("large_text")
	defer f.Close()
	var buf bytes.Buffer

	for n := 0; n < b.N; n++ {
		io.Copy(&buf, f)
	}
}

func Benchmark_ReadAll(b *testing.B) {
	f, _ := os.Open("large_text")
	defer f.Close()

	var pien []byte
	for n := 0; n < b.N; n++ {
		pien, _ = io.ReadAll(f)
	}

	assert.NotNil(b, pien)
}
