.DEFAULT_GOAL := build

setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest

clean:
	go clean

build: clean
	staticcheck ./...
	go vet ./...
	go build ./...

test: build
	go test ./...

cover: build
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
	open cover.html

dep:
	go mod download