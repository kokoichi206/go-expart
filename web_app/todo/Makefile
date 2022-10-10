.PHONY: help build build-local up down logs ps test
# make と打った時に何が出てくるか
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build:
	docker build -t kokoichi0206/todo:${DOCKER_TAG} \
		--target deploy ./

build-local:	## build docker image to local development
	docker compose build --no-cache


up:	## docker compose up with hot reload
	docker compose up -d

down:
	docker compose down

logs:	## Tail docker compose logs
	docker compose logs -f

ps:	## Check container status
	docker compose ps

test:
	go test -race -shuffle=on ./...

help:	## https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
