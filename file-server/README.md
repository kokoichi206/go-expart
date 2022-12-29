## file-server

## Libraries

- https://github.com/hashicorp/go-hclog
  - key, value logging
- https://pkg.go.dev/golang.org/x/xerrors
  - go 2 で提案されているエラーを、go 1 向けに外部ライブラリとして提供
  - go 公式がメンテ

```sh
# binary であることを明示する。
curl localhost:9090/images/1/sfa.png --data-binary @yoiwake.png

# content-type の指定
curl -X POST localhost:9090/images/1/sfa.png -F "file=@yoiwake.png;type=image/png"
```

## Links

- [multipart/form-data: mozilla](https://developer.mozilla.org/ja/docs/Web/HTTP/Methods/POST)
  - text and binary, separated by boundary
  - REST 的にはあんまりやらんかもなぁ

## Gzip

```sh
curl -v localhost:9090/images/1/test.png -o file.png

curl -v localhost:9090/images/1/test.png --compressed -o file.png
```
