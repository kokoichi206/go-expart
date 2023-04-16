BINARY_NAME=Markdown.app
APP_NAME=Markdown
VERSION=1.0.0

build:
	rm -rf ${BINARY_NAME}
	rm -r fyne-md
	fyne package -appVersion 1.0.0 -name ${APP_NAME} -release

run:
	go run .

test:
	go test -v ./...
