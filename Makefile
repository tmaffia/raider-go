.DEFAULT_GOAL := build

clean:
	go clean

build: clean
	go build ./...

test: build
	go test ./...

dep:
	go mod download