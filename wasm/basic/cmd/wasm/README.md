``` sh
GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm

cd ../../assets && cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"
```

## js

https://developer.mozilla.org/ja/docs/WebAssembly/JavaScript_interface/instantiate

## tiny-go

https://tinygo.org/getting-started/install/macos/

``` sh
brew tap tinygo-org/tools
brew install tinygo

$ tinygo version

tinygo version 0.27.0 darwin/amd64 (using go version go1.20 and LLVM version 15.0.0)
```

### wasm file size

``` sh
ls -lah ../../assets
...
-rwxr-xr-x  1 kokoichi  staff   2.0M Apr 23 20:54 main.wasm
-rw-r--r--  1 kokoichi  staff    16K Apr 23 19:06 wasm_exec.js

# tinyGo
ls -lah ../../assets
total 4856
-rwxr-xr-x  1 kokoichi  staff   313K Apr 23 21:20 main-tiny.wasm
-rw-r--r--  1 kokoichi  staff    16K Apr 23 21:20 wasm_exec_tiny.js
```

## Links

- [syscall/js documentation](https://pkg.go.dev/syscall/js)
