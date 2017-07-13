SHELL := /bin/sh
GOPROCS := 4
SRC := $(wildcard *.go)
COVFILE := coverage.out
GOFILES_NOVENDOR := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

help:                  ## Show this help.
	@echo -e "Rules:\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: all
all: vet test          ## Run the targets 'vet' and 'test'

.PHONY: clean
clean:                 ## Clean removes object files from package source directories
	@go clean -i

.PHONY: fmtcheck
fmtcheck:              ## Checks the files format
	@if [ "$(shell gofmt -l $(GOFILES_NOVENDOR) | wc -l)" != "0" ]; then \
		echo "Files missing go fmt: $(shell gofmt -l $(GOFILES_NOVENDOR))"; exit 2; \
	fi
	@gofmt -l $(GOFILES_NOVENDOR)

.PHONY: format
format:                ## Checks the files format
	@go fmt

.PHONY: cov
cov: $(COVFILE)        ## Shows coverage
	@go tool cover -func=$(COVFILE)

.PHONY: htmlcov
htmlcov: $(COVFILE)    ## Shows coverage in html
	@go tool cover -html=$(COVFILE)

$(COVFILE):            ## Creates the coverage.out
	@go test -covermode=count -coverprofile=$(COVFILE)

.PHONY: test
test:                  ## Run the tests
	@go test -v -coverprofile=$(COVFILE)

.PHONY: vet
vet:                   ## Run go vet
	@go vet
