BIN=goatmo
TAG := $(shell git describe --tags --abbrev=0)
DIST_FOLDER=dist
BUILD_FOLDER=build

prepare_build:
	mkdir -p build

prepare_dist:
	mkdir -p dist

build-linux-arm64: prepare_build
	CC=aarch64-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-extldflags=-static -s -w" -o $(BUILD_FOLDER)/$(BIN)-$(TAG)-linux-arm64

build-linux-amd64: prepare_build
	GOOS=linux GOARCH=amd64 go build -ldflags="-extldflags=-static -s -w" -o $(BUILD_FOLDER)/$(BIN)-$(TAG)-linux-amd64

package: prepare_dist
	@for bin in $(shell cd $(BUILD_FOLDER) && ls $(BIN)-$(TAG)-*); do\
		echo $$bin;\
		cd $(BUILD_FOLDER) && tar czf ../$(DIST_FOLDER)/$$bin.tar.gz $$bin && cd ..;\
	done

clean:
	rm -r $(BUILD_FOLDER) $(DIST_FOLDER)

all: build-linux-amd64 build-linux-arm64 package

