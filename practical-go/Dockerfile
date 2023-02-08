FROM golang:1.19-bullseye AS builder
# WORKDIR /opt/app
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY ./*.go ./
# COPY ./cmd/sample/*.go ./cmd/sample/
# RUN cd cmd/sample && go build -trimpath -ldflags="-w -s" -o "sample-app"
RUN go build -ldflags="-w -s" -o /go/bin/app

FROM gcr.io/distroless/base-debian11
# COPY --from=builder /opt/app/cmd/sample/sample-app /sapmle-app
COPY --from=builder /go/bin/app /
# CMD ["/sample-app"]
CMD ["/app"]
