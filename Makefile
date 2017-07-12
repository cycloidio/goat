SHELL := /bin/sh
GOPROCS := 4
SRC := $(wildcard *.go)
COVFILE := coverage.out
GOFILES_NOVENDOR := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all
all: vet test

.PHONY: clean
clean:
	go clean -i

.PHONY: fmtcheck
fmtcheck:
	@if [ "$(shell gofmt -l $(GOFILES_NOVENDOR) | wc -l)" != "0" ]; then \
		echo "Files missing go fmt: $(shell gofmt -l $(GOFILES_NOVENDOR))"; exit 2; \
	fi
	gofmt -l $(GOFILES_NOVENDOR)

.PHONY: format
format:
	go fmt

.PHONY: cov
cov: $(COVFILE)
	go tool cover -func=$(COVFILE)

.PHONY: htmlcov
htmlcov: $(COVFILE)
	go tool cover -html=$(COVFILE)

$(COVFILE):
	go test  -covermode=count -coverprofile=$(COVFILE)

.PHONY: test
test:
	go test -v -coverprofile=$(COVFILE)

.PHONY: vet
vet:
	go vet
