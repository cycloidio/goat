SHELL := /bin/sh
GOPROCS := 4
SRC := $(wildcard *.go)
COVFILE := coverage.out

.PHONY: all
all: vet test

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: format
format:
	go fmt ./...

#.PHONY: cov
cov: $(COVFILE)
	go tool cover -func=$(COVFILE)

.PHONY: htmlcov
htmlcov: $(COVFILE)
	go tool cover -html=$(COVFILE)

$(COVFILE):
	go test ./... -covermode=count -coverprofile=$(COVFILE)

.PHONY: test
test:
	go test -v ./... -coverprofile=$(COVFILE)

.PHONY: vet
vet:
	go vet ./...
