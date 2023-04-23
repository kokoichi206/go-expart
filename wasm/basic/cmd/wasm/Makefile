build:
	GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm

build-tiny:
	cp $$(tinygo env TINYGOROOT)/targets/wasm_exec.js ../../assets/wasm_exec_tiny.js
	tinygo build -o ../../assets/main-tiny.wasm -target wasm main.go
