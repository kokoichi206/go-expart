.PHONY: protos

protos:
	protoc protos/*.proto \
		--go_out=protos/currency \
		--go-grpc_out=protos/currency \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=protos
