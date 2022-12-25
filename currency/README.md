## Setup

https://grpc.io/docs/languages/go/quickstart/

```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

- https://github.com/grpc/grpc-go

```sh
protoc --version
libprotoc 3.21.7
```

## [grpcurl](https://github.com/fullstorydev/grpcurl)

```sh
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

grpcurl --plaintext localhost:9092 list
Failed to list services: server does not support the reflection API

grpcurl --plaintext localhost:9092 list Currency
Currency.GetRate

grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );

grpcurl --plaintext localhost:9092 describe .RateRequest
RateRequest is a message:
message RateRequest {
  string Base = 1;
  string Destination = 2;
}

grpcurl --plaintext -d '{"Base": "GBP", "Destination": "USD"}' localhost:9092 Currency.GetRate
{
  "Rate": 0.5
}
```
