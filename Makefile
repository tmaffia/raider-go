.DEFAULT_GOAL := build

clean:
	go clean

build: clean
	go build ./...

test: build
	go test ./...

cover: build
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
	open cover.html

dep:
	go mod download