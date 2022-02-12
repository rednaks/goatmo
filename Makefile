BIN=goatmo
TAG := $(shell git describe --tags)
build-linux-arm64:
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-extldflags=-static -s -w" -o $(BIN)-$(TAG)-linux-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -ldflags="-extldflags=-static -s -w" -o $(BIN)-$(TAG)-linux-amd64

clean:
	rm $(BIN)-v*

all: build-linux-amd64 build-linux-arm64

