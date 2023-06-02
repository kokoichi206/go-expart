storage-up:
	docker run \
		-p 9001:9000 \
		-p 9091:9090 \
		--user $(id -u):$(id -g) \
		--name minio \
		-e "MINIO_ROOT_USER=ROOTUSER" \
		-e "MINIO_ROOT_PASSWORD=CHANGEME123" \
		-v ${PWD}data:/data \
		quay.io/minio/minio server /data --console-address ":9090"

.DEFAULT_GOAL := run
run:
	go run *.go
