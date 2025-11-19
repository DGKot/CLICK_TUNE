BIN := "./bin/click_tune"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/app

run: build
	$(BIN) -config ./configs/config.toml

version: build
	$(BIN) version

test:
	go test -race ./...

lint: 
	golangci-lint run ./...