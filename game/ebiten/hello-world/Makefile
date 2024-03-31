.PHONY: run
run:
	go run main.go

.PHONY: dev
dev:	## Hot reload 付きでサーバーを起動する。
	air -c .air.toml

.PHONY: wasm-build
wasm-build:	## WebAssembly をビルドする。
	GOOS=js GOARCH=wasm go build -o ./web/main.wasm ./main.go

.PHONY: web-deploy
web-deploy:
	cp web/* ../../../docs

.PHONY: android-build
android-build:	## Android 用の aar ファイルをビルドする。
	ebitenmobile bind -target android -androidapi 33 -javapkg jp.mydns.kokoichi206.game.cat -o ./mobile/android/cat/cat.aar ./mobile

.PHONY: ios-build
ios-build:	## iOS 用の framework ファイルをビルドする。
	ebitenmobile bind -target ios -o ./mobile/ios/Mobile.xcframework ./mobile

.PHONY: prepare-wasm
prepare-wasm:	## WebAssembly 用のファイルをコピーする。
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./web
