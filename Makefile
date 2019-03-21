ROOT_DIR:=$(shell git rev-parse --show-toplevel)
GOFILES:=$(shell find . -name '*.go' | grep -v -E '(./vendor)')
GOFLAGS=
IMAGE_REPO=
IMAGE_TAG:=

$( shell mkdir -p bin )

build: bin/dbinit

check:
	@golint -set_exit_status $(shell go list ./...)
	@go vet $(shell go list ./...)

bin/dbinit: $(GOFILES)
	@go build $(GOFLAGS) -o $(ROOT_DIR)/bin/dbinit github.com/hexfusion/dbinit/cmd/dbinit

vendor:
	@go mod vendor -v

clean:
	rm -rf $(ROOT_DIR)/bin

.PHONY: build check clean vendor
