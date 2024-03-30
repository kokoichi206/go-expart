.PHONY: run
run:
	go run main.go

.PHONY: dev
dev:	## Hot reload 付きでサーバーを起動する。
	air -c .air.toml

.PHONY: web
wasm-build:	## WebAssembly をビルドする。
	GOOS=js GOARCH=wasm go build -o ./web/main.wasm ./main.go

.PHONY: web-deploy
web-deploy:
	cp web/* ../../../docs

.PHONY: prepare-wasm
prepare-wasm:	## WebAssembly 用のファイルをコピーする。
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./web
