LOCAL_BIN:=$(CURDIR)/bin

install-dependencies:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.0
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@v1.12.0

.PHONY: lint
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.20.0

.PHONY: test
test:
	go clean -testcache
	$(LOCAL_BIN)/gotestsum -- -race -v -cover ./...

.PHONY: format
format:
	$(LOCAL_BIN)/goimports -local assesment_1 -w .

.PHONY: build
build:
	go build -o bin/filerec -ldflags="-s -w" ./cmd/main.go