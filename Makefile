ROOT_DIR:=$(shell git rev-parse --show-toplevel)
GOFILES:=$(shell find . -name '*.go' | grep -v -E '(./vendor)')
GOFLAGS=
IMAGE_REPO=
IMAGE_TAG:=

$( shell mkdir -p bin )

build: bin/dbutil

check:
	@golint -set_exit_status $(shell go list ./...)
	@go vet $(shell go list ./...)

bin/dbutil: $(GOFILES)
	@go build $(GOFLAGS) -o $(ROOT_DIR)/bin/dbutil github.com/hexfusion/dbutil/cmd/dbutil

vendor:
	@go mod vendor -v

clean:
	rm -rf $(ROOT_DIR)/bin

.PHONY: build check clean vendor
