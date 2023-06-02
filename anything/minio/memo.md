## MinIO?

https://min.io/

- High Performance Object Storage for AI
  - どういう意味で for AI?
    - large scale AI/ML, data lake and database workloads
- S3 compatible object store
  - https://min.io/product/s3-compatibility
  - [aws-cli から操作する](https://dev.classmethod.jp/articles/s3-compatible-storage-minio/#aws-cli%E3%81%8B%E3%82%89%E6%93%8D%E4%BD%9C%E3%81%99%E3%82%8B:~:text=%E3%83%AD%E3%83%BC%E3%83%89%E5%AE%8C%E4%BA%86%E3%81%A7%E3%81%99%E3%80%82-,aws%2Dcli%E3%81%8B%E3%82%89%E6%93%8D%E4%BD%9C%E3%81%99%E3%82%8B,-aws%2Dcli%E3%81%8B%E3%82%89)

## Getting Started

- https://min.io/docs/minio/container/index.html

``` sh
mkdir -p ./data

docker run \
   -p 9001:9000 \
   -p 9091:9090 \
   --user $(id -u):$(id -g) \
   --name minio \
   -e "MINIO_ROOT_USER=ROOTUSER" \
   -e "MINIO_ROOT_PASSWORD=CHANGEME123" \
   -v ${PWD}data:/data \
   quay.io/minio/minio server /data --console-address ":9090"
```

http://127.0.0.1:9091/login

## SDK

https://min.io/docs/minio/linux/developers/minio-drivers.html#minio-drivers

```
.NET
Golang
Haskell
Java
JavaScript
Python
```

For golang

https://min.io/docs/minio/linux/developers/minio-drivers.html#go-sdk

Compare: https://github.com/aws/aws-sdk-go-v2

## 操作

- login
