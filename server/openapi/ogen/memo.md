docker run --rm \
  --volume ".:/workspace" \
  ghcr.io/ogen-go/ogen:latest --target workspace/api3p0.yml --clean workspace/api3p0.yml


docker run --rm \
  --volume ".:/workspace" \
  ghcr.io/ogen-go/ogen:latest --target workspace/petstore --clean workspace/petstore.yml


go install github.com/ogen-go/ogen




ogen --target ../api3p0.yml --clean ../api3p0.yml

ogen --target ../api3p1.yml --clean ../api3p1.yml

