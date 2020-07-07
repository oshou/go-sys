export GO111MODULE=on

GO ?= $(shell which go)

$(GOPATH)/bin/golangci-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

clean:
	$(GO) mod tidy

fmt: clean
	$(GO) fmt ./...

lint: fmt $(GOPATH)/bin/golangci-lint
	golangci-lint run

test: lint
	$(GO) test ./...

cov: lint
	$(GO) test ./... -cover

.PHONY: clean fmt lint cov
