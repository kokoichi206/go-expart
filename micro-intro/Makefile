check_install:
	which swagger || Go111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swaggers

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models
#  GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
