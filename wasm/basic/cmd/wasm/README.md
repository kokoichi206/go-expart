``` sh
GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm

cd ../../assets && cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"
```

## js

https://developer.mozilla.org/ja/docs/WebAssembly/JavaScript_interface/instantiate

##  Links

- [syscall/js documentation](https://pkg.go.dev/syscall/js)
