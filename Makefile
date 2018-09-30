.PHONY: build test

build: test
	go build -o bin/mset cmd/*.go

build-full: test build-linux build-darwin

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/mset-linux cmd/*.go
build-darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/mset-darwin cmd/*.go
test:
	go test -race -cover *.go
