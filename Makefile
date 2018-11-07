APPURL=github.com/alekssaul/si2ci

VERSION=$(shell ./scripts/git-version)
GOPATH_BIN:=$(shell echo ${GOPATH} | awk 'BEGIN { FS = ":" }; { print $1 }')/bin

.PHONY: all
all: build

.PHONY: build
build:
	@go build -o bin/si2ci -v ${APPURL}

.PHONY: vendor
vendor:
	@glide update --strip-vendor

.PHONY: test
test:
	@go test ./ 

.PHONY: docker
docker:
	@docker build . -t alekssaul/si2ci:v0.5
