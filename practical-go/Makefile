.PHONY: install test run build

install:
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	go generate
	go mod download

test:
	go test -race ./...

run:
	go run main.go

build:
	go build -ldflags '-s -w -X main.version=1.0.0' main.go
