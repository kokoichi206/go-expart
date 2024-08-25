## update go version

``` sh
go install golang.org/dl/go1.23.0@latest
go1.23.0 download
go1.23.0 version

# ~/.zshrc など
export GOROOT=$(go1.23.0 env GOROOT)
export PATH=$GOROOT/bin:$PATH
```

## unique package

interning (インターン化) によりメモリを節約する。

- [issue(Proposal?)](https://github.com/golang/go/issues/62483)

## GODEBUG

https://tip.golang.org/doc/godebug

``` sh
go list -f '{{.DefaultGODEBUG}}'
```

## Iterator

- https://pkg.go.dev/github.com/BooleanCat/go-functional/v2@v2.0.0-beta.6/it

## Links

- [Go 1.23 Release Notes](https://tip.golang.org/doc/go1.23)
