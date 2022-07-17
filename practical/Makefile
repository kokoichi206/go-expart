# メタ情報
NAME := myproj
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# テストを実行する
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...
