BIN=goatmo
build-linux-arm64:
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-extldflags=-static" -o ${BIN}-linux-arm64

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ${BIN}-linux-amd64

